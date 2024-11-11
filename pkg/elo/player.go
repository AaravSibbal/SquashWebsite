package elo

type Player struct {
	Player_ID    string `json:"playerID"`
	Name         string `json:"name"`
	EloRating    int    `json:"eloRating"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
	Draws        int    `json:"draws"`
	TotalMatches int    `json:"totalMatches"`
	Ranking      int    `json:"ranking"`
	Discord_ID   string `json:"discordID"`
}
