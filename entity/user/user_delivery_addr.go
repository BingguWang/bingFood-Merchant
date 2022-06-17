package user

import "time"

/**
    用户配送地址
 */
type UserDeliveryAddr struct {
    UserDeliveryAddrId int
    UserId             string
    Receiver           string // 接收人名称
    ProvinceId         int    // 省id
    Province           string // 省名称
    CityId             int    // 城市id
    City               string // 城市名称
    AreaId             int    // 区id
    Area               string // 区名称
    Detail             string // 详细地址

    CreateAt time.Time // 创建时间
    UpdateAt time.Time // 修改时间
}
