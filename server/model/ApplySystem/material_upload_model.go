package ApplySystem

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type MaterialUploadModel struct {
	global.GVA_MODEL
	Path               string           `json:"path"`
	CompetitionPrizeID uint             // 外键
	CompetitionPrize   CompetitionPrize // 属于

}
