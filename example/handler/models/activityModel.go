package models

//Kind
type ActivityKind struct {
	Model
	Name   string `gorm:"name;type:varchar(50);not null;default:'';comment:''" json:"name"`
	Sort   int    `gorm:"sort;type:int;not null;default:0;comment:''" json:"sort"`
	Status int    `gorm:"status;type:int;not null;default:1;comment:'1:显示'" json:"status"`
}

//Activity
type Activity struct {
	Model
	KindID        uint   `gorm:"kind;type:int;not null;default:0;comment:''" json:"kind_id"`
	Name          string `gorm:"name;type:varchar(100);not null;default:'';comment:''" json:"name"`
	Cover         string `gorm:"cover;type:varchar(255);not null;default:'';comment:'封面'" json:"cover"`
	Img           string `gorm:"img;type:varchar(255);not null;default:'';comment:''" json:"img,omitempty"`
	Begin         int64  `gorm:"begin;type:int;not null;comment:'活动开始'" json:"begin"`
	End           int64  `gorm:"end;type:int;not null;comment:'活动结束'" json:"end"`
	Address       string `gorm:"address;type:varchar(255);not null;default:'';comment:'活动地址'" json:"address,omitempty"`
	Tips          string `gorm:"tips;type:varchar(255);not null;default:'';comment:'提醒事项'" json:"tips,omitempty"`
	Desc          string `gorm:"desc;type:varchar(255);not null;default:'';comment:''" json:"desc,omitempty"`
	Status        int    `gorm:"status;type:int;not null;default:1;comment:'2：取消'" json:"-"`
	Sort          int    `gorm:"sort;type:int;not null;default:0;comment:''" json:"sort,omitempty"`
	JumpType      int8   `gorm:"jump_type;type:tinyint(1);not null;default:0;comment:'0:内部跳转，1：h5，2：小程序'" json:"jump_type"`
	SignUpType    int8   `gorm:"sign_up_type;type:tinyint(1);not null;default:0;comment:'0:本小程序报名'" json:"sign_up_type"`
	SignUpTotal   int8   `gorm:"sign_up_total;type:int;not null;default:0;comment:'可报名人数'" json:"sign_up_total"`
	OverSingUpNum int8   `gorm:"over_sing_up_num;type:int;not null;default:0;comment:'剩余人数'" json:"over_sing_up_num"`
	Url           string `gorm:"url;type:varchar(255);not null;default:'';comment:'跳转地址'" json:"url"`
	NeedIntegral  uint   `gorm:"need_integral;type:int;not null;default:0;comment:'消耗积分'" json:"need_integral"`
	AppID         string `gorm:"app_id;type:int;not null;default:0;comment:''" json:"app_id"`
}

func (Activity) TableName() string {
	return "activitys"
}

type ActivityUser struct {
	Model
	UID        uint   `gorm:"uid;type:int;not null;default:0;comment:'用户'"`
	ActivityID uint   `gorm:"activity_id;type:int;not null;default:0;comment:''"`
	Mobile     string `gorm:"mobile;type:varchar(15);not null;default:'';comment:''"`
	Status     int    `gorm:"status;type:tinyint(1);not null;default:0;comment:'1:参加'"`
}
