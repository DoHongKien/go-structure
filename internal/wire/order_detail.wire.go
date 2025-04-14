// //go::build wireinject
// +build wireinject

package wire

import (
	"github.com/DoHongKien/go-structure/internal/controller"
	"github.com/DoHongKien/go-structure/internal/repo"
	"github.com/DoHongKien/go-structure/internal/service"
	"github.com/google/wire"
)

func InitOrderDetailRouterHandler() (*controller.OrderDetailController, error) {

	wire.Build(
		repo.NewOrderDetailRepository,
		service.NewOrderDetailService,
		controller.NewOrderDetailController,
	)

	return new(controller.OrderDetailController), nil
}