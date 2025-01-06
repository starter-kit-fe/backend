package repository

import (
	"admin/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	IsEmailExists(email string) bool
	GetUserSecretKey(userID uint) (string, error)
	GenerateNewSecretKey(userID uint) (string, error)
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	FindByID(id uint) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) IsEmailExists(email string) bool {
	var count int64
	err := r.db.Model(&model.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		// 处理其他错误（可以根据需要记录日志或处理）
		return false
	}
	return count > 0
}
func (r *userRepository) GetUserSecretKey(userID uint) (string, error) {
	var user model.User
	if err := r.db.Select("uuid").First(&user, userID).Error; err != nil {
		return "", err
	}
	return user.UUID, nil
}

func (r *userRepository) GenerateNewSecretKey(userID uint) (string, error) {
	newSecretKey := uuid.New().String()
	if err := r.db.Model(&model.User{}).Where("id = ?", userID).
		Update("uuid", newSecretKey).Error; err != nil {
		return "", err
	}
	return newSecretKey, nil
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
