package ApplySystem

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
)

type ApiGroup struct {
	CompetitionPrizeApi
	ScientificResearchApi
	AcademicPapersApi
	PatentAuthorizationApi
}

// ResultDetailView 提交后查询
func ResultDetailView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*request.CustomClaims)

	competitionPrize, _ := CPService.GetCompetitionPrize(claims.ID)

}
