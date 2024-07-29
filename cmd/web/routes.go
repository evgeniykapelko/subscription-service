package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(app.SessionLoad)

	mux.Get("/", app.HomePage)

	mux.Get("/login", app.LoginPage)
	mux.Post("/login", app.PostLoginPage)
	mux.Get("/logout", app.Logout)
	mux.Get("/register", app.RegisterPage)
	mux.Post("/register", app.PostRegisterPage)
	mux.Get("/activate-account", app.ActivateAccount)

	//mux.Get("test-email", func(writer http.ResponseWriter, request *http.Request) {
	//	m := Mail{
	//		Domain:      "localhost",
	//		Host:        "localhost",
	//		Port:        1025,
	//		Encryption:  "none",
	//		FromAddress: "info@example.com",
	//		FromName:    "info",
	//		ErrorChan:   make(chan error),
	//	}
	//
	//	msg := Message{
	//		To:      "me@here.com",
	//		Subject: "Welcome",
	//		Data:    "Hello world!",
	//	}
	//
	//	m.SendMail(msg, make(chan error))
	//
	//})

	return mux
}
