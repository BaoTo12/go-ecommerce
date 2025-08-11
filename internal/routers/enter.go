package routers

import (
	"github.com/BaoTo12/go-ecommerce/internal/routers/admin"
	"github.com/BaoTo12/go-ecommerce/internal/routers/user"
)

type RouterGroup struct {
	User  user.UserRouterGroup
	Admin admin.AdminRouterGroup
}
