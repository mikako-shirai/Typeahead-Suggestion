package server

import (
	"net/http"
)

func setRoutes() {
	http.HandleFunc("/headers", headers)
  http.HandleFunc("/search", search)
}

func InitializeServer() {
	setRoutes()
	http.ListenAndServe(":8090", nil)
}
