// 自动生成模板AcademicPapers
package ApplySystem

import (
	"time"
)

// 学术论文 结构体  AcademicPapers
type AcademicPapers struct {
	ID               uint       `json:"id" gorm:"primarykey"` // 主键ID
	CreatedAt        string     `json:"created_at"`           // 创建时间
	UpdatedAt        string     `json:"updated_at"`
	Student_id       string     `json:"student_id" form:"student_id" gorm:"column:student_id;comment:;size:20;"`           //学号
	Student_name     string     `json:"student_name" form:"student_name" gorm:"column:student_name;comment:;size:20;"`     //学生姓名
	Papers_name      string     `json:"papers_name" form:"papers_name" gorm:"column:papers_name;comment:;"`                //论文名称
	Publish_at       string     `json:"publish_at" form:"publish_at" gorm:"column:publish_at;comment:;"`                   //发表期刊
	Publish_time     *time.Time `json:"publish_time" form:"publish_time" gorm:"column:publish_time;comment:;"`             //发表时间
	Inclusion_status string     `json:"inclusion_status" form:"inclusion_status" gorm:"column:inclusion_status;comment:;"` //收录情况
}

// TableName 学术论文 AcademicPapers自定义表名 AcademicPapers
func (AcademicPapers) TableName() string {
	return "AcademicPapers"
}
