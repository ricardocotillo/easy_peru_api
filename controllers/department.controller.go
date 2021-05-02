package controllers

import (
	"net/http"

	"cotillo.tech/inei_api/serializers"
)

type DepartamentController Controller

func (dc DepartamentController) Index(w http.ResponseWriter, r *http.Request) {
	var departments []serializers.DepartmentSerializer
	dc.DB.Table("departments").Find(&departments)
	jsonResponse(w, departments, http.StatusOK)
}
func (dc DepartamentController) Create(w http.ResponseWriter, r *http.Request)  {}
func (dc DepartamentController) Show(w http.ResponseWriter, r *http.Request)    {}
func (dc DepartamentController) Update(w http.ResponseWriter, r *http.Request)  {}
func (dc DepartamentController) Destroy(w http.ResponseWriter, r *http.Request) {}
