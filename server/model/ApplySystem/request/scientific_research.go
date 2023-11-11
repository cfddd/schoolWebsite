package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ScientificResearchSearch struct {
	ApplySystem.ScientificResearch
	StartCreatedAt             *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt               *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	StartProject_approval_time *time.Time `json:"startProject_approval_time" form:"startProject_approval_time"`
	EndProject_approval_time   *time.Time `json:"endProject_approval_time" form:"endProject_approval_time"`
	request.PageInfo
}
