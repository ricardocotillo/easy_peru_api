package controllers

import (
	"net/http"
	"strconv"

	"cotillo.tech/inei_api/models"
	"cotillo.tech/inei_api/serializers"
)

type ProductController Controller

func (pc ProductController) Index(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	st := params.Get("structure")
	vt := params.Get("valueType")
	y := params.Get("year")
	d, _ := strconv.Atoi(params.Get("department"))
	ea, _ := (strconv.Atoi(params.Get("economicActivity")))
	if y != "" || d > 0 || ea > 0 {
		var ps []serializers.ProductBySerializer
		pc.DB.Table("products").Joins("Department").Joins("EconomicActivity").Where(&models.Product{Structure: st, ValueType: vt, Year: y, EconomicActivityID: uint(ea), DepartmentID: uint(d)}).Find(&ps)
		jsonResponse(w, ps, http.StatusOK)
		return
	}
	var ps []serializers.ProductAggregateSerializer
	pc.DB.Table("products").Where(&models.Product{Structure: st, ValueType: vt, Year: y, EconomicActivityID: uint(ea), DepartmentID: uint(d)}).Find(&ps)
	jsonResponse(w, ps, http.StatusOK)
}
