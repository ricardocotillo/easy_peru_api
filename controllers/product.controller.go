package controllers

import (
	"net/http"

	"cotillo.tech/inei_api/models"
)

type ProductController Controller

func (pc ProductController) Index(w http.ResponseWriter, r *http.Request) {
	var ps []models.Product
	pc.DB.Where(&models.Product{Structure: "miles_de_soles", ValueType: "constante", Year: "2019"}).Find(&ps)
	jsonResponse(w, ps, http.StatusOK)
}
