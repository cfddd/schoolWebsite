// 自动生成模板PatentAuthorization
package ApplySystem

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// PatentAuthorization 结构体  PatentAuthorization
type PatentAuthorization struct {
	global.GVA_MODEL
	Student_id      string     `json:"student_id" form:"student_id" gorm:"column:student_id;comment:;size:20;"`        //学号
	Student_name    string     `json:"student_name" form:"student_name" gorm:"column:student_name;comment:;size:20;"`  //学号
	Patent_name     string     `json:"patent_name" form:"patent_name" gorm:"column:patent_name;comment:;"`             //专利名称
	Patent_type     string     `json:"patent_type" form:"patent_type" gorm:"column:patent_type;comment:;"`             //专利类别
	Patent_grant_id string     `json:"patent_grant_id" form:"patent_grant_id" gorm:"column:patent_grant_id;comment:;"` //专利授权号
	Approval_time   *time.Time `json:"approval_time" form:"approval_time" gorm:"column:approval_time;comment:;"`       //获批时间
	Is_first_invent *bool      `json:"is_first_invent" form:"is_first_invent" gorm:"column:is_first_invent;comment:;"` //是否第一发明人
}

// TableName PatentAuthorization PatentAuthorization自定义表名 PatentAuthorization
func (PatentAuthorization) TableName() string {
	return "PatentAuthorization"
}
