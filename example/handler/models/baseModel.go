package models

type Ad struct {
	Model
	Img      string `gorm:"img;type:varchar(255);not null;default:'';comment:''" json:"img"`
	Path     string `gorm:"path;type:varchar(255);not null;default:'';comment:''" json:"path"`
	Position int8   `gorm:"position;type:tinyint(1);not null;default:0;comment:'1：开屏'" json:"-"`
	Sort     int8   `gorm:"sort;type:tinyint(1);not null;default:0;comment:''" json:"-"`
}

//District 区
type District struct {
	Model
	Name   string `gorm:"name;type:varchar(100);not null;default:'';comment:''" json:"name"`
	Sort   int    `gorm:"sort;type:int;not null;default:0;comment:''" json:"sort"`
	Status int8   `gorm:"status;type:int;not null;default:1;comment:'1:显示'" json:"status"`
}

//Floor 楼层
type Floor struct {
	Model
	DistrictID uint   `gorm:"district_id;type:int;not null;default:0;comment:''" json:"district_id"`
	Name       string `gorm:"name;type:varchar(100);not null;default:'';comment:'楼层名'" json:"name"`
	Sort       int    `gorm:"sort;type:int;not null;default:0;comment:''" json:"sort"`
	Status     int8   `gorm:"status;type:int;not null;default:1;comment:'1:显示'" json:"status"`
}
