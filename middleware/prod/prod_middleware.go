package prod

import (
	"bingFood-Merchant/common/response"
	"bingFood-Merchant/entity/prod"
	"bingFood-Merchant/entity/prod/req"
	"bingFood-Merchant/global"
	"bingFood-Merchant/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

// 新增商品
func AddProdMiddle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var param req.AddProdReq
		ctx.ShouldBindJSON(&param)
		fmt.Println(utils.ToJsonString(param.Prod))

		db := global.MYSQL_DB
		if err := db.Transaction(func(tx *gorm.DB) error {
			err := tx.Create(&param.Prod).Error
			return err
		}); err != nil {
			log.Printf("insert prod failed : %v", err.Error())
			response.FailWithMessage(err.Error(), ctx)
			ctx.Abort()
			return
		}
		response.OkWithMessage("插入商品成功", ctx)
		ctx.Abort()
		return
	}
}

// 获取商品列表
func GetProdList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var param req.ListProdReq
		_ = ctx.ShouldBindJSON(&param)

		fmt.Println(utils.ToJsonString(&param))
		db := global.MYSQL_DB

		var prodList []prod.Prod
		limit := param.PageInfo.PageSize
		offset := param.PageInfo.PageSize * (param.PageInfo.Page - 1)
		var total int64
		db.Where("shop_id = ?", param.ShopId).Count(&total)

		fmt.Println(utils.ToJsonString(param.ShopId))

		if err := db.Limit(limit).Offset(offset).
			Preload("Skus").Preload("Properties").
			Where(&prod.Prod{ShopId: param.ShopId}). // 这样当没传入shopId时零值就不会作为条件了
			Find(&prodList).Error; err != nil {
			response.FailWithMessage(fmt.Sprintf("获取列表失败, err:%v", err.Error()), ctx)
			ctx.Abort()
			return
		}
		response.OkWithDetailed(response.PageResult{
			List:     prodList,
			Total:    total,
			Page:     param.PageInfo.Page,
			PageSize: param.PageInfo.PageSize,
		}, "获取列表成功", ctx)
		ctx.Abort()
		return
	}
}
