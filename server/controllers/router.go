package controllers

import (
	"server/controllers/api"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", api.Register)
	r.POST("/api/auth/login", api.Login)
	r.GET("api/auth/info", middleware.AuthMiddleware(), api.Info)
	return r
}

/*
func Router() {
	router := mux.NewRouter()
	staticHandler := http.FileServer(http.Dir("./template"))
	// 注册路由
	http.Handle("/", staticHandler)
	router.Handle("/diary", staticHandler)
	router.Handle("/wiki", staticHandler)
	router.Handle("/find", staticHandler)
	router.Handle("/my", staticHandler)
}
*/
