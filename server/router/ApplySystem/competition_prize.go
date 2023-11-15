package ApplySystem

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CompetitionPrizeRouter struct {
}

// InitCompetitionPrizeRouter 初始化 比赛获奖申报 路由信息
func (s *CompetitionPrizeRouter) InitCompetitionPrizeRouter(Router *gin.RouterGroup) {
	//CPRouter := Router.Group("CP").Use(middleware.OperationRecord())
	CPRouter := Router.Group("CP")
	CPRouterWithoutRecord := Router.Group("CP")
	var CPApi = v1.ApiGroupApp.ApplySystemApiGroup.CompetitionPrizeApi
	{
		CPRouter.POST("createCompetitionPrize", CPApi.CreateCompetitionPrize)             // 新建比赛获奖申报
		CPRouter.DELETE("deleteCompetitionPrize", CPApi.DeleteCompetitionPrize)           // 删除比赛获奖申报
		CPRouter.DELETE("deleteCompetitionPrizeByIds", CPApi.DeleteCompetitionPrizeByIds) // 批量删除比赛获奖申报
		CPRouter.PUT("updateCompetitionPrize", CPApi.UpdateCompetitionPrizeStudent)       // 更新比赛获奖申报
	}
	{
		CPRouterWithoutRecord.GET("findCompetitionPrize", CPApi.FindCompetitionPrize)       // 根据ID获取比赛获奖申报
		CPRouterWithoutRecord.GET("getCompetitionPrizeList", CPApi.GetCompetitionPrizeList) // 获取比赛获奖申报列表
	}
}
