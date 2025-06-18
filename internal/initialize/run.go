package initialize

import (
	"chapapp-backend-api/global"
	"chapapp-backend-api/internal/socket"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func Run() {
	LoadConfig()
	InitLogger()
	global.Logger.Info("Load Config Success")
	InitMysql()
	InitRedis()

	go func() {
		InitRouter()
	}()

	func() {
		socketServer := socket.InitSocketServer()
		socketServer.Listen(":"+strconv.Itoa(global.Config.Server.SocketPort), nil)
	}()

	// Graceful shutdown
	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	<-exit
}
