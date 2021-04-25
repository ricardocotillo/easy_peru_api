package router

import (
	"cotillo.tech/inei_api/controllers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()
	dc := controllers.DepartamentController{DB: db}
	r.HandleFunc("/departments/", dc.Index).Methods("GET")
	r.HandleFunc("/departments/", dc.Create).Methods("POST")
	return r
}
