package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AaravSibbal/SquashWebsite/pkg/chart"
	psql "github.com/AaravSibbal/SquashWebsite/pkg/sql"
)

func (app *application) pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	htmlFile, err := app.readHTMLFile("index.html")
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(htmlFile)
}

// TODO: change this function so it gives out json rather than html
func (app *application) playerRankings(w http.ResponseWriter, r *http.Request){
	players, err := psql.GetRanking(app.db, app.ctx)
	if err != nil {
		app.serverError(w, err)
		return 
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
	name := r.URL.Query().Get(":name")
	if name == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}	

	player, err := psql.GetPlayerData(app.db, app.ctx, name)
	if err != nil {
		app.serverError(w, err)
		return
	}

	playerJson, err := json.Marshal(player)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(playerJson)
}

func (app *application) playerGraph(w http.ResponseWriter, r *http.Request){
	name := r.URL.Query().Get(":name")
	if name == ""{
		app.clientError(w, http.StatusBadRequest)
		return
	}

	matches, err := psql.GetPlayerMatches(app.db, app.ctx, name, 0)
	if err != nil {
		app.serverError(w, err)
		return
	}

	if len(matches) == 0 {
		 html := "<p>No Match Data Found</p>"
		 w.Header().Set("Content-Type", "text/html")
		 w.Write([]byte(html))
		 return
	}

	line := chart.GetEloChart(matches, name)

	w.Header().Set("Content-Type", "text/html")
	w.Write(line.RenderContent())
}

func (app *application) playerMatches(w http.ResponseWriter, r *http.Request){
	var pageInt int
	name := r.URL.Query().Get(":name")
	if name == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	pageStr := r.URL.Query().Get("record")
	if pageStr == "" {
		pageInt = 0
	} else {
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			pageInt = 0
		} else{
			pageInt = page
		}
	}

	matches, err := psql.GetPlayerMatches(app.db, app.ctx, name, pageInt)
	if err != nil {
		app.serverError(w, err)
		return
	}

	matchesJson, err := json.Marshal(matches)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(matchesJson)
}