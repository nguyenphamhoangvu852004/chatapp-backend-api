// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"chapapp-backend-api/internal/controller"
	"chapapp-backend-api/internal/reporitory"
	"chapapp-backend-api/internal/service"
)

// Injectors from auth.wire.go:

func InitModuleAuth() (*controller.AuthController, error) {
	iAccountRepository := reporitory.NewAccountRepository()
	iAuthRepository := reporitory.NewAuthRepository()
	iAuthService := service.NewAuthService(iAccountRepository, iAuthRepository)
	authController := controller.NewAuthController(iAuthService)
	return authController, nil
}

func InitModuleProfile() (*controller.ProfileController, error) {
	iProfileRepository := reporitory.NewProfileRepository()
	iProfileService := service.NewProfileService(iProfileRepository)
	profileController := controller.NewProfileController(iProfileService)
	return profileController, nil
}

func InitModuleAccount() (*controller.AccountController, error) {
	iAccountRepository := reporitory.NewAccountRepository()
	iBlockRepository := reporitory.NewBlockRepository()
	iAccountService := service.NewAccountService(iAccountRepository, iBlockRepository)
	accountController := controller.NewAccountController(iAccountService)
	return accountController, nil
}

func InitModuleFriendShip() (*controller.FriendShipController, error) {
	iFriendShipRepository := reporitory.NewFriendShipRepository()
	iAccountRepository := reporitory.NewAccountRepository()
	iProfileRepository := reporitory.NewProfileRepository()
	iParticipantRepository := reporitory.NewParticiapntRepository()
	iConversationRepository := reporitory.NewConversationRepository()
	iBlockRepository := reporitory.NewBlockRepository()
	iFriendShipService := service.NewFriendShipService(iFriendShipRepository, iAccountRepository, iProfileRepository, iParticipantRepository, iConversationRepository, iBlockRepository)
	friendShipController := controller.NewFriendShipController(iFriendShipService)
	return friendShipController, nil
}

func InitModuleBlock() (*controller.BlockController, error) {
	iBlockRepository := reporitory.NewBlockRepository()
	iFriendShipRepository := reporitory.NewFriendShipRepository()
	iBlockService := service.NewBlockService(iBlockRepository, iFriendShipRepository)
	blockController := controller.NewBlockController(iBlockService)
	return blockController, nil
}

func InitModuleMessage() (*controller.MessageController, error) {
	iMessageRepository := reporitory.NewMessageRepository()
	iAccountRepository := reporitory.NewAccountRepository()
	iConversationRepository := reporitory.NewConversationRepository()
	iMessageService := service.NewMessageService(iMessageRepository, iAccountRepository, iConversationRepository)
	messageController := controller.NewMessageController(iMessageService)
	return messageController, nil
}

func InitModuleConversation() (*controller.ConversationController, error) {
	iConversationRepository := reporitory.NewConversationRepository()
	iParticipantRepository := reporitory.NewParticiapntRepository()
	iConversationSerivce := service.NewConversationService(iConversationRepository, iParticipantRepository)
	conversationController := controller.NewConversationController(iConversationSerivce)
	return conversationController, nil
}

func InitModuleBan() (*controller.BanController, error) {
	iBanRepository := reporitory.NewBanRepository()
	iAccountRepository := reporitory.NewAccountRepository()
	iBanService := service.NewBanService(iBanRepository, iAccountRepository)
	banController := controller.NewBanController(iBanService)
	return banController, nil
}
