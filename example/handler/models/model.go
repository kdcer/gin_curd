package models

import (
	//"github.com/liujunren93/curl_gin/db"
	"gorm.io/gorm"
	"time"
)

func init() {
	initMigrate()
}

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
var ModelList []interface{}
// 迁移数据
func initMigrate() {
	ModelList = []interface{}{
		&Banner{},
		&User{},
		&Brand{},
		&BrandKind{},
		&Ad{},
		&Activity{},
		&ActivityUser{},
		&ActivityKind{},
		&District{},
		&Floor{},
		&IntegralLog{},
		&UserProduct{},
		&Task{},
		&UserTask{},
		&UserSignUp{},
		&Product{},

	}

	//db.DB.AutoMigrate(ModelList...)
}

