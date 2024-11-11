package server

import (
	"encoding/json"
	"html/template"
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

func (app *application) playerRankings(w http.ResponseWriter, r *http.Request) {
	players, err := psql.GetRanking(app.db, app.ctx)
	if err != nil {
		app.serverError(w, err)
		return
	}

	playersJson, err := json.Marshal(players)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(playersJson)
}

func (app *application) playerStat(w http.ResponseWriter, r *http.Request) {
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

func (app *application) playerGraph(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get(":name")
	if name == "" {
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

	tmpl := "{{.Element}} {{.Script}} {{.Option}}"
	t := template.New("snippet")
	t, err = t.Parse(tmpl)
	if err != nil {
		panic(err)
	}

	line := chart.GetEloChart(matches, name)
	lineSnippet := line.RenderSnippet()

	data := struct {
		Element template.HTML
		Script  template.HTML
		Option  template.HTML
	}{
		Element: template.HTML(lineSnippet.Element),
		Script:  template.HTML(lineSnippet.Script),
		Option:  template.HTML(lineSnippet.Option),
	}

	err = t.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func (app *application) playerMatches(w http.ResponseWriter, r *http.Request) {
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
		} else {
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

func (app *application) playerHtml(w http.ResponseWriter, r *http.Request) {
	html, err := app.readHTMLFile("player.html")
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(html)
}
