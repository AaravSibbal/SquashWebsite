package elo

import (
	"time"

)

type Match struct {
	Id string
	PlayerA    *Player
	PlayerB    *Player
	PlayerWon  *Player
	PlayerARating int
	PlayerBRating int
	PlayerAName string
	PlayerBName string
	When       *time.Time
}

func (ms *Match) New(playerA, playerB, playerWon *Player) {
	ms.PlayerA = playerA
	ms.PlayerB = playerB
	ms.PlayerWon =  playerWon
	ms.PlayerARating = playerA.EloRating
	ms.PlayerBRating = playerB.EloRating
	ms.PlayerAName = playerA.Name
	ms.PlayerBName = playerB.Name
}
