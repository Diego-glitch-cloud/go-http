package main 

import (
"encoding/json" // convertir datos de json a struct de go
"log"			// imprimir mensajes
"net/http"		// maneja las peticioens y respuestas HTTP
"os"			// leer archivos e interactuar con el sistema
"strconv"		// convertir texto a numeros
)


type Band struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Genre string `json:"genre"`
	Year int `json:"year"`			// año de formación
	Albums int `json:"albums"`		// cantidad de albumes 
	Members int `json:"members"`	// cantidad de miembros
}

type Message struct {
	Message string `json:"message"`
}

var bands []Band

func main() {
	loadBandas()

	http.HandleFunc("/api/ping", pingHandler)
	http.HandleFunc("/api/bands", bandsHandler)

	log.Println("POST JSON API running on :80")
	log.Fatal(http.ListenAndServe(":80", nil))
}