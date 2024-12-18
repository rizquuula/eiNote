package main

import (
	noteappservice "core/internal/application/service/note"
	notebookappservice "core/internal/application/service/notebook"
	"core/internal/config"
	noterepository "core/internal/infrastructure/postgresql/note"
	notebookrepository "core/internal/infrastructure/postgresql/notebook"
	notecontroller "core/internal/interfaces/rest/note"
	notebookcontroller "core/internal/interfaces/rest/notebook"
	"core/pkg/database"
	"log"
	"net/http"
	"os"

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

	config := config.Config{
		Table: config.Table{
			Note:     os.Getenv("TABLE_NOTE"),
			Notebook: os.Getenv("TABLE_NOTEBOOK"),
		},
	}

	db, err := database.Driver{
		User:         os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASSWORD"),
		DatabaseName: os.Getenv("DB_NAME"),
		Host:         os.Getenv("DB_HOST"),
		Port:         os.Getenv("DB_PORT"),
	}.NewPostgresConn()

	if err != nil {
		log.Fatalf("Error db connection: %v", err)
	}

	defer db.Close()

	noteRepository := noterepository.New(db, config.Table.Note)
	notebookRepository := notebookrepository.New(db, config.Table.Notebook)

	noteService := noteappservice.New(noteRepository)
	notebookService := notebookappservice.New(notebookRepository)

	noteController := notecontroller.New(noteService)
	notebookConteroller := notebookcontroller.New(notebookService)

	r := mux.NewRouter()
	v1 := r.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/notes", noteController.ReadNotes).Methods("GET")

	v1.HandleFunc("/note", noteController.WriteNote).Methods("POST")
	v1.HandleFunc("/note", noteController.DeleteNote).Methods("DELETE")

	v1.HandleFunc("/notebooks", notebookConteroller.ReadNotebooks).Methods("GET")

	v1.HandleFunc("/notebook", notebookConteroller.WriteNotebook).Methods("POST")
	v1.HandleFunc("/notebook", notebookConteroller.DeleteNotebook).Methods("DELETE")

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
