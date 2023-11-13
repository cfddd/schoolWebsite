// 自动生成模板CompetitionPrize
package ApplySystem

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"time"
)

// 比赛获奖申报 结构体  CompetitionPrize
type CompetitionPrize struct {
	ID                   uint                  `gorm:"primarykey"` // 主键ID
	CreatedAt            string                // 创建时间
	UpdatedAt            string                // 更新时间                                                                          // 删除时间
	Student_id           string                `json:"student_id" form:"student_id" gorm:"column:student_id;comment:;size:50;"`                    //学号
	Student_name         string                `json:"student_name" form:"student_name" gorm:"column:student_name;comment:;size:20;"`              //学生姓名
	Competition_name     string                `json:"competition_name" form:"competition_name" gorm:"column:competition_name;comment:;size:100;"` //比赛名称
	Award_time           *time.Time            `json:"award_time" form:"award_time" gorm:"column:award_time;comment:;"`                            //获奖时间
	Award_type           string                `json:"award_type" form:"award_type" gorm:"column:award_type;comment:;size:20;"`                    //获奖类别
	Award_level          string                `json:"award_level" form:"award_level" gorm:"column:award_level;comment:;"`                         //获奖等级
	Competition_type     string                `json:"competition_type" form:"competition_type" gorm:"column:competition_type;comment:;"`          //比赛类型
	Description          string                `json:"description" form:"description" gorm:"column:description;comment:;"`                         //说明
	Audit_status         int                   `json:"audit_status" form:"audit_status" gorm:"column:audit_status;comment:;" validate:"oneof=0 1 2"`
	Hint_message         string                `json:"hint_message" form:"hint_message" gorm:"column:hint_message;comment:;"` // 提示消息
	MaterialUploadModels []MaterialUploadModel `json:"material_upload_models" form:"material_upload_models"`                  // 拥有的证明材料
	SysUserID            uint                  `json:"sys_user_id" form:"sys_user_id" gorm:"column:sys_user_id;comment:;"`    // 外键
	SysUser              system.SysUser        `json:"sys_user"`                                                              // 外键
}

// TableName 比赛获奖申报 CompetitionPrize自定义表名 CompetitionPrize
func (CompetitionPrize) TableName() string {
	return "CompetitionPrize"
}
