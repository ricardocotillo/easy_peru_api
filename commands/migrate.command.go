package commands

import (
	"cotillo.tech/inei_api/database"
	"cotillo.tech/inei_api/models"
)

func migrate() {
	db := database.InitDB()
	db.AutoMigrate(&models.Department{}, &models.EconomicActivity{}, &models.Product{})
}
