//go::build wireinject
// +build wireinject

package wire

import (
	"github.com/DoHongKien/go-structure/internal/controller"
	"github.com/DoHongKien/go-structure/internal/repo"
	"github.com/DoHongKien/go-structure/internal/service"
	"github.com/google/wire"
)

func InitAuthRouterHandler() (*controller.AuthController, error) {
	wire.Build(
		repo.NewAuthRepository,
		service.NewAuthService,
		controller.NewAuthController,
	)
	return new(controller.AuthController), nil
}
