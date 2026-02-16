package auth

import (
	"context"
	"errors"
	"time"

	"github.com/JscorpTech/auth/internal/config"
	"github.com/JscorpTech/auth/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthUsecase interface {
	Login(context.Context, string, string) (*User, error)
	Register(context.Context, *User) (*User, error)
	IsExists(context.Context, string) bool
	ValidateToken(string) (*jwt.MapClaims, error)
	AccessToken(*User) string
	RefreshToken(*User) string
	SendOtp(context.Context, string) error
	ValidateOtp(context.Context, string, string) bool
	IsConfirm(context.Context, *User) bool
	GetUserByPhone(context.Context, string) (*User, error)
	Confirm(context.Context, *User)
}

type AuthUsecaseImpl struct {
	repo   AuthRepository
	cfg    *config.Config
	logger *zap.Logger
}

func NewAuthUsecase(repo AuthRepository, cfg *config.Config, logger *zap.Logger) AuthUsecase {
	return &AuthUsecaseImpl{
		repo:   repo,
		cfg:    cfg,
		logger: logger,
	}
}

func (a *AuthUsecaseImpl) IsConfirm(ctx context.Context, user *User) bool {
	return user.ValidatedAT != nil
}

func (a *AuthUsecaseImpl) ValidateToken(token string) (*jwt.MapClaims, error) {
	claims, err := utils.VerifyJWT(token, a.cfg.PublicKey)
	if err != nil {
		return nil, err
	}
	if claims["type"] != "refresh" {
		return nil, ErrInvalidRefreshToken
	}
	return &claims, nil
}

func (a *AuthUsecaseImpl) Login(ctx context.Context, phone string, password string) (*User, error) {
	user, err := a.repo.GetByPhone(ctx, phone)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidCredentions
		}
		return nil, err
	}
	if !a.IsConfirm(ctx, user) {
		return nil, ErrPhoneNumberNotConfirmed
	}

	if res := utils.CheckPasswordHash(password, user.Password); !res {
		return nil, ErrInvalidPassword
	}
	return user, nil
}

func (a *AuthUsecaseImpl) IsExists(ctx context.Context, phone string) bool {
	return a.repo.IsExists(ctx, phone)
}

func (a *AuthUsecaseImpl) Register(ctx context.Context, user *User) (*User, error) {
	userInstance, err := a.repo.GetByPhone(ctx, user.Phone)
	if err == nil && a.IsConfirm(ctx, userInstance) {
		return nil, ErrUserAlreadyExists
	}
	if err != nil {
		userInstance, err = a.repo.Create(ctx, user)
	}
	if err := a.SendOtp(ctx, user.Phone); err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return userInstance, nil
}

func (a *AuthUsecaseImpl) AccessToken(user *User) string {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Minute * time.Duration(a.cfg.AccessExp)).Unix(),
		"type":    "access",
		"role":    user.Role,
	}
	token, err := utils.CreateJWT(claims, a.cfg.PrivateKey)
	if err != nil {
		return ""
	}
	return token
}

func (a *AuthUsecaseImpl) RefreshToken(user *User) string {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Minute * time.Duration(a.cfg.RefreshExp)).Unix(),
		"type":    "refresh",
		"role":    user.Role,
	}
	token, err := utils.CreateJWT(claims, a.cfg.PrivateKey)
	if err != nil {
		return ""
	}
	return token
}

func (a *AuthUsecaseImpl) SendOtp(ctx context.Context, phone string) error {
	code := utils.RandomOtp(6)
	a.logger.Info("New otp", zap.String("otp", code))
	otp, err := a.repo.GetOtpByPhone(ctx, phone)
	if err != nil {
		otp, err = a.repo.CreateOtp(ctx, phone, code)
		if err != nil {
			return err
		}
	} else if time.Since(otp.UpdatedAt) < time.Minute*2 {
		return ErrRateLimit
	}
	if err := a.repo.UpdateOtp(ctx, phone, code); err != nil {
		return err
	}
	return nil
}

func (a *AuthUsecaseImpl) ValidateOtp(ctx context.Context, phone string, otp string) bool {
	otpInstance, err := a.repo.GetOtp(ctx, phone, otp)
	if err != nil {
		a.logger.Info("invalid otp", zap.Error(err))
		return false
	}
	a.repo.DeleteOtp(ctx, otpInstance)
	return true
}

func (a *AuthUsecaseImpl) GetUserByPhone(ctx context.Context, phone string) (*User, error) {
	return a.repo.GetByPhone(ctx, phone)
}

func (a *AuthUsecaseImpl) Confirm(ctx context.Context, user *User) {
	a.repo.Update(ctx, user, map[string]any{
		"validated_at": time.Now(),
	})
}
