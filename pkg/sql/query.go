package psql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/AaravSibbal/SquashWebsite/pkg/elo"
)

func GetRanking(db *sql.DB, ctx *context.Context) ([]*elo.Player, error) {
	// need to get all players sorted by elo
	
	newCtx, cancel := context.WithTimeout(*ctx, time.Second*5)
	defer cancel()

	stmt, err := db.PrepareContext(newCtx, `SELECT *, COUNT(*) OVER() FROM player ORDER BY elo_rating DESC;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	var totalCount int
	count := 1
	firstRow := true
	var players []*elo.Player

	for rows.Next() {
		player := &elo.Player{}

		err := rows.Scan(&player.Player_ID, &player.Name, &player.EloRating, &player.Wins, &player.Losses, &player.Draws, &player.TotalMatches, &player.Discord_ID, &totalCount)
		
		player.Ranking = count
		if err != nil {
			return nil, err	
		}

		if firstRow {
			firstRow = false
			players = make([]*elo.Player, 0, totalCount)
		}

		count++
		players = append(players, player)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return players, nil
}

func GetPlayerData(db *sql.DB, ctx *context.Context, name string) (*elo.Player, error){

	newCtx, cancel := context.WithTimeout(*ctx, 5*time.Second)
	defer cancel()

	stmt, err := db.PrepareContext(newCtx, `SELECT * FROM player WHERE name=$1;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	
	row := stmt.QueryRowContext(newCtx, name)
	
	player := &elo.Player{}

	err = row.Scan(&player.Player_ID, &player.Name, &player.EloRating,
		&player.Wins, &player.Losses, &player.Draws, &player.TotalMatches, &player.Discord_ID)
	if err != nil {
		return nil, err
	}

	return player, nil
}

func GetPlayerMatches(db *sql.DB, ctx *context.Context, name string, offset int) ([]*elo.MatchJson, error) {
	
	newCtx, cancel := context.WithTimeout(*ctx, 5*time.Second)
	defer cancel()

	stmt, err := db.PrepareContext(newCtx, `SELECT * from match WHERE 
	player_a_name=$1 OR player_b_name=$1 
	ORDER BY match_time DESC LIMIT 20 OFFSET $2;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(newCtx, name, offset)
	if err != nil {
		return nil, err
	}

	matches := make([]*elo.MatchJson, 0, 20)

	for rows.Next() {

		var id, playerAID, playerBID, playerWonID, playerAName, playerBName string
		var playerARating, playerBRating int
		when := time.Time{}
	
		err := rows.Scan(&id, &playerAID, &playerBID, &playerWonID, &when,
				&playerARating, &playerBRating, &playerBName, &playerAName)
		if err != nil {
			return nil, err
		}

		year, month, day := when.Date()

		match := &elo.MatchJson{
			PlayerA: playerAName,
			PlayerB: playerBName,
			PlayerWon: GetPlayerWon(playerAID, playerWonID, playerAName, playerBName),
			PlayerARating: playerARating,
			PlayerBRating: playerBRating,
			When: fmt.Sprintf("%d %s, %d", day, month.String(), year),
		}

		matches = append(matches, match)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return matches, nil
}

func GetPlayerWon(playerAID, playerWonID, playerAName, playerBName string) string {
	if playerWonID == playerAID{
		return playerAName
	}

	return playerBName
}