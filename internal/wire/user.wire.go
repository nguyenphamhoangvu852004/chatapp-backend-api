//go:build wireinject

package wire

import (
	"chapapp-backend-api/internal/controller"
	"chapapp-backend-api/internal/reporitory"
	"chapapp-backend-api/internal/service"

	"github.com/google/wire"
)

func InitModuleUser() (*controller.UserController, error) {
	wire.Build(controller.NewUserController, service.NewUserService, reporitory.NewUserRepository)
	return new(controller.UserController), nil
}
