package prod

import (
    "bingFood-Merchant/common/response"
    "bingFood-Merchant/entity/prod/req"
    "bingFood-Merchant/global"
    "bingFood-Merchant/utils"
    "fmt"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "log"
)

func AddProdMiddle() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        var param req.AddProdReq
        ctx.ShouldBindJSON(&param)
        fmt.Println(utils.ToJsonString(param.Prod))

        db := global.MYSQL_DB
        if err := db.Transaction(func(tx *gorm.DB) error {
            err := tx.Omit( "CreateAt","UpdateAt").Create(&param.Prod).Error
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
