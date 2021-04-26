package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

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

	basedir, _ := os.Getwd()

	f, _ := ioutil.ReadFile(basedir + "/data/constante_soles.json")
	var result map[string]interface{}
	json.Unmarshal(f, &result)
	deps := result["departamentos"].([]interface{})
	for _, d := range deps {
		dm := d.(map[string]interface{})
		dn := dm["nombre"]
		var dep models.Departament
		db.Where(&models.Departament{Name: dn.(string)}).Find(&dep)
		ys := dm["años"].([]interface{})
		for _, y := range ys {
			ym := y.(map[string]interface{})
			acs := ym["actividades"].([]interface{})
			for _, a := range acs {
				am := a.(map[string]interface{})
				ac := models.EconomicActivity{}
				db.Where(&models.EconomicActivity{Name: am["nombre"].(string)}).Find(&ac)
				p := models.Product{Year: ym["año"].(string), Value: am["valor"].(float64), Structure: "miles_de_soles", ValueType: "constante", DepartamentID: dep.ID, EconomicActivityID: ac.ID}
				result := db.Create(&p)
				if result.Error != nil {
					fmt.Println(am["nombre"])
				}
			}
		}

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
