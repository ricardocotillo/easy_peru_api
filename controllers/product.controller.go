package controllers

import (
	"net/http"
	"strconv"

	"cotillo.tech/inei_api/models"
)

type ProductController Controller

func (pc ProductController) Index(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	st := params.Get("structure")
	vt := params.Get("valueType")
	y := params.Get("year")
	d, _ := strconv.Atoi(params.Get("department"))
	ea, _ := (strconv.Atoi(params.Get("economicActivity")))
	var ps []models.Product
	pc.DB.Joins("EconomicActivity").Joins("Department").Where(&models.Product{Structure: st, ValueType: vt, Year: y, EconomicActivityID: uint(ea), DepartmentID: uint(d)}).Find(&ps)
	jsonResponse(w, ps, http.StatusOK)
}
