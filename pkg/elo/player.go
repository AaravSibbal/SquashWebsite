package elo

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type Player struct {
	Player_ID 	*uuid.UUID
	Name         string
	EloRating    int
	Wins         int
	Losses       int
	Draws        int
	TotalMatches int
}

func (player *Player) Equals(p *Player) bool {
	result := strings.Compare(player.Name, p.Name)
	return result == 0
}

func (player *Player) UpdatePlayer(match *Match){
	result:= player.DidPlayerWin(match)

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
	players map[string]*Player
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

func (pls *Players) AddPlayer(p *Player) bool {
	_, exists := pls.players[p.Name]
	if exists {
		return false
	}

	pls.players[p.Name] = p
	return true
}

func (pls *Players) RemovePlayer(name string) bool {
	_, exists := pls.players[name]
	if !exists {
		return false
	}

	delete(pls.players, name)
	return true
}

func (pls *Players) GetPlayer(name string) (*Player, bool) {
	player, exists := pls.players[name]
	if !exists {
		return nil, false
	}
	return player, true
}
