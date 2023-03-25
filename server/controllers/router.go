package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

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
