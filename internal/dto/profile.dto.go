package dto

type GetProfileDetailOutputDTO struct {
	Id        string `json:"id"`
	FullName  string `json:"fullname"`
	Bio       string `json:"bio"`
	AvatarURL string `json:"avatarUrl"`
	CoverURL  string `json:"coverUrl"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type CreateProfileInputDTO struct {
	AccountID string `json:"accountId" binding:"required"`
	FullName  string `json:"fullname" binding:"required"`
	Bio       string `json:"bio" binding:"required"`
	AvatarURL string `json:"avatarUrl" binding:"required"`
	CoverURL  string `json:"coverUrl" binding:"required"`
}

type CreateProfileOutputDTO struct {
	FullName  string `json:"fullname"`
	Bio       string `json:"bio"`
	AvatarURL string `json:"avatarUrl"`
	CoverURL  string `json:"coverUrl"`
}

type UpdateProfileInputDTO struct {
	ProfileId string `json:"profileId" binding:"required"`
	FullName  string `json:"fullname" binding:"required"`
	Bio       string `json:"bio" binding:"required"`
	AvatarURL string `json:"avatarUrl" binding:"required"`
	CoverURL  string `json:"coverUrl" binding:"required"`
}

type UpdateProfileOutputDTO struct {
	ProfileId string `json:"profileId" binding:"required"`
	FullName  string `json:"fullname" binding:"required"`
	Bio       string `json:"bio" binding:"required"`
	AvatarURL string `json:"avatarUrl" binding:"required"`
	CoverURL  string `json:"coverUrl" binding:"required"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
