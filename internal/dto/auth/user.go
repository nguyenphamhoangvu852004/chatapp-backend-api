package auth

type RegisterInputDTO struct {
	Email           string `json:"email"`
	Username        string `json:"username"`
	PhoneNumber     string `json:"phoneNumber"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type RegisterOutputDTO struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
type LoginInputDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginOutputDTO struct {
	Id           string `json:"id"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
