package service

// import (
// 	"chapapp-backend-api/internal/reporitory"
// 	"chapapp-backend-api/internal/utils"
// 	"chapapp-backend-api/pkg/response"
// 	"fmt"
// 	"time"
// )

// type IUserService interface {
// 	Register(email string, purpose string) int
// }

// type userService struct {
// 	userRepo     reporitory.IUserRepository
// 	userAuthRepo reporitory.IUserAuthRepository
// }

// // Register implements IUserService.
// func (us *userService) Register(email string, purpose string) int {

// 	strHashed := utils.GetHash(email)
// 	fmt.Println(strHashed)
// 	// check exist
// 	if us.userRepo.GetUserByEmail(email) {
// 		return response.ErrCodeParamInvalid
// 	}

// 	// new otp
// 	otp := utils.GenerateSixDigitNumber()
// 	if purpose == "otp" {
// 		otp = otp
// 	}

// 	if err := us.userAuthRepo.AddOTP(strHashed, otp, int64(10*time.Minute)); err != nil {
// 		return response.ErrCodeParamInvalid
// 	}
// 	if err := utils.SendTextEmailOTP([]string{email}, utils.AdminReceiver, fmt.Sprintf("OTP: %d", otp)); err != nil {
// 		return response.ErrCodeSendMail
// 	}
// 	return response.ErrCodeSuccess
// }

// func NewUserService(userRepo reporitory.IUserRepository, userAuthRepo reporitory.IUserAuthRepository) IUserService {
// 	return &userService{
// 		userRepo:     userRepo,
// 		userAuthRepo: userAuthRepo,
// 	}
// }
