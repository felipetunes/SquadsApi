package structs

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Structure that represents a team
type Team struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
}

// Structure that represents a player
type Player struct {
	ID      int            `json:"id"`
	Name    string         `json:"name"`
	City    string         `json:"city"`
	Country string         `json:"country"`
	Birth   time.Time      `json:"birth"`
	IdTeam  int            `json:"idteam"`
	Height  sql.NullString `json:"height"`
}
