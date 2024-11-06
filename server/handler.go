package server

import (
	"net/http"

	psql "github.com/AaravSibbal/SquashWebsite/pkg/sql"
)

func (app *application) pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	htmlFile, err := app.readHTMLFile("index.html")
	if err != nil {
		app.serverError(w, err)
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(htmlFile)
}

func (app *application) playerRankings(w http.ResponseWriter, r *http.Request){
	players, err := psql.GetRanking(app.db, app.ctx)
	if err != nil {
		app.serverError(w, err)
	}

	if len(players) == 0 {
		errHTML := app.errorHTML("There are no players in the system yet")
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(errHTML))
		return
	}

	playerHTML := app.createPlayerRankingHTML(players)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(playerHTML))
}

func (app *application) playerStat(w http.ResponseWriter, r *http.Request){

}

func (app *application) playerGraph(w http.ResponseWriter, r *http.Request){

}

func (app *application) playerMatches(w http.ResponseWriter, r *http.Request){

}