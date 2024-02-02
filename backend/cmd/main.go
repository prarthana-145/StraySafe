package main

import (
	"log"
	"net/http"

	"github.com/aadi-1024/straysafe/internals/mailer"
)

var App *Config

func main() {
	App = &Config{}

	App.Mail = mailer.NewMailer("username", "password", "localhost")
	go App.Mail.StartService()

	srv := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: NewRouter(),
	}

	log.Println("starting server on port 8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
