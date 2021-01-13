package models

type Product struct {
	Model
	Name     string `gorm:"name;type:varchar(255);not null;default:'';comment:''" json:"name,omitempty"`
	Cover    string `gorm:"cover;type:varchar(255);not null;default:'';comment:'封面'" json:"cover,omitempty"`
	Img      string `gorm:"img;type:varchar(255);not null;default:'';comment:''" json:"img,omitempty"`
	Integral uint   `gorm:"integral;type:int;not null;default:0;comment:''" json:"integral,omitempty"`
	Desc     string `gorm:"desc;type:varchar(500);not null;default:'';comment:''" json:"desc,omitempty"`
	Total    uint   `gorm:"total;type:int;not null;default:0;comment:'总数" json:"total,omitempty"`
	Overage  uint   `gorm:"overage;type:int;not null;default:0;comment:'剩余'" json:"overage,omitempty"`
	Begin    int64  `gorm:"begin;type:int;not null;default:0;comment:'核销开始'" json:"take_begin,omitempty"`
	End      int64  `gorm:"end;type:int;not null;default:0;comment:'核销结束'" json:"take_end,omitempty"`
	Address  string `gorm:"address;type:varchar(255);not null;default:'';comment:''" json:"address,omitempty"`
	Status   int8   `gorm:"status;type:int;not null;default:1;comment:'1:正常，'" json:"status,omitempty"`
}
