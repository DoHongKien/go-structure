package routers

import (
	"github.com/DoHongKien/go-structure/internal/routers/user"
)

type RouterGroup struct {
	User user.UserRouterGroup
}

var RouterGroupApp = new(RouterGroup)
