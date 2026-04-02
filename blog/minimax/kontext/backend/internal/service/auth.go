package service

import (
	"blog/internal/config"
	"blog/internal/model/dto"
	"blog/internal/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
	config  *config.Config
}

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewAuthService() *AuthService {
	cfg := config.GlobalConfig
	if cfg == nil {
		cfg = &config.Config{
			JWTSecret: "your-secret-key-change-in-production",
			JWTExpire: 72,
		}
	}
	return &AuthService{
		userRepo: repository.NewUserRepository(),
		config:   cfg,
	}
}

func (s *AuthService) Register(req *dto.RegisterRequest) (*repository.User, error) {
	existing, _ := s.userRepo.GetByUsername(req.Username)
	if existing != nil {
		return nil, errors.New("用户名已存在")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("密码加密失败")
	}

	user := &repository.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Nickname: req.Nickname,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.New("创建用户失败")
	}

	return user, nil
}

func (s *AuthService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.userRepo.GetByUsername(req.Username)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	token, expiresAt, err := s.generateToken(user)
	if err != nil {
		return nil, errors.New("生成令牌失败")
	}

	return &dto.LoginResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		UserID:    user.ID,
		Username:  user.Username,
		Nickname:  user.Nickname,
	}, nil
}

func (s *AuthService) generateToken(user *repository.User) (string, int64, error) {
	expireTime := time.Now().Add(time.Duration(s.config.JWTExpire) * time.Hour)
	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.config.JWTSecret))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expireTime.Unix(), nil
}

func (s *AuthService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		return []byte(s.config.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无效的令牌")
}

func (s *AuthService) GetUserByID(id uint) (*repository.User, error) {
	return s.userRepo.GetByID(id)
}
