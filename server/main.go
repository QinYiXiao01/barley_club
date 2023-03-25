package main

import (
	// "net/http"
	"server/common"
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	_ = common.InitDB()

	r := gin.Default()
	r = controllers.CollectRoute(r)
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}

// 程序入口
// func main() {
// 	controllers.Router()
// 	http.ListenAndServe(":8080", nil)
// }
