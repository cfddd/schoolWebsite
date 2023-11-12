package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CompetitionPrizeSearch struct{
    ApplySystem.CompetitionPrize
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    StartAward_time  *time.Time  `json:"startAward_time" form:"startAward_time"`
    EndAward_time  *time.Time  `json:"endAward_time" form:"endAward_time"`
    request.PageInfo
}
