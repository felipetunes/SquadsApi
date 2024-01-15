package structs

import (
	_ "github.com/go-sql-driver/mysql"
)

// Estrutura que representa um time
type Team struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
}

// Estrutura que representa um jogador
type Player struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
	Age     int    `json:"age"`
	IdTeam  int    `json:"idteam"`
}
