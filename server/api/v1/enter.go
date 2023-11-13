package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/ApplySystem"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
)

type ApiGroup struct {
	SystemApiGroup      system.ApiGroup
	ExampleApiGroup     example.ApiGroup
	ApplySystemApiGroup ApplySystem.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

var MUService = service.ServiceGroupApp.ApplySystemServiceGroup.MaterialUploadService
