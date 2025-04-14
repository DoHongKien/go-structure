//go::build wireinject
// +build wireinject

package wire

import (
	"github.com/DoHongKien/go-structure/internal/controller"
	"github.com/DoHongKien/go-structure/internal/repo"
	"github.com/DoHongKien/go-structure/internal/service"
	"github.com/google/wire"
)

func InitOrderRouterHandler() (*controller.OrderController, error) {

	wire.Build(
		repo.NewOrderRepository,
		service.NewOrderService,
		controller.NewOrderController,
	)

	return new(controller.OrderController), nil
}