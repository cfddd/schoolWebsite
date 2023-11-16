package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/ApplySystem"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/test"
)

type ServiceGroup struct {
	SystemServiceGroup      system.ServiceGroup
	ExampleServiceGroup     example.ServiceGroup
	ApplySystemServiceGroup ApplySystem.ServiceGroup
	TestServiceGroup        test.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
