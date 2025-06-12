package dto

type (
	CreateMessageInputDTO struct {
		SenderId       string `json:"senderId"`
		ConversationId string `json:"conversationId"`
		Content        string `json:"content"`
	}
	CreateMessageOutputDTO struct {
		MessageId string `json:"messageId"`
		Content   string `json:"content"`
	}
)

type (
	GetListMessageInputDTO struct {
		Me             string `form:"me"`
		ConversationId string `form:"conversationId"`
	}
	GetListMessageOutputDTO struct {
		ConversationId string    `json:"conversationId"`
		Messages       []Message `json:"messages"`
	}
	Message struct {
		ID        string `json:"id"`
		SenderId  string `json:"senderId"`
		Content   string `json:"content"`
		Type      string `json:"type"`
		CreatedAt string `json:"createdAt"`
	}
)
