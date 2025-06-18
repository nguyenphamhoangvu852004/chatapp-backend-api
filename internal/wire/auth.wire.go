//go:build wireinject

package wire

import (
	"chapapp-backend-api/internal/controller"
	"chapapp-backend-api/internal/reporitory"
	"chapapp-backend-api/internal/service"

	"github.com/google/wire"
)

func InitModuleAuth() (*controller.AuthController, error) {
	wire.Build(
		controller.NewAuthController,
		service.NewAuthService,
		reporitory.NewAccountRepository,
		reporitory.NewAuthRepository,
	)
	return new(controller.AuthController), nil
}

func InitModuleProfile() (*controller.ProfileController, error) {
	wire.Build(
		controller.NewProfileController,
		service.NewProfileService,
		reporitory.NewProfileRepository,
	)
	return new(controller.ProfileController), nil
}

func InitModuleAccount() (*controller.AccountController, error) {
	wire.Build(
		controller.NewAccountController,
		service.NewAccountService,
		reporitory.NewAccountRepository,
		reporitory.NewBlockRepository,
	)
	return new(controller.AccountController), nil
}

func InitModuleFriendShip() (*controller.FriendShipController, error) {
	wire.Build(
		reporitory.NewFriendShipRepository,
		reporitory.NewAccountRepository,
		reporitory.NewProfileRepository,
		reporitory.NewConversationRepository,
		reporitory.NewParticiapntRepository,
		reporitory.NewBlockRepository,
		service.NewFriendShipService,
		controller.NewFriendShipController,
	)
	return new(controller.FriendShipController), nil
}

func InitModuleBlock() (*controller.BlockController, error) {
	wire.Build(
		reporitory.NewBlockRepository,
		reporitory.NewFriendShipRepository,
		service.NewBlockService,
		controller.NewBlockController,
	)
	return new(controller.BlockController), nil
}

func InitModuleMessage() (*controller.MessageController, error) {
	wire.Build(
		reporitory.NewMessageRepository,
		service.NewMessageService,
		controller.NewMessageController,
	)
	return new(controller.MessageController), nil
}

func InitModuleConversation() (*controller.ConversationController, error) {
	wire.Build(
		reporitory.NewConversationRepository,
		reporitory.NewParticiapntRepository,
		service.NewConversationService,
		controller.NewConversationController,
	)
	return new(controller.ConversationController), nil
}

func InitModuleBan() (*controller.BanController, error) {
	wire.Build(
		reporitory.NewBanRepository,
		reporitory.NewAccountRepository,
		service.NewBanService,
		controller.NewBanController,
	)
	return new(controller.BanController), nil
}
