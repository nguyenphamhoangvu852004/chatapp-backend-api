package auth

type SendOTPInputDTO struct {
	Email string `json:"email"`
}

type SendOTPOutputDTO struct {
	Email          string `json:"email"`
	Message        string `json:"message"`
	ExpirationTime int    `json:"expirationTime"`
}

type VerifyOTPInputDTO struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

type VerifyOTPOutputDTO struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}
