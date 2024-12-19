package main

import (
	"database/sql"
	"log"
	"maker-checker/internal/db"
	"maker-checker/internal/handler"
	"maker-checker/internal/mailer"
	"maker-checker/internal/repository"
	"maker-checker/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, err := sql.Open("sqlite3", "file:messages.db?_foreign_keys=on")
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	queries := db.New(database)
	repo := repository.NewMessageRepository(queries)
	mlr := mailer.NewConsoleMailer()
	svc := service.NewMessageService(repo, mlr)
	h := handler.NewMessageHandler(svc)

	r := chi.NewRouter()
	r.Post("/messages", h.Create)
	r.Get("/messages", h.ListPending)
	r.Post("/messages/{id}/approve", h.Approve)
	r.Post("/messages/{id}/reject", h.Reject)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
