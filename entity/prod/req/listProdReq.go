package req

import "bingFood-Merchant/common/request"

type ListProdReq struct {
	PageInfo request.PageInfo `json:"pageInfo"`

	ShopId int `json:"shopId"`
}
