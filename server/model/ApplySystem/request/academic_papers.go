package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type AcademicPapersSearch struct {
	ApplySystem.AcademicPapers
	StartCreatedAt    *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt      *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	StartPublish_time *time.Time `json:"startPublish_time" form:"startPublish_time"`
	EndPublish_time   *time.Time `json:"endPublish_time" form:"endPublish_time"`
	request.PageInfo
}
