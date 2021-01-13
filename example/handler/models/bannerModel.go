package models

//BannerModel
type Banner struct {
	Model
	StoreID  int    `gorm:"type;type:tinyint(1);default:0;not null" json:"store_id"` //门店id
	Title    string `gorm:"title;type:varchar(100);not null" json:"title"`
	Cover    string `gorm:"cover;type:varchar(100);not null" json:"cover"`
	AppID    string `gorm:"app_id;type:varchar(255);default:'';not null" json:"app_id"`
	Path     string `gorm:"path;type:varchar(100);default:'';not null" json:"path"`
	JumpType int8   `gorm:"jump_type;type:tinyint(1);not null;default:0;comment:'-1:不跳转,0:内部跳转，1：h5，2：小程序'" json:"jump_type"`
	Sort     uint   `gorm:"sort;type:tinyint(1);default:100;not null" json:"sort"`
	EndAt    int64  `gorm:"end_at;type:int;not null;default:0;comment:''" json:"end_at"`
}
