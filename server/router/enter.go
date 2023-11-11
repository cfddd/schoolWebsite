package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/ApplySystem"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System      system.RouterGroup
	Example     example.RouterGroup
	ApplySystem ApplySystem.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
