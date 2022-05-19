package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func setRoutes() {
    http.HandleFunc("/all", getAllWords)
    http.HandleFunc("/search", getWordsByPrefix)
    http.HandleFunc("/insert", insert)
}

func setPort() {
    PORT := os.Getenv("PORT")
    if PORT == "" {
        PORT = "8080"
        log.Printf(" - Defaulting to port %s\n\n", PORT)
    }

    fmt.Printf("server running on http://localhost:%s\n", PORT)
    if err := http.ListenAndServe(":"+PORT, nil); err != nil {
        log.Fatal(err)
    }
}

func InitializeServer() {
    setRoutes()
    setPort()
}
