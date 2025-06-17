package socket

import (
	"chapapp-backend-api/internal/dto"
	"chapapp-backend-api/internal/reporitory"
	"chapapp-backend-api/internal/service"
	"fmt"

	"github.com/zishang520/socket.io/v2/socket"
)

// func RegisterHandlers() {
// 	IO.On("connection", func(clients ...any) {
// 		client := clients[0].(*socket.Socket)
// 		fmt.Println("✅ New client connected:", client.Id())

// 		client.On("chat message", func(args ...any) {
// 			data := args[0].(map[string]any)
// 			content := data["content"].(string)
// 			conversationID := uint(data["conversationId"].(float64)) // HTML gửi number
// 			messageType := data["type"].(string)

// 			fmt.Printf("📩 Message received: %s (conversation %d)\n", content, conversationID)

// 			// Test broadcast lại
// 			room := fmt.Sprintf("conversation_%d", conversationID)
// 			IO.To(socket.Room(room)).Emit("chat message", map[string]any{
// 				"senderId":       123, // test cứng
// 				"conversationId": conversationID,
// 				"content":        content,
// 				"type":           messageType,
// 			})
// 		})

// 		client.On("disconnect", func(...any) {
// 			fmt.Println("❌ Client disconnected:", client.Id())
// 		})
// 	})
// }

func RegisterHandlers() {
	IO.On("connection", func(clients ...any) {
		client := clients[0].(*socket.Socket)
		fmt.Println("✅ New client connected:", client.Id())

		client.On("join room", func(args ...any) {
			room := args[0].(string)
			fmt.Println("➡️ Joining room:", room)
			client.Join(socket.Room(room))
		})

		client.On("chat message", func(args ...any) {
			messageRepo := reporitory.NewMessageRepository()
			messageService := service.NewMessageService(messageRepo)

			data := args[0].(map[string]any)
			content := data["content"].(string)
			conversationID := uint(data["conversationId"].(float64))
			senderId := uint(data["senderId"].(float64))
			originFilename := data["originFilename"].(string)
			size := data["size"].(string)
			messageType := data["type"].(string)

			room := fmt.Sprintf("conversation_%d", conversationID)

			fmt.Printf("📩 Message: [%s] in room %s from id %d\n", content, room, senderId)

			_, err := messageService.Create(dto.CreateMessageInputDTO{
				Content:        content,
				ConversationId: fmt.Sprintf("%d", conversationID),
				OriginFilename: &originFilename,
				SenderId:       fmt.Sprintf("%d", senderId),
				Size:           &size,
			})

			if err != nil {
				fmt.Println(err)
				panic(err)
			}

			IO.To(socket.Room(room)).Emit("chat message", map[string]any{
				"senderId":       senderId, // giả lập user
				"conversationId": conversationID,
				"content":        content,
				"originFilename": originFilename,
				"size":           size,
				"type":           messageType,
			})
			fmt.Println("📤 Emitted to room:", room)
		})

		client.On("read message", func(args ...any) {
			data := args[0].(map[string]any)
			conversationID := uint(data["conversationId"].(float64))
			readerID := uint(data["readerId"].(float64))
			room := fmt.Sprintf("conversation_%d", conversationID)

			fmt.Printf("Read message in conversation %d by user %d\n", conversationID, readerID)

			IO.To(socket.Room(room)).Emit("message read", map[string]any{
				"conversationId": conversationID,
				"readerId":       readerID,
			})
		})
		client.On("disconnect", func(...any) {
			fmt.Println("❌ Client disconnected:", client.Id())
		})
	})
}
