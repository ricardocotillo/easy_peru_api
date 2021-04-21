package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"cotillo.tech/inei_api/models"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/pbi/", pbiHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func pbiHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	err := checkParams(params)
	check(err)
	e := params.Get("estructura")
	v := params.Get("valor")
	d := params.Get("departamento")
	a := params.Get("año")

	b, err := readFile(v + "_" + e)
	check(err)
	cd := filterF(b, d, a)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cd)
}

func filterF(b string, d string, a string) models.Cuadro {
	var result models.Cuadro
	json.Unmarshal([]byte(b), &result)
	if d == "" && a == "" {
		return result
	}
	if d != "" {
		ds := []models.Departamento{}
		for _, dep := range result.Departamentos {
			if dep.Nombre == d {
				ds = append(ds, dep)
			}
		}
		result.Departamentos = ds
	}

	if a != "" {
		ds := []models.Departamento{}
		for _, dep := range result.Departamentos {
			as := []models.Año{}
			for _, an := range dep.Años {
				if an.Año == a {
					as = append(as, an)
				}
			}
			dep.Años = as
			ds = append(ds, dep)
		}
		result.Departamentos = ds
	}
	return result
}

func readFile(fp string) (string, error) {
	dir_path := "data"
	dat, err := ioutil.ReadFile(dir_path + "/" + fp + ".json")
	return string(dat), err
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkParams(vs url.Values) error {
	e := vs.Get("estructura")
	v := vs.Get("valor")
	if e == "" || v == "" {
		return errors.New("Los parámetros extructura y valor son requeridos")
	}
	return nil
}
