package controllers

import (
	"net/http"

	"cotillo.tech/inei_api/models"
)

type ProductController Controller

func (pc ProductController) Index(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	st := params.Get("structure")
	vt := params.Get("valueType")
	var ps []models.Product
	pc.DB.Where(&models.Product{Structure: st, ValueType: vt, Year: "2019"}).Find(&ps)
	jsonResponse(w, ps, http.StatusOK)
}
