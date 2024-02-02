package structs

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Structure that represents a team
type Team struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
	Color1  string `json:"color1"`
}

// Structure that represents a player
type Player struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	City     string    `json:"city"`
	Country  string    `json:"country"`
	Birth    time.Time `json:"birth"`
	IdTeam   int       `json:"idteam"`
	Height   float64   `json:"height"`
	Position string    `json:"position"`
}

type Live struct {
	ID             int       `json:"id"`
	IdTeam1        int       `json:"teamid1"`
	IdTeam2        int       `json:"teamid2"`
	IdChampionship int       `json:"idchampionship"`
	DateMatch      time.Time `json:"datematch"`
	Stadium        string    `json:"stadium"`
	TeamPoints1    int       `json:"teampoints1"`
	TeamPoints2    int       `json:"teampoints2"`
}
