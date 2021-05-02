package controllers

import (
	"net/http"

	"cotillo.tech/inei_api/serializers"
)

type EconomicActivityController Controller

func (eac EconomicActivityController) Index(w http.ResponseWriter, r *http.Request) {
	var eas []serializers.EconomicActivitySerializer
	eac.DB.Table("economic_activities").Find(&eas)
	jsonResponse(w, eas, http.StatusOK)
}
