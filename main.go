package main

import (
	"log"

	"bingFood-Merchant/initialize"
	"bingFood-Merchant/router"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
	r = router.SetupRouter()
	initialize.Viper()
	initialize.MySql()
	initialize.Redis()
}
func main() {
	if err := r.Run(":8089"); err != nil {
		log.Fatal(err)
	}
}
