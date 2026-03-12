package main

import (
	"encoding/json" // convertir datos de json a struct de go
	"log"           // imprimir mensajes
	"net/http"      // maneja las peticioens y respuestas HTTP
	"os"            // leer archivos e interactuar con el sistema
	"strconv"       // convertir texto a numeros
	"strings"
)

type Band struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Genre   string `json:"genre"`
	Year    int    `json:"year"`    // año de formación
	Albums  int    `json:"albums"`  // cantidad de albumes
	Members int    `json:"members"` // cantidad de miembros
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
	response := Message{
		Message: "pong",
	}

	writeJSON(w, http.StatusOK, response)
}

func bandsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetBands(w, r)
	case http.MethodPost:
		handleCreateBand(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetBands(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	idParam := query.Get("id") // busca por ?id=...
	nameParam := query.Get("name")
	genreParam := query.Get("genre")
	albumsParam := query.Get("albums")
	yearParam := query.Get("year")
	membersParam := query.Get("members")

	if idParam != "" { // Buscar por ID
		id, err := strconv.Atoi(idParam)
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
		return
	}

	var filtered []Band
	for _, band := range bands {
		// filtro por nombre
		if nameParam != "" && !strings.Contains(strings.ToLower(band.Name), strings.ToLower(nameParam)) {
			continue
		}
		// Filtro por Género
		if genreParam != "" && !strings.EqualFold(band.Genre, genreParam) {
			continue
		}
		// Filtro por Año
		if yearParam != "" {
			y, err := strconv.Atoi(yearParam)
			if err != nil || band.Year != y {
				continue
			}
		}
		// Filtro por Álbumes
		if albumsParam != "" {
			a, err := strconv.Atoi(albumsParam)
			if err != nil || band.Albums != a {
				continue
			}
		}
		// Filtro por Miembros
		if membersParam != "" {
			m, err := strconv.Atoi(membersParam)
			if err != nil || band.Members != m {
				continue
			}
		}

		filtered = append(filtered, band)
	}

	// Si el slice está vacío, lo inicializamos para que devuelva "[]" en vez de "null"
	if filtered == nil {
		filtered = []Band{}
	}

	writeJSON(w, http.StatusOK, filtered)
}

func handleCreateBand(w http.ResponseWriter, r *http.Request) {
	var newband Band

	err := json.NewDecoder(r.Body).Decode(&newband) // para decodificar el cuerpo del json de la request
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if newband.Name == "" || newband.Genre == "" || newband.Year == 0 || newband.Albums == 0 || newband.Members == 0 {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	newband.ID = generateNextID()

	bands = append(bands, newband)

	writeJSON(w, http.StatusCreated, newband)

}

func generateNextID() int {
	maxID := 0
	for _, band := range bands {
		if band.ID > maxID {
			maxID = band.ID
		}
	}
	return maxID + 1
}

func saveBands() {
	data, err := json.MarshalIndent(bands, "", "  ")
	if err != nil {
		log.Println("Error marshaling JSON: ", err)
		return
	}

	err = os.WriteFile("./data/bands.json", data, 0644)
	if err != nil {
		log.Println("Error writing file: ", err)
	}
}

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
