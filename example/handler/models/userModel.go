package models

type User struct {
	Model
	OpenID   string `gorm:"open_id;type:varchar(255);default:'';not null;unique"`
	UnionID  string `gorm:"union_id;type:varchar(255);default:'';not null"`
	NickName string `gorm:"nick_name;type:varchar(255);default:'';not null"`
	Avatar   string `gorm:"avatar;type:varchar(255);default:'';not null"`
	Gender   int    `gorm:"gender;type:int;default:0;not null"`
	Country  string `gorm:"country;type:varchar(100);default:'';not null"`
	Province string `gorm:"province;type:varchar(100);default:'';not null"`
	City     string `gorm:"city;type:varchar(100);default:'';not null"`
	Mobile   string `gorm:"mobile;type:varchar(15);default:'';not null"`
	Integral uint   `gorm:"integral;type:int;default:0;not null"`
	Status   int    `gorm:"status;type:tinyint(1);not null;default:0;comment:'1:有用户信息'"`
}

//IntegralLog 积分记录
type IntegralLog struct {
	Model
	UID      uint   `gorm:"uid;type:int;not null;default:0;comment:''" json:"-"`
	Integral int    `gorm:"integral;type:int;default:0;not null" json:"integral" json:"integral"`
	Desc     string `gorm:"desc;type:varchar(255);not null;default:0;comment:'描述'" json:"desc,omitempty"`
}

//UserProduct 用户奖品
type UserProduct struct {
	Model
	UID       uint `gorm:"uid;type:int;not null;default:0;comment:''"`
	ProductID uint `gorm:"product_id;type:int;not null;default:0;comment:''" `
	Status    int8 `gorm:"status;type:int;not null;default:0;comment:'0:待核销，1：已核销'" json:"status"`
}

//UserTask 用户任务
type UserTask struct {
	Model
	UID    uint `gorm:"uid;type:int;not null;default:0;comment:''" json:"uid"`
	TaskID uint `gorm:"task_id;type:int;not null;default:0;comment:''" json:"task_id"`
}

//UserSignUp 用户签到
type UserSignUp struct {
	Model
	UID uint `gorm:"uid;type:int;not null;default:0;comment:''" json:"uid"`
}
