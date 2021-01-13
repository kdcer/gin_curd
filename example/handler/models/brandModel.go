package models

//BrandModel model
type Brand struct {
	Model
	KindID     uint   `gorm:"kind_id;type:int;not null;default:0;comment:''" json:"kind_id,omitempty"`
	DistrictID uint   `gorm:"district_id;type:int;not null;default:0;comment:''" json:"district_id,omitempty"`
	FloorID    uint   `gorm:"floor_id;type:int;not null;default:0;comment:''" json:"floor_id,omitempty"`
	Cover      string `gorm:"cover;type:varchar(255);not null;default:0;comment:''" json:"cover,omitempty"`
	Name       string `gorm:"name;type:varchar(100);not null;default:'';comment:''" json:"name,omitempty"`
	SearchName string `gorm:"search_name;type:varchar(100);default:'';not null"json:"search_name,omitempty"`
	Logo       string `gorm:"name;type:varchar(255);default:'';not null" json:"logo,omitempty"`
	Desc       string `gorm:"name;type:text;not null" json:"desc,omitempty"`
	Sort       int    `gorm:"sort;type:int;not null;index;default:0;comment:'';index" json:"sort,omitempty"`
	Status     int8   `gorm:"status;type:tinyint(1);not null;default:1;comment:'1:显示'" json:"status,omitempty"`
}

//BrandKind 分类表
type BrandKind struct {
	Model
	Name   string `gorm:"name;type:varchar(255);default:'';not null" json:"name"`
	Sort   int    `gorm:"sort;type:tinyint(1);default:100;not null;index" json:"sort,omitempty"`
	Status int8   `gorm:"status;type:tinyint(1);not null;default:1;comment:'1:显示'"json:"status"`
}
