package socket

import "fmt"

var onlineUsers = make(map[uint]map[string]struct{}) // userId → set(socketId)

func AddOnlineUser(accountId uint, socketId string) {
	if onlineUsers[accountId] == nil {
		onlineUsers[accountId] = make(map[string]struct{})
	}
	onlineUsers[accountId][socketId] = struct{}{}
	fmt.Printf("🟢 AddOnlineUser: %d -> %s\n", accountId, socketId)
}

func RemoveOnlineUserBySocket(socketId string) {
	for id, socketSet := range onlineUsers {
		if _, exists := socketSet[socketId]; exists {
			delete(socketSet, socketId)
			if len(socketSet) == 0 {
				delete(onlineUsers, id)
			}
			break
		}
	}
}

func GetOnlineUserIDs() []uint {
	ids := make([]uint, 0)
	for id := range onlineUsers {
		ids = append(ids, id)
	}
	return ids
}
