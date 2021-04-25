package controllers

import (
	"encoding/json"
	"net/http"
)

type Controller interface {
	index()
	create()
	show()
	update()
	destroy()
}

func httpBadRequest(msg string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(msg))
}

func jsonResponse(w http.ResponseWriter, d interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(d)
}
