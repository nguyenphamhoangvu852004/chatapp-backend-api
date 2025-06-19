package socket

import (
	"chapapp-backend-api/internal/dto"
	"chapapp-backend-api/internal/reporitory"
	"chapapp-backend-api/internal/service"
	"fmt"

	"github.com/zishang520/socket.io/v2/socket"
)

func RegisterHandlers() {
	IO.On("connection", func(clients ...any) {
		client := clients[0].(*socket.Socket)
		fmt.Println("✅ New client connected:", client.Id())
		client.On("identify", func(args ...any) {
			if len(args) == 0 || args[0] == nil {
				return
			}
			floatId, ok := args[0].(float64)
			if !ok {
				return
			}
			accountId := uint(floatId)
			AddOnlineUser(accountId, string(client.Id())) // Không xoá map
			IO.Emit("online users", GetOnlineUserIDs())   // Gửi tất cả ID đang online
		})
		// client.On("disconnect", func(args ...any) {
		// 	fmt.Println("Client disconnected:", client.Id())
		// 	RemoveOnlineUserBySocket(string(client.Id()))
		// 	IO.Emit("online users", GetOnlineUserIDs())
		// })
		client.On("join room", func(args ...any) {
			room := args[0].(string)
			fmt.Println("➡️ Joining room:", room)
			client.Join(socket.Room(room))
		})

		client.On("chat message", func(args ...any) {
			messageRepo := reporitory.NewMessageRepository()
			accountRepo := reporitory.NewAccountRepository()
			conversationRepo := reporitory.NewConversationRepository()
			messageService := service.NewMessageService(messageRepo, accountRepo, conversationRepo)

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
			RemoveOnlineUserBySocket(string(client.Id()))
			IO.Emit("online users", GetOnlineUserIDs())
		})
		client.On("get online users", func(...any) {
			fmt.Println("📨 Client yêu cầu danh sách người online")
			client.Emit("online users", GetOnlineUserIDs())
		})
	})
}
