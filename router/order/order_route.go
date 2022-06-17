package order

import (
    "bingFood-Merchant/middleware"
    "github.com/gin-gonic/gin"
)

func OrderRouter(r *gin.Engine) {
    group := r.Group("/order")
    group.POST("/add", middleware.AddOrderMiddleware())
}