package ApplySystem

import "time"

type MaterialUploadModel struct {
	ID                 uint             `json:"id" gorm:"primarykey"` // 主键ID
	CreatedAt          time.Time        `json:"created_at"`           // 创建时间
	UpdatedAt          time.Time        `json:"updated_at"`
	Path               string           `json:"path"`
	CompetitionPrizeID uint             // 外键
	CompetitionPrize   CompetitionPrize // 属于
	AcademicPapersID   uint             // 外键
	AcademicPapers     AcademicPapers   // 属于
}
