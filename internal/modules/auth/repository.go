package auth

import (
	"context"

	"gorm.io/gorm"
)

type AuthRepository interface {
	GetID(context.Context, int64) *User
	Update(context.Context, *User, map[string]any)
	Create(context.Context, *User) (*User, error)
	IsExists(context.Context, string) bool
	GetByPhone(context.Context, string) (*User, error)
	GetOtp(context.Context, string, string) (*Otp, error)
	DeleteOtp(context.Context, *Otp)
	CreateOtp(context.Context, string, string) (*Otp, error)
}

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{
		db: db,
	}
}

func (a *AuthRepositoryImpl) GetID(ctx context.Context, id int64) *User {
	var user User
	if err := a.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil
	}
	return &user
}

func (a *AuthRepositoryImpl) GetByPhone(ctx context.Context, phone string) (*User, error) {
	var user User
	if err := a.db.WithContext(ctx).Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (a *AuthRepositoryImpl) IsExists(ctx context.Context, phone string) bool {
	var count int64
	a.db.WithContext(ctx).Model(&User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

func (a *AuthRepositoryImpl) Create(ctx context.Context, user *User) (*User, error) {
	if err := a.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (a *AuthRepositoryImpl) Update(ctx context.Context, user *User, update map[string]any) {
	a.db.WithContext(ctx).Model(user).Updates(update)
}

func (a *AuthRepositoryImpl) GetOtp(ctx context.Context, phone string, otp string) (*Otp, error) {
	var otpInstance Otp
	if err := a.db.WithContext(ctx).Where("phone = ? and code = ?", phone, otp).First(&otpInstance).Error; err != nil {
		return nil, err
	}
	return &otpInstance, nil
}

func (a *AuthRepositoryImpl) DeleteOtp(ctx context.Context, otp *Otp) {
	a.db.WithContext(ctx).Unscoped().Delete(otp)
}

func (a *AuthRepositoryImpl) CreateOtp(ctx context.Context, phone string, otp string) (*Otp, error) {
	otpInstance := &Otp{}
	err := a.db.WithContext(ctx).Where("phone = ?", phone).First(otpInstance).Error
	if err != nil {
		otpInstance = &Otp{Phone: phone, Code: otp}
		if err := a.db.WithContext(ctx).Create(&otpInstance).Error; err != nil {
			return nil, err
		}
	}
	if err := a.db.WithContext(ctx).Model(otpInstance).Update("code", otp).Error; err != nil {
		return nil, err
	}
	return otpInstance, nil
}
