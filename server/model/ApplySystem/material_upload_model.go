package ApplySystem

type MaterialUploadModel struct {
	ID                 uint             `json:"id" gorm:"primarykey"` // 主键ID
	CreatedAt          string           `json:"created_at"`           // 创建时间
	UpdatedAt          string           `json:"updated_at"`
	Path               string           `json:"path"`
	CompetitionPrizeID uint             // 外键
	CompetitionPrize   CompetitionPrize // 属于

}
