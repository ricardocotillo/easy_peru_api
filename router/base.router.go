package router

import (
	"cotillo.tech/inei_api/controllers"
	"cotillo.tech/inei_api/middlewares"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()
	r.Use(middlewares.CorsMiddleware)
	dc := controllers.DepartamentController{DB: db}
	pc := controllers.ProductController{DB: db}
	r.HandleFunc("/departments/", dc.Index).Methods("GET")
	r.HandleFunc("/departments/", dc.Create).Methods("POST")
	r.HandleFunc("/products/", pc.Index).Methods("GET")
	return r
}
