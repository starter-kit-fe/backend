package service

import (
	"context"
	"fmt"
	"strings"
	"time"
	"admin/internal/constant"
	"admin/internal/dto"
	"admin/internal/model"
	"admin/internal/repository"
	"admin/pkg/cloudflare"
	"admin/pkg/email"
	"admin/pkg/google"
	"admin/pkg/jwt"
	"admin/pkg/totp"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

const (
	tokenExpireDuration = constant.JWT_EXP
	verificationCodeTTL = 5 * time.Minute
)

type UserService interface {
	GoogleSignin(ctx context.Context, token string) (string, error)
	Signout(ctx context.Context, id uint) error
	SendVerificationCode(ctx context.Context, token, email, ip string) error
	IsEmailExists(email string) bool
	Signup(ctx context.Context, req *dto.SignupRequest, ip string) (string, error)
	Signin(ctx context.Context, req *dto.SigninRequest, ip string) (string, error)
	GetUserByID(id uint) (*dto.UserResponse, error)
}

type userService struct {
	userRepo     repository.UserRepository
	cfClient     *cloudflare.Client
	emailService *email.Service
	googleServer google.GoogleService
	totpClient   totp.TOTPGenerator
	rdb          *redis.Client
	jwt          *jwt.JWTMaker
}

type UserServiceConfig struct {
	UserRepo     repository.UserRepository
	CfClient     *cloudflare.Client
	EmailService *email.Service
	GoogleServer google.GoogleService
	TotpClient   totp.TOTPGenerator
	RDB          *redis.Client
	JWT          *jwt.JWTMaker
}

func NewUserService(cfg UserServiceConfig) *userService {
	return &userService{
		userRepo:     cfg.UserRepo,
		cfClient:     cfg.CfClient,
		emailService: cfg.EmailService,
		googleServer: cfg.GoogleServer,
		totpClient:   cfg.TotpClient,
		rdb:          cfg.RDB,
		jwt:          cfg.JWT,
	}
}

func (s *userService) GoogleSignin(ctx context.Context, token string) (string, error) {
	googleUser, err := s.googleServer.GetGoogleUserInfo(ctx, token)
	if err != nil {
		return "", err
	}

	user, err := s.getOrCreateUser(googleUser)
	if err != nil {
		return "", err
	}

	return s.generateAndStoreToken(ctx, user)
}

func (s *userService) getOrCreateUser(googleUser *google.GoogleUser) (*model.User, error) {
	exists := s.userRepo.IsEmailExists(googleUser.Email)
	if exists {
		return s.userRepo.FindByEmail(googleUser.Email)
	}

	uuid := uuid.NewString()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(googleUser.Email), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		UUID:     uuid,
		NickName: googleUser.Name,
		Password: string(hashedPassword),
		Email:    googleUser.Email,
		Avatar:   googleUser.Picture,
	}
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) generateAndStoreToken(ctx context.Context, user *model.User) (string, error) {
	token, err := s.jwt.CreateToken(user.ID, user.UUID, tokenExpireDuration)
	if err != nil {
		return "", err
	}
	if err := s.rdb.Set(ctx, user.UUID, token, tokenExpireDuration).Err(); err != nil {
		return "", err
	}
	return token, nil
}

func (s *userService) Signout(ctx context.Context, id uint) error {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return err
	}
	return s.rdb.Del(ctx, user.UUID).Err()
}

func (s *userService) SendVerificationCode(ctx context.Context, token, email, ip string) error {
	if err := s.verifyCloudflareToken(ctx, token, ip); err != nil {
		return err
	}

	if exists, err := s.rdb.Exists(ctx, email).Result(); err != nil {
		return err
	} else if exists > 0 {
		return fmt.Errorf("验证码五分钟内有效，检查邮箱。请勿重复发送！")
	}

	code, err := s.generateAndSendVerificationCode(email)
	if err != nil {
		return err
	}

	return s.rdb.Set(ctx, email, code, verificationCodeTTL).Err()
}

func (s *userService) verifyCloudflareToken(ctx context.Context, token, ip string) error {
	req := &cloudflare.VerifyRequest{
		Token:    token,
		RemoteIP: ip,
	}
	_, err := s.cfClient.Verify(ctx, req)
	return err
}

func (s *userService) generateAndSendVerificationCode(email string) (string, error) {
	key, err := s.totpClient.GenerateOTP(email)
	if err != nil {
		return "", err
	}
	code, err := s.totpClient.GenerateTotpCode(key.Secret())
	if err != nil {
		return "", err
	}
	if err := s.emailService.SendVerificationCode(email, code); err != nil {
		return "", err
	}
	return code, nil
}

func (s *userService) IsEmailExists(email string) bool {
	return s.userRepo.IsEmailExists(email)
}

func (s *userService) Signup(ctx context.Context, req *dto.SignupRequest, ip string) (string, error) {
	if err := s.verifyCloudflareToken(ctx, req.Token, ip); err != nil {
		return "", err
	}

	code, err := s.rdb.Get(ctx, req.Email).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("验证码已过期，请重新获取")
		}
		return "", err
	}
	if code != req.Code {
		return "", fmt.Errorf("验证码错误")
	}

	user, err := s.createUser(req)
	if err != nil {
		return "", err
	}

	if err := s.rdb.Del(ctx, req.Email).Err(); err != nil {
		return "", err
	}

	return s.generateAndStoreToken(ctx, user)
}

func (s *userService) createUser(req *dto.SignupRequest) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	nickName := strings.Split(req.Email, "@")[0]
	uuid := uuid.NewString()
	user := &model.User{
		UUID:     uuid,
		NickName: nickName,
		Password: string(hashedPassword),
		Email:    req.Email,
	}
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Signin(ctx context.Context, req *dto.SigninRequest, ip string) (string, error) {
	if err := s.verifyCloudflareToken(ctx, req.Token, ip); err != nil {
		return "", err
	}

	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", fmt.Errorf("密码错误，请重新输入")
	}

	return s.generateAndStoreToken(ctx, user)
}

func (s *userService) GetUserByID(id uint) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:        user.ID,
		NickName:  user.NickName,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Phone:     user.Phone,
		Gender:    user.Gender,
		CreatedAt: user.CreatedAt.Format(constant.TIME_FORMAT),
		UpdatedAt: user.UpdatedAt.Format(constant.TIME_FORMAT),
	}, nil
}
