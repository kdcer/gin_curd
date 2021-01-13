package models

type AdminModel struct {
	Model
	Account  string `gorm:"account"`
	Password string `gorm:"password"`
	RealName string `gorm:"real_name"`
}
