package main

import (
	"database/sql"
	"log"
	"net/http"

	"br.com.goalbums/internal/service"
	"br.com.goalbums/internal/web"

	_ "github.com/lib/pq"
)

func main() {
	dbConfig := "host=localhost port=5432 user=admin password=root dbname=postgresdb sslmode=disable"

	db, err := sql.Open("postgres", dbConfig)

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		panic(err)
	}

	defer db.Close()

	albumService := service.NewAlbumService(db)
	albumHandlers := web.NewAlbumHandler(albumService)

	router := http.NewServeMux()

	router.HandleFunc("GET /albums", albumHandlers.GetAlbums)
	router.HandleFunc("GET /albums/{id}", albumHandlers.GetAlbumByID)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
