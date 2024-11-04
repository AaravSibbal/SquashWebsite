package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/AaravSibbal/SquashWebsite/pkg/elo"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s \n\n %s", err.Error(), debug.Stack())

	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) logRequest(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.RequestURI)

		next.ServeHTTP(w, r)
	})
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				w.Header().Set("Connection", "Close")
				app.serverError(w, fmt.Errorf("%s", err))
			}

		}()

		next.ServeHTTP(w, r)
	})
}

func (app *application) secureHeaders(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-Xss-Protection", "true")

		next.ServeHTTP(w, r)
	})
}

func (app *application) readHTMLFile(name string) ([]byte, error) {
	file, err := os.OpenFile(fmt.Sprintf("ui/html/%s", name), os.O_RDONLY, 0644)
	if err != nil {
		app.errorLog.Printf("couldn't open file: %s\n%v",name, err )
		return nil, err
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		app.errorLog.Printf("couldn't read file: %s\n%v", name, err)
		return nil, err
	}

	return content, nil
}

func (app *application) createPlayerRankingHTML(players []*elo.Player) string { 
	html := "<tr>"

	for _, player := range players {
		html += fmt.Sprintf(`<td>%d</td>
		<td>
			<a href="/player/%s" target="_blank" rel="noopener noreferrer">
				%s
			</a>
		</td>
		
		<td>%d</td>
		<td>%d</td><td>%d</td><td>%d</td>`, player.Ranking, player.Name,
		player.Name, player.EloRating, player.Wins, player.Losses, 
		player.TotalMatches)
	}

	html += "</tr>"
	return html
}

func (app *application) errorHTML(errorString string) string {
	html := fmt.Sprintf(`<tr>
		<td>Error:%s</td>
	</tr>`, errorString)

	return html
}