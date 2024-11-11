package elo

import (
	"time"
)

type Match struct {
	Id            string
	PlayerA       *Player
	PlayerB       *Player
	PlayerWon     *Player
	PlayerARating int
	PlayerBRating int
	PlayerAName   string
	PlayerBName   string
	When          *time.Time
}

type MatchJson struct {
	PlayerA       string `json:"playerA"`
	PlayerB       string `json:"playerB"`
	PlayerWon     string `json:"playerWon"`
	PlayerARating int    `json:"playerARating"`
	PlayerBRating int    `json:"playerBRating"`
	When          string `json:"when"`
}

func (m *Match) ToJsonObj() *MatchJson {
	matchJson := &MatchJson{
		PlayerA:       m.PlayerA.Name,
		PlayerB:       m.PlayerB.Name,
		PlayerWon:     m.PlayerWon.Name,
		PlayerARating: m.PlayerARating,
		PlayerBRating: m.PlayerBRating,
		When:          m.When.Format("31 Jan 2024"),
	}

	return matchJson
}
