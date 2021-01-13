package models

type Task struct {
	Model
	Name     string `gorm:"name;type:varchar(50);not null;default:'';comment:''" json:"name"`
	Integral int    `gorm:"integral;type:int;not null;default:0;comment:''" json:"integral"`
	Icon     string `gorm:"icon;type:varchar(255);not null;default:'';comment:'图标'" json:"icon"`
	Sort     int    `gorm:"sort;type:int;not null;default:0;comment:''" json:"sort,omitempty"`
	Type     int    `gorm:"type;type:int;not null;default:0;comment:'0:阅读，1：扫码，2：分享'" json:"type"`
}
