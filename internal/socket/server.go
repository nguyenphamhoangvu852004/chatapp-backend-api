package socket

import (

	"github.com/zishang520/engine.io/v2/types"
	"github.com/zishang520/socket.io/v2/socket"
)

var IO *socket.Server

func InitSocketServer() *types.HttpServer {
	httpServer := types.CreateServer(nil)
	IO = socket.NewServer(httpServer, nil)

		RegisterHandlers()

	return httpServer
}

// Dummy ParseToken function and Claims struct for compilation
type Claims struct {
	UserID string
}

func ParseToken(token string) (*Claims, error) {
	// TODO: Replace with actual JWT parsing logic
	return &Claims{UserID: "dummyUserID"}, nil
}
