package elo

import (
	"errors"
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

type Matches struct {
	matches []*Match
}

func (ms Matches) New(playerA *Player, playerB *Player, playerWon *Player,
	id string, when *time.Time) *Match {

	m := &Match{
		Id:            id,
		PlayerA:       playerA,
		PlayerB:       playerB,
		PlayerWon:     playerWon,
		PlayerARating: playerA.EloRating,
		PlayerBRating: playerB.EloRating,
		PlayerAName:   playerA.Name,
		PlayerBName:   playerB.Name,
		When:          when,
	}

	return m
}

func (ms *Matches) AddMatch(match *Match) bool {
	if match == nil {
		return false
	}

	ms.matches = append(ms.matches, match)
	return true
}

func (ms *Matches) RemoveMatch(id string) (bool, error) {
	matchIdx := ms.GetMatchIdx(id)
	if matchIdx == -1 {
		return false, errors.New("couldn't find the match")
	}
	if matchIdx < 0 || matchIdx > len(ms.matches) {
		return false, errors.New("index out of bounds for matches")
	}

	ms.matches = append(ms.matches[:matchIdx], ms.matches[matchIdx+1:]...)
	return true, nil
}

func (ms *Matches) GetMatchIdx(id string) int {
	for idx, match := range ms.matches {
		if match.Id == id {
			return idx
		}
	}

	return -1
}
