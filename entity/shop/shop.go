package shop

import "time"

type Shop struct {
	ShopId   uint64 `gorm:"primaryKey"` // 店铺id
	ShopName string // 店铺名称
	Intro    string // 店铺简介
	ShopLogo string // 店铺logo(可修改)

	MobileBackgroundPic string // 店铺移动端背景图

	ShopStatus uint8 // 店铺状态(-1:已删除 0: 停业中 1:营业中,因为我这没有写后台系统所以就是直接默认是营业的)

	Type     uint8     // 店铺类型
	CreateAt time.Time `json:"createAt" gorm:"autoCreateTime"` // 创建时间
	UpdateAt time.Time `json:"updateAt" gorm:"autoUpdateTime"` // 最近更新时间
}
