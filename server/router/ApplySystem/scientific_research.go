package ApplySystem

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ScientificResearchRouter struct {
}

// InitScientificResearchRouter 初始化 科研项目填报 路由信息
func (s *ScientificResearchRouter) InitScientificResearchRouter(Router *gin.RouterGroup) {
	SRRouter := Router.Group("SR").Use(middleware.OperationRecord())
	SRRouterWithoutRecord := Router.Group("SR")
	var SRApi = v1.ApiGroupApp.ApplySystemApiGroup.ScientificResearchApi
	{
		SRRouter.POST("createScientificResearch", SRApi.CreateScientificResearch)             // 新建科研项目填报
		SRRouter.DELETE("deleteScientificResearch", SRApi.DeleteScientificResearch)           // 删除科研项目填报
		SRRouter.DELETE("deleteScientificResearchByIds", SRApi.DeleteScientificResearchByIds) // 批量删除科研项目填报
		SRRouter.PUT("updateScientificResearch", SRApi.UpdateScientificResearch)              // 更新科研项目填报
	}
	{
		SRRouterWithoutRecord.GET("findScientificResearch", SRApi.FindScientificResearch)       // 根据ID获取科研项目填报
		SRRouterWithoutRecord.GET("getScientificResearchList", SRApi.GetScientificResearchList) // 获取科研项目填报列表
	}
}
