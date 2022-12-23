package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()                      // gin router初始化
	router.GET("/", func(context *gin.Context) { // 建立一開始的載入點，測試gin router 是否正常運作
		context.JSON(http.StatusOK, gin.H{"message": "Hello World!!"})
	})

	err := router.Run(":8088") // 專案運作的port號，打開瀏覽器：http://localhost:8088/
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}
