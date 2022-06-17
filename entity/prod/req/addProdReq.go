package req

import "bingFood-Merchant/entity/prod"

type AddProdReq struct {
    Prod prod.Prod `json:"prod"`

    // 操作人信息
    ShopId int `json:"shopId"`
}
