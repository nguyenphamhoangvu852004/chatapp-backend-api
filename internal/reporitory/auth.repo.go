package reporitory

import (
	"chapapp-backend-api/global"
	"time"
)

type IAuthRepository interface {
	AddOTP(email string, otp int, expirationTime int64) error
	GetOTP(email string) (storedEmail string, otp int, err error)
	CanSendOTP(email string) (bool, error)
	RemoveOTP(email string) error
}

type AuthRepository struct {
}

// RemoveOTP implements IAuthRepository.
func (u *AuthRepository) RemoveOTP(email string) error {
	key := "otp:" + email
	err := global.Rdb.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetOTP implements IAuthRepository.
func (u *AuthRepository) GetOTP(email string) (storedEmail string, otp int, err error) {
	key := "otp:" + email
	otp, err = global.Rdb.Get(ctx, key).Int()
	if err != nil {
		return "", 0, err
	}
	return email, otp, nil
}

// AddOTP implements IUserAuthRepository.
func (u *AuthRepository) AddOTP(email string, otp int, expirationTime int64) error {
	key := "otp:" + email
	err := global.Rdb.SetEx(ctx, key, otp, time.Duration(expirationTime)).Err()
	if err != nil {
		return err
	}
	return nil
}
func (u *AuthRepository) CanSendOTP(email string) (bool, error) {
	key := "otp:" + email
	exists, err := global.Rdb.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	// Nếu key đã tồn tại → Không được gửi tiếp
	return exists == 0, nil
}

func NewAuthRepository() IAuthRepository {
	return &AuthRepository{}
}
