package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/news/create", http.HandlerFunc(app.createNewsForm))
	mux.Post("/news/create", http.HandlerFunc(app.createNews))
	mux.Get("/category", http.HandlerFunc(app.category))
	mux.Get("/contact", http.HandlerFunc(app.contact))
	mux.Get("/news/:id", http.HandlerFunc(app.showNews))
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static/", fileServer))
	return standardMiddleware.Then(mux)
}
