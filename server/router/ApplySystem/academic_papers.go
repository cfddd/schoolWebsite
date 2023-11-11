package ApplySystem

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AcademicPapersRouter struct {
}

// InitAcademicPapersRouter 初始化 学术论文 路由信息
func (s *AcademicPapersRouter) InitAcademicPapersRouter(Router *gin.RouterGroup) {
	APRouter := Router.Group("AP").Use(middleware.OperationRecord())
	APRouterWithoutRecord := Router.Group("AP")
	var APApi = v1.ApiGroupApp.ApplySystemApiGroup.AcademicPapersApi
	{
		APRouter.POST("createAcademicPapers", APApi.CreateAcademicPapers)             // 新建学术论文
		APRouter.DELETE("deleteAcademicPapers", APApi.DeleteAcademicPapers)           // 删除学术论文
		APRouter.DELETE("deleteAcademicPapersByIds", APApi.DeleteAcademicPapersByIds) // 批量删除学术论文
		APRouter.PUT("updateAcademicPapers", APApi.UpdateAcademicPapers)              // 更新学术论文
	}
	{
		APRouterWithoutRecord.GET("findAcademicPapers", APApi.FindAcademicPapers)       // 根据ID获取学术论文
		APRouterWithoutRecord.GET("getAcademicPapersList", APApi.GetAcademicPapersList) // 获取学术论文列表
	}
}
