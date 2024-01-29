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
