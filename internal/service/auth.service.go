package service

import (
	"chapapp-backend-api/internal/dto/auth"
	"chapapp-backend-api/internal/entity"
	exception "chapapp-backend-api/internal/exeption"
	"chapapp-backend-api/internal/reporitory"
	"chapapp-backend-api/internal/utils"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type IAuthService interface {
	Login(data auth.LoginInputDTO) (auth.LoginOutputDTO, error)
	Register(data auth.RegisterInputDTO) (auth.RegisterOutputDTO, error)
	SendOTP(data auth.SendOTPInputDTO) (auth.SendOTPOutputDTO, error)
	VerifyOTP(data auth.VerifyOTPInputDTO) (auth.VerifyOTPOutputDTO, error)
	ResetPassword(data auth.ResetPasswordInputDTO) (auth.ResetPasswordOutputDTO, error)
}

type authService struct {
	accountRepo reporitory.IAccountRepository
	authRepo    reporitory.IAuthRepository
}

// ResetPassword implements IAuthService.
func (authService *authService) ResetPassword(data auth.ResetPasswordInputDTO) (auth.ResetPasswordOutputDTO, error) {
	// tìm cái entity account theo email

	account, err := authService.accountRepo.GetUserByEmail(data.Email)
	if err != nil {
		return auth.ResetPasswordOutputDTO{}, exception.NewCustomError(http.StatusNotFound, "Not found account")
	}

	// so sanh 2 cai passs tu client
	if data.ConfirmPassword != data.Password {
		return auth.ResetPasswordOutputDTO{}, exception.NewCustomError(http.StatusBadRequest, "Password and confirm password do not match")
	}

	// doi qua mat khau moi
	account.Password = data.Password
	account.UpdatedAt = time.Now()

	// luu vao db
	_, err = authService.accountRepo.Update(account)
	if err != nil {
		return auth.ResetPasswordOutputDTO{}, exception.NewCustomError(http.StatusInternalServerError, "Failed to update account")
	}

	outDTO := auth.ResetPasswordOutputDTO{
		Id:      strconv.FormatUint(uint64(account.ID), 10),
		Email:   account.Email,
		Message: "Reset password success",
	}

	return outDTO, nil
}

// Login implements IAuthService.
func (authService *authService) Login(data auth.LoginInputDTO) (auth.LoginOutputDTO, error) {
	// tim theo username
	account, err := authService.accountRepo.GetUserByUsername(data.Username)
	if err != nil {
		return auth.LoginOutputDTO{}, exception.NewCustomError(http.StatusNotFound, "Not found account")
	}

	if account.Password != data.Password {
		return auth.LoginOutputDTO{}, exception.NewCustomError(http.StatusUnauthorized, "Invalid password")
	}

	var accToken, _ = utils.GenerateAccessToken(account.ID, account.Email)
	var refToken, _ = utils.GenerateRefreshToken(account.ID)

	return auth.LoginOutputDTO{
		Id:           strconv.FormatUint(uint64(account.ID), 10),
		AccessToken:  accToken,
		RefreshToken: refToken,
	}, nil

}

func (authService *authService) VerifyOTP(data auth.VerifyOTPInputDTO) (auth.VerifyOTPOutputDTO, error) {

	var otpEmailKey = utils.GetHash(data.Email)
	// tìm coi có cái key trong redis hay không
	email, otp, _ := authService.authRepo.GetOTP(otpEmailKey)
	fmt.Println(email, otp)
	if email != "" && otp == utils.StringToInt(data.OTP) {
		return auth.VerifyOTPOutputDTO{
			Email:   data.Email,
			Message: "Verify OTP success",
		}, nil
	}
	authService.authRepo.RemoveOTP(otpEmailKey)
	return auth.VerifyOTPOutputDTO{
		Email:   data.Email,
		Message: "Verify OTP fail",
	}, errors.New("verify OTP failed")
}

// SendOTP implements IAuthService.
func (authService *authService) SendOTP(data auth.SendOTPInputDTO) (auth.SendOTPOutputDTO, error) {
	// tạo otp
	otp := utils.GenerateSixDigitNumber()
	// mã hoá hash
	strHashed := utils.GetHash(data.Email)

	if _, err := authService.authRepo.CanSendOTP(strHashed); err != nil {
		return auth.SendOTPOutputDTO{
			Email:   data.Email,
			Message: "Please wait for 10 minutes",
		}, errors.New("please wait for 10 minutes")
	}

	if err := authService.authRepo.AddOTP(strHashed, otp, int64(10*time.Minute)); err != nil {
		return auth.SendOTPOutputDTO{
			Email:   data.Email,
			Message: "Send OTP fail",
		}, errors.New("send OTP fail")
	}

	// thực hiện gữi vào email
	err := utils.SendTextEmailOTP([]string{data.Email}, "nguyenphamhoangvu852004@gmail.com", "OTP:"+strconv.Itoa(otp))
	if err != nil {
		return auth.SendOTPOutputDTO{
			Email:   data.Email,
			Message: "Send OTP fail",
		}, errors.New("send OTP fail")
	}
	return auth.SendOTPOutputDTO{
		Email:          data.Email,
		Message:        "Send OTP success",
		ExpirationTime: 10,
	}, nil
}

// Register implements IAuthService.
func (s *authService) Register(input auth.RegisterInputDTO) (auth.RegisterOutputDTO, error) {
	// 1. Check email đã tồn tại chưa
	_, err := s.accountRepo.GetUserByEmail(input.Email)
	if err == nil {
		// Email đã tồn tại, trả về lỗi
		return auth.RegisterOutputDTO{}, exception.NewCustomError(http.StatusConflict, "Email already exists")
	}
	// 2. Kiểm tra password và confirm password
	if input.Password != input.ConfirmPassword {
		return auth.RegisterOutputDTO{}, exception.NewCustomError(http.StatusBadRequest, "Password and confirm password do not match")
	}

	// 3. Tạo entity account mới
	newAccount := entity.Account{
		Email:       input.Email,
		Password:    input.Password, // TODO: mã hóa sau
		Username:    input.Username,
		PhoneNumber: input.PhoneNumber,
	}

	// 4. Lưu vào DB
	createdAccount, err := s.accountRepo.Create(newAccount)
	if err != nil {
		return auth.RegisterOutputDTO{}, exception.NewCustomError(http.StatusInternalServerError, "Failed to create account")
	}

	// 5. Chuẩn bị output DTO
	output := auth.RegisterOutputDTO{
		Id:       strconv.FormatUint(uint64(createdAccount.ID), 10),
		Username: createdAccount.Username,
		Email:    createdAccount.Email,
	}
	return output, nil
}

func NewAuthService(accountRepo reporitory.IAccountRepository, authRepo reporitory.IAuthRepository) IAuthService {
	return &authService{
		accountRepo: accountRepo,
		authRepo:    authRepo,
	}
}
