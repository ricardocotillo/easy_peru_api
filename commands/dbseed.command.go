package commands

import (
	"cotillo.tech/inei_api/database"
	"cotillo.tech/inei_api/models"
)

func dbseed() {
	db := database.InitDB()
	for _, n := range departments {
		d := models.Departament{Name: n}
		db.Create(&d)
	}
	for _, n := range activities {
		a := models.EconomicActivity{Name: n}
		db.Create(&a)
	}
}

var departments = []string{
	"amazonas",
	"áncash",
	"apurímac",
	"arequipa",
	"ayacucho",
	"cajamarca",
	"cusco",
	"huancavelica",
	"huánuco",
	"ica",
	"junín",
	"la libertad",
	"lambayeque",
	"lima",
	"loreto",
	"madre de dios",
	"moquegua",
	"pasco",
	"piura",
	"puno",
	"san martín",
	"tacna",
	"tumbes",
	"ucayali",
}

var activities = []string{
	"Agricultura, Ganadería, Caza y Silvicultura",
	"Pesca y Acuicultura",
	"Extracción de Petróleo, Gas y Minerales",
	"Manufactura",
	"Electricidad, Gas y Agua",
	"Construcción",
	"Comercio",
	"Transporte, Almacen., Correo y Mensajería",
	"Alojamiento y Restaurantes",
	"Telecom. y Otros Serv. de Información",
	"Administración Pública y Defensa",
	"Otros Servicios",
}
