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
	ID                            int          `json:"id"`
	HomeTeam                      Team         `json:"hometeam"`
	VisitingTeam                  Team         `json:"visitingteam"`
	Championship                  Championship `json:"championship"`
	DateMatch                     time.Time    `json:"datematch"`
	Stadium                       string       `json:"stadium"`
	StatusMatch                   string       `json:"statusmatch"`
	TeamPoints1                   int          `json:"teampoints1"`
	TeamPoints2                   int          `json:"teampoints2"`
	HomeTeamWins                  int          `json:"hometeamwins"`
	VisitingTeamWins              int          `json:"visitingteamwins"`
	Draws                         int          `json:"draws"`
	HomeTeamRecentPerformance     float64      `json:"hometeamrecentperformance"`
	VisitingTeamRecentPerformance float64      `json:"visitingteamrecentperformance"`
}

// Structure that represents a user
type User struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Photo    []byte  `json:"photo"`
	Cash     float64 `json:"cash"`
}

type Championship struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Year    int     `json:"year"`
	Matches []Match `json:"matches"` // Lista de partidas que fazem parte deste campeonato
}

type Bet struct {
	Id              int     `json:"id"`
	MatchId         int     `json:"matchId"`         // Identificador da partida na qual a aposta é feita
	UserId          int     `json:"userId"`          // Identificador do usuário que fez a aposta
	SelectedOutcome string  `json:"selectedOutcome"` // O resultado selecionado pelo usuário ("HomeTeam", "VisitingTeam" ou "Draw")
	BetAmount       float64 `json:"betAmount"`       // A quantidade de dinheiro apostada pelo usuário
	PossibleReturn  float64 `json:"possibleReturn"`  // O retorno possível se a aposta for bem-sucedida
}

type Match struct {
	Id             int          `json:"id"`             // Identificador único para a partida
	ChampionshipId int          `json:"championshipId"` // Identificador do campeonato
	Championship   Championship `json:"championship"`   // Campeonato da partida
}
