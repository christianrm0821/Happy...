package main

import (
	"encoding/json"
	"log"
	"net/http"
	//"github.com/christianrm0821/happy/backend/structs"
)

func respondWithError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	generalErr := errorResponse{
		Error: message,
	}
	res, _ := json.Marshal(generalErr)
	w.Write(res)

}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	res, err := json.Marshal(payload)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error marshalling response")
		return
	}
	w.Write(res)
}

func main() {
	const filePath = "frontend"
	const port = "8080"

	serveMux := http.NewServeMux()

	serveMux.Handle("/", http.FileServer(http.Dir(filePath)))

	myServeMux := &http.Server{
		Addr:    ":" + port,
		Handler: serveMux,
	}

	err := myServeMux.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatal("Server error: ", err)
	}

}
