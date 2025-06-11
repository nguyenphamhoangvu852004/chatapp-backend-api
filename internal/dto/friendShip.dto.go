package dto

type CreateFriendShipInputDTO struct {
	SenderID   string `json:"senderId" binding:"required"`
	ReceiverID string `json:"receiverId" binding:"required"`
}

type CreateFriendShipOutputDTO struct {
	SenderID   string `json:"senderId" binding:"required"`
	ReceiverID string `json:"receiverId" binding:"required"`
	Status     string `json:"status" binding:"required"`
}

type UpdateFriendShipInputDTO struct {
	SenderID   string `json:"senderId" binding:"required"`
	ReceiverID string `json:"receiverId" binding:"required"`
	Status     string `json:"status" binding:"required"`
}

type UpdateFriendShipOutputDTO struct {
	SenderID   string `json:"senderId" binding:"required"`
	ReceiverID string `json:"receiverId" binding:"required"`
}

type GetListFriendShipInputDTO struct {
	Me     string `form:"me"`
	Status string `form:"status"` // lọc theo trạng thái: PENDING, ACCEPT, REJECT
}

type FriendShipItemDTO struct {
	ID        uint      `json:"id"`
	Friend    FriendDTO `json:"friend"`
	Status    string    `json:"status"`
	CreatedAt string    `json:"createdAt"`
}

type FriendDTO struct {
	ID        uint   `json:"id"`
	Fullname  string `json:"fullname"`
	AvatarURL string `json:"avatarUrl"`
}

type GetListFriendShipOutputDTO struct {
	Me   uint              `json:"me"`
	Data []FriendShipItemDTO `json:"data"` // danh sách mối quan hệ
}


type GetFriendShipOutputDTO struct {
	Me     Sender     `json:"me"`
	Others []Receiver `json:"others"`
}

type Sender struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Receiver struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	ImageURL string `json:"imageUrl"`
	Status   string `json:"status"`
}
