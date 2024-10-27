package main

import (
	"log"
	"net/http"
	"os"
	noteappservice "stub/internal/application/service/note"
	notebookappservice "stub/internal/application/service/notebook"
	"stub/internal/interfaces/rest/note"
	"stub/internal/interfaces/rest/notebook"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	sPort := os.Getenv("PORT")
	if sPort == "" {
		log.Fatal("env PORT is required")
	}
	log.Printf("Application will run on port %s", sPort)

	noteService := noteappservice.New()
	notebookService := notebookappservice.New()

	noteController := note.New(noteService)
	notebookConteroller := notebook.New(notebookService)

	r := mux.NewRouter()
	v1 := r.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/note", noteController.WriteNote).Methods("POST")
	v1.HandleFunc("/notes", noteController.ReadNotes).Methods("GET")
	v1.HandleFunc("/notebooks", notebookConteroller.ReadNotebooks).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allow specific origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	if err := http.ListenAndServe(":"+sPort, handler); err != nil {
		log.Fatal(err.Error())
	}
}
