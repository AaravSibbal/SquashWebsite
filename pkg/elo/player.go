package elo

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type Player struct {
	Player_ID    *uuid.UUID
	Name         string
	EloRating    int
	Wins         int
	Losses       int
	Draws        int
	TotalMatches int
	Ranking      int
}

func (player *Player) Equals(p *Player) bool {
	result := strings.Compare(player.Name, p.Name)
	return result == 0
}

func (player *Player) UpdatePlayer(match *Match) {
	result := player.DidPlayerWin(match)

	if result {
		player.Wins++
	} else {
		player.Losses++
	}
	player.TotalMatches++
}

func (player *Player) DidPlayerWin(match *Match) bool {
	return player.Equals(match.PlayerWon)
}

func (player *Player) String() string {
	return fmt.Sprintf("Name: %s\nElo Rating: %d\nWins: %d\nLosses: %d\nTotal Matches: %d\n",
		player.Name, player.EloRating, player.Wins, player.Losses, player.TotalMatches)
}

type Players struct {
	players []*Player
}

func (pls Players) New(name string) *Player {
	p := &Player{
		Name:         name,
		EloRating:    400,
		Wins:         0,
		Losses:       0,
		Draws:        0,
		TotalMatches: 0,
	}
	return p
}

func (pls *Players) CreatePlayer(playerID *uuid.UUID, name string, eloRating int,
	wins int, losses int, draws int, totalMatches int) *Player {
	p := &Player{
		Player_ID:    playerID,
		Name:         name,
		EloRating:    eloRating,
		Wins:         wins,
		Losses:       losses,
		Draws:        draws,
		TotalMatches: totalMatches,
	}

	return p
}

func (pls *Players) AddPlayer(p *Player) bool {
	if p == nil {
		return false
	}

	pls.players = append(pls.players, p)
	return true
}

/*
removes player given the name
returns true if name was for a valid Player
otherwise returns false
*/
func (pls *Players) RemovePlayer(name string) bool {
	idx := pls.GetPlayerIdx(name)
	if idx == -1 {
		return false
	}
	player := pls.GetPlayerByIdx(idx)
	if player == nil {
		return false
	}
	pls.players = append(pls.players[idx:], pls.players[:idx+1]...)
	return true
}

func (pls *Players) GetPlayer(name string) *Player {
	for _, player := range pls.players {
		if player.Name == name {
			return player
		}
	}

	return nil
}

/*
return player Object if index is valid return nil otherwise
*/
func (pls *Players) GetPlayerByIdx(idx int) *Player {
	if idx < 0 || idx >= len(pls.players) {
		return nil
	}

	return pls.players[idx]
}

/*
returns index of a player if found returns -1 otherwise
*/
func (pls *Players) GetPlayerIdx(name string) (idx int) {
	for i, player := range pls.players {
		if player.Name == name {
			return i
		}
	}

	return -1
}
