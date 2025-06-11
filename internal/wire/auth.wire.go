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
	)
	return new(controller.AccountController), nil
}

func InitModuleFriendShip() (*controller.FriendShipController, error) {
	wire.Build(
		reporitory.NewFriendShipRepository,
		reporitory.NewAccountRepository,
		reporitory.NewProfileRepository,
		service.NewFriendShipService,
		controller.NewFriendShipController,
	)
	return new(controller.FriendShipController), nil
}
