package middleware

import (
	"context"
	"strings"
	"admin/internal/repository"
	"admin/pkg/jwt"
	"admin/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type JWTAuthMiddleware struct {
	jwtMaker *jwt.JWTMaker
	userRepo repository.UserRepository
	rdb      *redis.Client
}

func NewJWTAuthMiddleware(jwtMaker *jwt.JWTMaker, userRepo repository.UserRepository, rdb *redis.Client) *JWTAuthMiddleware {
	return &JWTAuthMiddleware{
		jwtMaker: jwtMaker,
		userRepo: userRepo,
		rdb:      rdb,
	}
}

func (m *JWTAuthMiddleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.FailWithCode(c, response.UNAUTHORIZED)
			c.Abort()
			return
		}
		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			response.FailWithCode(c, response.UNAUTHORIZED)
			c.Abort()
			return
		}
		authType := strings.ToLower(fields[0])
		if authType != "bearer" {
			response.FailWithCode(c, response.UNAUTHORIZED)
			c.Abort()
			return
		}
		accessToken := fields[1]
		preliminaryClaims, err := m.jwtMaker.ParseTokenWithoutVerification(accessToken)
		if err != nil {
			response.FailWithCode(c, response.UNAUTHORIZED)
			c.Abort()
			return
		}
		secretKey, err := m.userRepo.GetUserSecretKey(preliminaryClaims.UserID)
		if err != nil {
			response.FailWithCode(c, response.UNAUTHORIZED)
			c.Abort()
			return
		}
		claims, err := m.jwtMaker.VerifyToken(accessToken, secretKey)
		if err != nil {
			response.FailWithCode(c, response.UNAUTHORIZED)
			c.Abort()
			return
		}
		//扩展redis加强token验证
		exists, err := m.rdb.Exists(c.Request.Context(), secretKey).Result()
		if err != nil || exists <= 0 {
			response.FailWithCode(c, response.UNAUTHORIZED)
			c.Abort()
			return
		}
		token, err := m.rdb.Get(c.Request.Context(), secretKey).Result()
		if err != nil || token != accessToken {
			response.FailWithCode(c, response.UNAUTHORIZED)
			c.Abort()
			return
		}
		c.Set("userId", claims.UserID)
		ctx := context.WithValue(c.Request.Context(), "userId", claims.UserID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
