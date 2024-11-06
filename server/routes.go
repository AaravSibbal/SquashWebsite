package server

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) Routes() http.Handler {

	standardMiddleware := alice.New(app.logRequest, app.recoverPanic, app.secureHeaders)

	mux := pat.New()

	mux.Get("/ping", standardMiddleware.ThenFunc(app.pong))
	mux.Get("/", standardMiddleware.ThenFunc(app.home))
	mux.Get("/player/ranking", standardMiddleware.ThenFunc(app.playerRankings))
	mux.Get("/player/:name/stat", standardMiddleware.ThenFunc(app.playerStat))
	mux.Get("/player/:name/graph", standardMiddleware.ThenFunc(app.playerGraph))
	mux.Get("/player/:name/matches", standardMiddleware.ThenFunc(app.playerMatches))

	fileServer := http.FileServer(http.Dir("ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// http.Handle("/", fileServer)

	return standardMiddleware.Then(mux)

}