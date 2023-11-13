// 自动生成模板ScientificResearch
package ApplySystem

import (
	"time"
)

// 科研项目填报 结构体  ScientificResearch
type ScientificResearch struct {
	ID                      uint       `json:"id" gorm:"primarykey"` // 主键ID
	CreatedAt               string     `json:"created_at"`           // 创建时间
	UpdatedAt               string     `json:"updated_at"`
	Student_id              string     `json:"student_id" form:"student_id" gorm:"column:student_id;comment:;size:20;"`                                         //学号
	Student_name            string     `json:"student_name" form:"student_name" gorm:"column:student_name;comment:;size:20;"`                                   //学生姓名
	Scientific_project_name string     `json:"scientific_project_name" form:"scientific_project_name" gorm:"column:scientific_project_name;comment:;size:100;"` //科研项目名称
	Work_id                 string     `json:"work_id" form:"work_id" gorm:"column:work_id;comment:;size:20;"`                                                  //工号
	Project_approval_time   *time.Time `json:"project_approval_time" form:"project_approval_time" gorm:"column:project_approval_time;comment:;"`                //项目获批时间
}

// TableName 科研项目填报 ScientificResearch自定义表名 ScientificResearch
func (ScientificResearch) TableName() string {
	return "ScientificResearch"
}
