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
