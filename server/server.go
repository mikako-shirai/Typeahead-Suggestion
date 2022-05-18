package server

import (
	"net/http"
)

func setRoutes() {
    http.HandleFunc("/search", search)
    http.HandleFunc("/add", insert)
}

func InitializeServer() {
    setRoutes()
    http.ListenAndServe(":8090", nil)
}
