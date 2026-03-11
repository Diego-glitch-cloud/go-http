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
	loadBands()

	http.HandleFunc("/api/ping", pingHandler)
	http.HandleFunc("/api/bands", bandsHandler)

	log.Println("POST JSON API running on :80")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func loadBands() {
	file, err := os.ReadFile("./data/bands.json")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	err = json.Unmarshal(file, &bands)
	if err != nil {
		log.Fatal("Error parsing JSON:", err)
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	response := Message {
		Message: "pong",
	}

	writeJSON(w, http.StatusOK, response)
}

func bandsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			handleGetBands(w, r)
		case http.MethodPOst:
			handleCreateBand(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed	)
	}
}

func handleGetBands(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	idParam := query.Get("id") // busca por ?id=...

	if idParam == "" {
		writeJSON(w, http.StatusOK, bands)
		return
	}

	id, err := strconv.Atoi(idParam) // convierte el id de string a int
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	
	for _, band := range bands {
		if band.ID == id {
			writeJSON(w, http.StatusOK, band)
			return
		}
	}

	http.Error(w, "Band not found", http.StatusNotFound)
}
