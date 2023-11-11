package ApplySystem

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PatentAuthorizationRouter struct {
}

// InitPatentAuthorizationRouter 初始化 PatentAuthorization 路由信息
func (s *PatentAuthorizationRouter) InitPatentAuthorizationRouter(Router *gin.RouterGroup) {
	PARouter := Router.Group("PA").Use(middleware.OperationRecord())
	PARouterWithoutRecord := Router.Group("PA")
	var PAApi = v1.ApiGroupApp.ApplySystemApiGroup.PatentAuthorizationApi
	{
		PARouter.POST("createPatentAuthorization", PAApi.CreatePatentAuthorization)             // 新建PatentAuthorization
		PARouter.DELETE("deletePatentAuthorization", PAApi.DeletePatentAuthorization)           // 删除PatentAuthorization
		PARouter.DELETE("deletePatentAuthorizationByIds", PAApi.DeletePatentAuthorizationByIds) // 批量删除PatentAuthorization
		PARouter.PUT("updatePatentAuthorization", PAApi.UpdatePatentAuthorization)              // 更新PatentAuthorization
	}
	{
		PARouterWithoutRecord.GET("findPatentAuthorization", PAApi.FindPatentAuthorization)       // 根据ID获取PatentAuthorization
		PARouterWithoutRecord.GET("getPatentAuthorizationList", PAApi.GetPatentAuthorizationList) // 获取PatentAuthorization列表
	}
}
