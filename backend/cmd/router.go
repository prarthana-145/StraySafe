package main

import (
	"net/http"

	"github.com/aadi-1024/straysafe/internals/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() http.Handler {
	mux := chi.NewMux()

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("wow"))
	})
	mux.Post("/mail", handlers.SendMailPostHandler(App.Mail.Mail))
	return mux
}
