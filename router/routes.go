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

// Função que retorna todos os times
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

// Função que retorna todos os times
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
		err = rows.Scan(&player.ID, &player.Name, &player.City, &player.Country, &player.Age, &player.IdTeam)
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

// Função que insere um novo time
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

// Função que atualiza um time existente
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

// Função que deleta um time existente
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

// Função que retorna um time pelo seu ID
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

// Função que retorna os jogadores pelo ID do time
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
		err := rows.Scan(&player.ID, &player.Name, &player.City, &player.Country, &player.Age, &player.IdTeam)
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

// Função que retorna um time pelo país
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

// Função que retorna um time pelo seu nome
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

// Função que retorna um time pelo seu nome
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
	err = row.Scan(&player.ID, &player.Name, &player.City, &player.Country, &player.Age, &player.IdTeam)
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
