package server

import "net/http"

func (app *application) pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	htmlFile, err := app.readHTMLFile("index.html")
	if err != nil {
		app.serverError(w, err)
	}

	w.Write(htmlFile)
}
