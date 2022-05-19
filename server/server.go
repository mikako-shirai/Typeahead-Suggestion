package server

import (
	"net/http"
)

func setRoutes() {
    http.HandleFunc("/all", getAllWords)
    http.HandleFunc("/search", getWordsByPrefix)
    http.HandleFunc("/add", insert)
}

func InitializeServer() {
    setRoutes()
    http.ListenAndServe(":8090", nil)
}
