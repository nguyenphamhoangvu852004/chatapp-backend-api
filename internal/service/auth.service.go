package service

import (
	"chapapp-backend-api/internal/dto"
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
	Login(data dto.LoginInputDTO) (dto.LoginOutputDTO, error)
	Register(data dto.RegisterInputDTO) (dto.RegisterOutputDTO, error)
	SendOTP(data dto.SendOTPInputDTO) (dto.SendOTPOutputDTO, error)
	VerifyOTP(data dto.VerifyOTPInputDTO) (dto.VerifyOTPOutputDTO, error)
	ResetPassword(data dto.ResetPasswordInputDTO) (dto.ResetPasswordOutputDTO, error)
}

type dtoService struct {
	accountRepo reporitory.IAccountRepository
	dtoRepo     reporitory.IAuthRepository
}

// ResetPassword implements IAuthService.
func (dtoService *dtoService) ResetPassword(data dto.ResetPasswordInputDTO) (dto.ResetPasswordOutputDTO, error) {
	// tìm cái entity account theo email

	account, err := dtoService.accountRepo.GetUserByEmail(data.Email)
	if err != nil {
		return dto.ResetPasswordOutputDTO{}, exception.NewCustomError(http.StatusNotFound, "Not found account")
	}

	// so sanh 2 cai passs tu client
	if data.ConfirmPassword != data.Password {
		return dto.ResetPasswordOutputDTO{}, exception.NewCustomError(http.StatusBadRequest, "Password and confirm password do not match")
	}

	// doi qua mat khau moi
	account.Password = data.Password
	account.UpdatedAt = time.Now()

	// luu vao db
	_, err = dtoService.accountRepo.Update(account)
	if err != nil {
		return dto.ResetPasswordOutputDTO{}, exception.NewCustomError(http.StatusInternalServerError, "Failed to update account")
	}

	outDTO := dto.ResetPasswordOutputDTO{
		Id:      strconv.FormatUint(uint64(account.ID), 10),
		Email:   account.Email,
		Message: "Reset password success",
	}

	return outDTO, nil
}

// Login implements IAuthService.
func (dtoService *dtoService) Login(data dto.LoginInputDTO) (dto.LoginOutputDTO, error) {
	// tim theo username
	account, err := dtoService.accountRepo.GetUserByUsername(data.Username)
	if err != nil {
		return dto.LoginOutputDTO{}, exception.NewCustomError(http.StatusNotFound, "Not found account")
	}

	if account.IsBanned {
		return dto.LoginOutputDTO{}, exception.NewCustomError(http.StatusUnauthorized, "Account is banned")
	}

	// so sanh mat khau tu client voi mat khau trong db

	if !utils.CheckPassword(data.Password, account.Password) {
		return dto.LoginOutputDTO{}, exception.NewCustomError(http.StatusUnauthorized, "Invalid password")
	}

	var roles []string
	for _, r := range account.Roles {
		roles = append(roles, r.Rolename) 
	}

	var accToken, _ = utils.GenerateAccessToken(account.ID, account.Email, roles)
	var refToken, _ = utils.GenerateRefreshToken(account.ID)

	return dto.LoginOutputDTO{
		Id:           strconv.FormatUint(uint64(account.ID), 10),
		AccessToken:  accToken,
		RefreshToken: refToken,
	}, nil

}

func (dtoService *dtoService) VerifyOTP(data dto.VerifyOTPInputDTO) (dto.VerifyOTPOutputDTO, error) {

	var otpEmailKey = utils.GetHash(data.Email)
	// tìm coi có cái key trong redis hay không
	email, otp, _ := dtoService.dtoRepo.GetOTP(otpEmailKey)
	fmt.Println(email, otp)
	if email != "" && otp == utils.StringToInt(data.OTP) {
		return dto.VerifyOTPOutputDTO{
			Email:   data.Email,
			Message: "Verify OTP success",
		}, nil
	}
	dtoService.dtoRepo.RemoveOTP(otpEmailKey)
	return dto.VerifyOTPOutputDTO{
		Email:   data.Email,
		Message: "Verify OTP fail",
	}, errors.New("verify OTP failed")
}

// SendOTP implements IAuthService.
func (dtoService *dtoService) SendOTP(data dto.SendOTPInputDTO) (dto.SendOTPOutputDTO, error) {
	// tạo otp
	otp := utils.GenerateSixDigitNumber()
	// mã hoá hash
	strHashed := utils.GetHash(data.Email)

	if _, err := dtoService.dtoRepo.CanSendOTP(strHashed); err != nil {
		return dto.SendOTPOutputDTO{
			Email:   data.Email,
			Message: "Please wait for 10 minutes",
		}, errors.New("please wait for 10 minutes")
	}

	if err := dtoService.dtoRepo.AddOTP(strHashed, otp, int64(10*time.Minute)); err != nil {
		return dto.SendOTPOutputDTO{
			Email:   data.Email,
			Message: "Send OTP fail",
		}, errors.New("send OTP fail")
	}

	// thực hiện gữi vào email
	err := utils.SendTextEmailOTP([]string{data.Email}, "nguyenphamhoangvu852004@gmail.com", "OTP:"+strconv.Itoa(otp))
	if err != nil {
		return dto.SendOTPOutputDTO{
			Email:   data.Email,
			Message: "Send OTP fail",
		}, errors.New("send OTP fail")
	}
	return dto.SendOTPOutputDTO{
		Email:          data.Email,
		Message:        "Send OTP success",
		ExpirationTime: 10,
	}, nil
}

// Register implements IAuthService.
func (s *dtoService) Register(input dto.RegisterInputDTO) (dto.RegisterOutputDTO, error) {
	// 1. Check email đã tồn tại chưa
	_, err := s.accountRepo.GetUserByEmail(input.Email)
	if err == nil {
		// Email đã tồn tại, trả về lỗi
		return dto.RegisterOutputDTO{}, exception.NewCustomError(http.StatusConflict, "Email already exists")
	}
	// 2. Kiểm tra password và confirm password
	if input.Password != input.ConfirmPassword {
		return dto.RegisterOutputDTO{}, exception.NewCustomError(http.StatusBadRequest, "Password and confirm password do not match")
	}
	// ma hoa mat khau
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return dto.RegisterOutputDTO{}, exception.NewCustomError(http.StatusInternalServerError, "Failed to hash password")
	}

	// 3. Tạo entity account mới
	newAccount := entity.Account{
		Email:       input.Email,
		Password:    hashedPassword,
		Username:    input.Username,
		PhoneNumber: input.PhoneNumber,
		Profile: &entity.Profile{
			FullName:  "",
			Bio:       "",
			AvatarURL: "",
			CoverURL:  "",
		},
	}

	// 4. Lưu vào DB
	createdAccount, err := s.accountRepo.Create(newAccount)
	if err != nil {
		return dto.RegisterOutputDTO{}, exception.NewCustomError(http.StatusInternalServerError, "Failed to create account")
	}

	// 5. Chuẩn bị output DTO
	output := dto.RegisterOutputDTO{
		Id:       strconv.FormatUint(uint64(createdAccount.ID), 10),
		Username: createdAccount.Username,
		Email:    createdAccount.Email,
	}
	return output, nil
}

func NewAuthService(accountRepo reporitory.IAccountRepository, dtoRepo reporitory.IAuthRepository) IAuthService {
	return &dtoService{
		accountRepo: accountRepo,
		dtoRepo:     dtoRepo,
	}
}
