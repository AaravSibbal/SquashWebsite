package psql

import (
	"context"
	"database/sql"
	"time"

	"github.com/AaravSibbal/SquashWebsite/pkg/elo"
)

func GetRanking(db *sql.DB, ctx *context.Context) ([]*elo.Player, error) {
	// need to get all players sorted by elo
	stmt := `SELECT *, COUNT(*) OVER() FROM player ORDER BY elo_rating DESC;`

	newCtx, cancel := context.WithTimeout(*ctx, time.Second*5)
	defer cancel()

	rows, err := db.QueryContext(newCtx, stmt)

	if err != nil {
		return nil, err
	}

	var totalCount int
	firstRow := true
	var players []*elo.Player

	for rows.Next() {
		player := &elo.Player{}
		
		err := rows.Scan(player.Player_ID, player.Name, player.EloRating, player.Wins, player.Losses, player.Draws, player.TotalMatches, &totalCount)
		if err != nil {
			return nil, err	
		}

		if firstRow {
			firstRow = false
			players = make([]*elo.Player, 0, totalCount)
		}

		players = append(players, player)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return players, nil
}