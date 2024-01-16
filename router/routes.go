package router

import (
	"apiSquads/db"
	"apiSquads/structs"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

// GetAllTeams godoc
// @Summary Get all teams
// @Description Get all teams
// @Tags Teams
// @Accept  json
// @Produce  json
// @Success 200 {array} structs.Team
// @Router /api/v1/team/getall [get]
func GetAllTeams(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Executa a consulta SQL
	rows, err := db.Query("SELECT * FROM Team")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Cria um slice de times
	teams := []structs.Team{}

	// Lê os resultados
	for rows.Next() {
		// Cria um time vazio
		team := structs.Team{}

		// Preenche o time com os dados da linha
		err = rows.Scan(&team.ID, &team.Name, &team.City, &team.Country)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// Adiciona o time ao slice de times
		teams = append(teams, team)
	}
	// Fecha o conjunto de resultados
	rows.Close()

	// Converte os times em JSON
	teamsJSON, err := json.Marshal(teams)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Envia a resposta em JSON
	return c.JSONBlob(http.StatusOK, teamsJSON)
}

// GetAllPlayers godoc
// @Summary Get all players
// @Description Get all players
// @Tags Players
// @Accept  json
// @Produce  json
// @Success 200 {array} structs.Player
// @Router /api/v1/player/getall [get]
func GetAllPlayers(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Executa a consulta SQL
	rows, err := db.Query("SELECT * FROM Player")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Cria um slice de times
	players := []structs.Player{}

	// Lê os resultados
	for rows.Next() {
		// Cria um time vazio
		player := structs.Player{}

		// Preenche o time com os dados da linha
		err = rows.Scan(&player.Name, &player.City, &player.Country, &player.Age, &player.IdTeam, &player.ID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// Adiciona o time ao slice de players
		players = append(players, player)
	}

	// Fecha o conjunto de resultados
	rows.Close()

	// Converte os times em JSON
	playersJSON, err := json.Marshal(players)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Envia a resposta em JSON
	return c.JSONBlob(http.StatusOK, playersJSON)
}

// InsertTeam godoc
// @Summary Insert a team
// @Description Insert a team
// @Tags Teams
// @Accept  json
// @Produce  json
// @Success 200 {object} structs.Team
// @Router /api/v1/team/insert [post]
func InsertTeam(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Obtém os dados do time da URL
	name := c.QueryParam("name")
	city := c.QueryParam("city")
	country := c.QueryParam("country")

	// Executa a consulta SQL
	result, err := db.Exec("INSERT INTO Team (name, city, country) VALUES (?, ?, ?)", name, city, country)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Obtém o ID do time inserido
	id, err := result.LastInsertId()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Envia a resposta com o ID
	return c.String(http.StatusOK, fmt.Sprintf("Time inserido com o ID %d", id))
}

// UpdateTeam godoc
// @Summary Update a team
// @Description Update a team
// @Tags Teams
// @Accept  json
// @Produce  json
// @Success 200 {object} structs.Team
// @Router /api/v1/team/update [put]
func UpdateTeam(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Obtém os dados do time da URL
	id := c.QueryParam("id")
	name := c.QueryParam("name")
	city := c.QueryParam("city")
	country := c.QueryParam("country")

	// Executa a consulta SQL
	result, err := db.Exec("UPDATE Team SET name = ?, city = ?, country = ? WHERE id = ?", name, city, country, id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Obtém o número de linhas afetadas
	rows, err := result.RowsAffected()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Envia a resposta com o número de linhas afetadas
	return c.String(http.StatusOK, fmt.Sprintf("%d linha(s) afetada(s)", rows))
}

func DeleteTeam(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Obtém o ID do time da URL
	id := c.QueryParam("id")

	// Executa a consulta SQL
	result, err := db.Exec("DELETE FROM Team WHERE id = ?", id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Obtém o número de linhas afetadas
	rows, err := result.RowsAffected()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Envia a resposta com o número de linhas afetadas
	return c.String(http.StatusOK, fmt.Sprintf("%d linha(s) afetada(s)", rows))
}

// GetByIdTeam godoc
// @Summary Get a team by ID
// @Description Get a team by ID
// @Tags Teams
// @Accept  json
// @Produce  json
// @Param id path string true "Team ID"
// @Success 200 {object} structs.Team
// @Router /api/v1/team/getbyid/{id} [get]
func GetByIdTeam(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Obtém o valor do parâmetro id da URL da rota
	id := c.Param("id")

	// Executa a consulta SQL que seleciona o time com o ID informado
	row := db.QueryRow("SELECT * FROM Team WHERE id = ?", id)

	// Cria uma variável do tipo Team para armazenar os dados do time
	var team structs.Team

	// Lê o resultado da consulta e preenche a estrutura do time com os dados obtidos
	err = row.Scan(&team.ID, &team.Name, &team.City, &team.Country)
	if err != nil {
		// Lida com o erro
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Converte a estrutura do time em JSON
	teamJSON, err := json.Marshal(team)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Envia a resposta em JSON
	return c.JSONBlob(http.StatusOK, teamJSON)
}

// GetByIdPlayer godoc
// @Summary Get a player by ID
// @Description Get a player by ID
// @Tags Players
// @Accept  json
// @Produce  json
// @Param id path string true "Player ID"
// @Success 200 {object} structs.Player
// @Router /api/v1/player/getbyid/{id} [get]
func GetByIdPlayer(c echo.Context) error {
	// Connect to the database
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Get the value of the id parameter from the route URL
	id := c.Param("id")

	// Execute the SQL query that selects the player with the informed ID
	row := db.QueryRow("SELECT * FROM Player WHERE id = ?", id)

	// Create a variable of type Player to store the player data
	var player structs.Player

	// Read the result of the query and fill the player structure with the obtained data
	err = row.Scan(&player.Name, &player.City, &player.Country, &player.Age, &player.IdTeam, &player.ID)
	if err != nil {
		// Handle the error
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Convert the player structure into JSON
	playerJSON, err := json.Marshal(player)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Send the response in JSON
	return c.JSONBlob(http.StatusOK, playerJSON)
}

// GetByIdTeamPlayer godoc
// @Summary Get players by team ID
// @Description Get players by team ID
// @Tags Players
// @Accept  json
// @Produce  json
// @Param idteam path string true "Team ID"
// @Success 200 {array} structs.Player
// @Router /api/v1/player/getbyidteam/{idteam} [get]
func GetByIdTeamPlayer(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Obtém o valor do parâmetro country da URL da rota
	idteam := c.Param("idteam")

	// Executa a consulta SQL que seleciona todos os jogadores do time informado
	rows, err := db.Query("SELECT * FROM Player WHERE idteam = ?", idteam)
	if err != nil {
		// Lida com o erro
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	// Cria uma slice de Team para armazenar os dados dos times
	var players []structs.Player

	for rows.Next() {
		// Cria uma variável do tipo Player para cada linha do resultado
		var player structs.Player

		// Lê o resultado da consulta e preenche a estrutura do time com os dados obtidos
		err := rows.Scan(&player.Name, &player.City, &player.Country, &player.Age, &player.IdTeam, &player.ID)
		if err != nil {
			// Lida com o erro
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// Adiciona o jogador à slice de jogadores
		players = append(players, player)
	}

	// Converte a slice de jogadores em JSON
	playersJSON, err := json.Marshal(players)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Envia a resposta em JSON
	return c.JSONBlob(http.StatusOK, playersJSON)
}

// GetByCountryTeam godoc
// @Summary Get teams by country
// @Description Get teams by country
// @Tags Teams
// @Accept  json
// @Produce  json
// @Param country path string true "Country Name"
// @Success 200 {array} structs.Team
// @Router /api/v1/team/getbycountry/{country} [get]
func GetByCountryTeam(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Obtém o valor do parâmetro country da URL da rota
	country := c.Param("country")

	// Executa a consulta SQL que seleciona todos os times do país informado
	rows, err := db.Query("SELECT * FROM Team WHERE country = ?", country)
	if err != nil {
		// Lida com o erro
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	// Cria uma slice de Team para armazenar os dados dos times
	var teams []structs.Team

	for rows.Next() {
		// Cria uma variável do tipo Team para cada linha do resultado
		var team structs.Team

		// Lê o resultado da consulta e preenche a estrutura do time com os dados obtidos
		err := rows.Scan(&team.ID, &team.Name, &team.City, &team.Country)
		if err != nil {
			// Lida com o erro
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// Adiciona o time à slice de times
		teams = append(teams, team)
	}

	// Converte a slice de times em JSON
	teamsJSON, err := json.Marshal(teams)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Envia a resposta em JSON
	return c.JSONBlob(http.StatusOK, teamsJSON)
}

// GetByCountryPlayer godoc
// @Summary Get players by country
// @Description Get players by country
// @Tags Players
// @Accept  json
// @Produce  json
// @Param country path string true "Country Name"
// @Success 200 {array} structs.Player
// @Router /api/v1/player/getbycountry/{country} [get]
func GetByCountryPlayer(c echo.Context) error {
	// Connect to the database
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Get the value of the country parameter from the route URL
	country := c.Param("country")

	// Execute the SQL query that selects all players from the informed country
	rows, err := db.Query("SELECT * FROM Player WHERE Country = ?", country)
	if err != nil {
		// Handle the error
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	// Create a slice of Player to store the player data
	var players []structs.Player

	for rows.Next() {
		// Create a variable of type Player for each row of the result
		var player structs.Player

		// Read the result of the query and fill the player structure with the obtained data
		err := rows.Scan(&player.Name, &player.City, &player.Country, &player.Age, &player.IdTeam, &player.ID)
		if err != nil {
			// Handle the error
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// Add the player to the slice of players
		players = append(players, player)
	}

	// Convert the slice of players into JSON
	playersJSON, err := json.Marshal(players)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Send the response in JSON
	return c.JSONBlob(http.StatusOK, playersJSON)
}

// GetByNameTeam godoc
// @Summary Get a team by name
// @Description Get a team by name
// @Tags Teams
// @Accept  json
// @Produce  json
// @Param name path string true "Team Name"
// @Success 200 {object} structs.Team
// @Router /api/v1/team/getbyname/{name} [get]
func GetByNameTeam(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Obtém o valor do parâmetro name da URL da rota
	name := c.Param("name")

	// Executa a consulta SQL que seleciona o time com o nome informado
	row := db.QueryRow("SELECT * FROM Team WHERE name = ?", name)

	// Cria uma variável do tipo Team para armazenar os dados do time
	var team structs.Team

	// Lê o resultado da consulta e preenche a estrutura do time com os dados obtidos
	err = row.Scan(&team.ID, &team.Name, &team.City, &team.Country)
	if err != nil {
		// Lida com o erro
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Converte a estrutura do time em JSON
	teamJSON, err := json.Marshal(team)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Envia a resposta em JSON
	return c.JSONBlob(http.StatusOK, teamJSON)
}

// GetByNamePlayer godoc
// @Summary Get a player by name
// @Description Get a player by name
// @Tags Players
// @Accept  json
// @Produce  json
// @Param name path string true "Player Name"
// @Success 200 {object} structs.Player
// @Router /api/v1/player/getbyname/{name} [get]
func GetByNamePlayer(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Obtém o valor do parâmetro name da URL da rota
	name := c.Param("name")

	// Executa a consulta SQL que seleciona o player com o nome informado
	row := db.QueryRow("SELECT * FROM Player WHERE name = ?", name)

	// Cria uma variável do tipo player para armazenar os dados do time
	var player structs.Player

	// Lê o resultado da consulta e preenche a estrutura do time com os dados obtidos
	err = row.Scan(&player.Name, &player.City, &player.Country, &player.Age, &player.IdTeam, &player.ID)
	if err != nil {
		// Lida com o erro
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Converte a estrutura do time em JSON
	teamJSON, err := json.Marshal(player)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Envia a resposta em JSON
	return c.JSONBlob(http.StatusOK, teamJSON)
}

// InsertPlayer godoc
// @Summary Insert a player
// @Description Insert a player
// @Tags Players
// @Accept  json
// @Produce  json
// @Param name query string true "Player Name"
// @Param team query string true "Team ID"
// @Param city query string true "City"
// @Param country query string true "Country"
// @Param age query string true "Age"
// @Success 200 {object} structs.Player
// @Router /api/v1/player/insert [post]
func InsertPlayer(c echo.Context) error {
	// Connect to the database
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Get the player data from the URL
	name := c.QueryParam("name")
	idteam := c.QueryParam("team")
	city := c.QueryParam("city")
	country := c.QueryParam("country")
	age := c.QueryParam("age")

	// Execute the SQL query
	result, err := db.Exec("INSERT INTO Player (Name, IdTeam, City, Country, Age) VALUES (?, ?, ?, ?, ?)", name, idteam, city, country, age)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Get the ID of the inserted player
	id, err := result.LastInsertId()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Send the response with the ID
	return c.String(http.StatusOK, fmt.Sprintf("Player inserted with ID %d", id))
}
