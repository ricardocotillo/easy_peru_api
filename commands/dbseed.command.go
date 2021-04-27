package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"cotillo.tech/inei_api/database"
	"cotillo.tech/inei_api/models"
	"gorm.io/gorm"
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

	fcs, _ := ioutil.ReadFile(basedir + "/data/constante_soles.json")
	saveFileToDB(fcs, "miles_de_soles", "constante", db)
	fcp, _ := ioutil.ReadFile(basedir + "/data/constante_porcentual.json")
	saveFileToDB(fcp, "porcentual", "constante", db)
	fcvp, _ := ioutil.ReadFile(basedir + "/data/constante_variacion_porcentual.json")
	saveFileToDB(fcvp, "variacion_porcentual", "constante", db)
	fcs, _ = ioutil.ReadFile(basedir + "/data/corriente_soles.json")
	saveFileToDB(fcs, "miles_de_soles", "corriente", db)
	fcp, _ = ioutil.ReadFile(basedir + "/data/corriente_porcentual.json")
	saveFileToDB(fcp, "porcentual", "corriente", db)
	fcvp, _ = ioutil.ReadFile(basedir + "/data/corriente_variacion_porcentual.json")
	saveFileToDB(fcvp, "variacion_porcentual", "corriente", db)
}

func saveFileToDB(f []byte, st string, vt string, db *gorm.DB) {
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
				p := models.Product{Year: ym["año"].(string), Value: am["valor"].(float64), Structure: st, ValueType: vt, DepartamentID: dep.ID, EconomicActivityID: ac.ID}
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
