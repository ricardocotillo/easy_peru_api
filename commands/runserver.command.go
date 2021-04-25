package commands

import (
	"log"
	"net/http"

	"cotillo.tech/inei_api/database"
	"cotillo.tech/inei_api/router"
)

func runserver() {
	db := database.InitDB()
	r := router.InitRouter(db)
	log.Fatal(http.ListenAndServe(":8000", r))
}
