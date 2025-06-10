package main

import (
	"net/http"
)


func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("error: ", err)
		return
	}
	mux := http.NewServeMux()

	myServer := 
}
