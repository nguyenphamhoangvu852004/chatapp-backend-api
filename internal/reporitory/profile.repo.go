package reporitory

type IProfileRepository interface {
	AddOTP(email string, otp int, expirationTime int64) error
	GetOTP(email string) (storedEmail string, otp int, err error)
	CanSendOTP(email string) (bool, error)
	RemoveOTP(email string) error
}
