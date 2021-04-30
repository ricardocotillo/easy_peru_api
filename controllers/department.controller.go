package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"cotillo.tech/inei_api/models"
)

type DepartamentController Controller

func (dc DepartamentController) Index(w http.ResponseWriter, r *http.Request) {
	var departments []models.Department
	result := dc.DB.Find(&departments)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
func (dc DepartamentController) Create(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httpBadRequest("Nombre es requerido", w)
		return
	}
	var d models.Department
	err = json.Unmarshal(b, &d)
	if err != nil {
		httpBadRequest("Sucedió un error", w)
		return
	}
	dc.DB.Create(&d)
	jsonResponse(w, d, http.StatusOK)
}
func (dc DepartamentController) Show(w http.ResponseWriter, r *http.Request)    {}
func (dc DepartamentController) Update(w http.ResponseWriter, r *http.Request)  {}
func (dc DepartamentController) Destroy(w http.ResponseWriter, r *http.Request) {}
