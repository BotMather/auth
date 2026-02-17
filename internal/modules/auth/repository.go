package auth

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type AuthRepository interface {
	GetID(context.Context, int64) (*User, error)
	Update(context.Context, *User, map[string]any)
	Create(context.Context, *User) (*User, error)
	IsExists(context.Context, string) bool
	GetByPhone(context.Context, string) (*User, error)
	GetByEmail(context.Context, string) (*User, error)
	GetOtp(context.Context, string, string) (*Otp, error)
	DeleteOtp(context.Context, *Otp)
	CreateOtp(context.Context, string, string) (*Otp, error)
	GetOtpByPhone(context.Context, string) (*Otp, error)
	UpdateOtp(context.Context, string, string) error
	GetOldOtps(context.Context) ([]Otp, error)
}

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{
		db: db,
	}
}

func (a *AuthRepositoryImpl) GetByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	if err := a.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (a *AuthRepositoryImpl) GetID(ctx context.Context, id int64) (*User, error) {
	var user User
	if err := a.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
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

func (a *AuthRepositoryImpl) GetOtpByPhone(ctx context.Context, phone string) (*Otp, error) {
	otp := &Otp{}
	err := a.db.WithContext(ctx).Where("phone = ?", phone).First(otp).Error
	if err != nil {
		return nil, err
	}
	return otp, nil
}

func (a *AuthRepositoryImpl) DeleteOtp(ctx context.Context, otp *Otp) {
	a.db.WithContext(ctx).Unscoped().Delete(otp)
}

func (a *AuthRepositoryImpl) CreateOtp(ctx context.Context, phone string, code string) (*Otp, error) {
	otp := &Otp{
		Phone: phone,
		Code:  code,
	}
	if err := a.db.WithContext(ctx).Create(otp).Error; err != nil {
		return nil, err
	}
	return otp, nil
}

func (a *AuthRepositoryImpl) UpdateOtp(ctx context.Context, phone string, code string) error {
	if err := a.db.WithContext(ctx).Model(&Otp{}).Where("phone = ?", phone).Update("code", code).Error; err != nil {
		return err
	}
	return nil
}

func (a *AuthRepositoryImpl) GetOldOtps(ctx context.Context) ([]Otp, error) {
	var otps []Otp
	twoTimeAgo := time.Now().Add(-2 * time.Minute)
	if err := a.db.WithContext(ctx).Model(&Otp{}).Where("updated_at <= ?", twoTimeAgo).Find(&otps).Error; err != nil {
		return nil, err
	}
	return otps, nil
}
