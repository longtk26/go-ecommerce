package routers

import (
	"github.com/longtk26/go-ecommerce/internal/routers/manager"
	"github.com/longtk26/go-ecommerce/internal/routers/user"
)

type RouterGroup struct {
	User user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)