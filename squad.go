package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

// Estrutura que representa um time
type Team struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
}

// Função que conecta ao banco de dados
func connectDB() (*sql.DB, error) {
	// Abre a conexão com o banco de dados
	db, err := sql.Open("mysql", "admin:Tunes1313#@tcp(teammate.cr2mw0ioqij3.us-east-1.rds.amazonaws.com:3306)/Teammate")
	if err != nil {
		return nil, err
	}

	// Testa se a conexão está ativa
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Retorna a conexão como resultado
	return db, nil
}

// Função que retorna todos os times
func getAllTeams(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := connectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Executa a consulta SQL
	rows, err := db.Query("SELECT * FROM Team")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Cria um slice de times
	teams := []Team{}

	// Lê os resultados
	for rows.Next() {
		// Cria um time vazio
		team := Team{}

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

// Função que insere um novo time
func insertTeam(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := connectDB()
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
func updateTeam(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := connectDB()
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
func deleteTeam(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := connectDB()
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
func getByIdTeam(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := connectDB()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Obtém o valor do parâmetro id da URL da rota
	id := c.Param("id")

	// Executa a consulta SQL que seleciona o time com o ID informado
	row := db.QueryRow("SELECT * FROM Team WHERE id = ?", id)

	// Cria uma variável do tipo Team para armazenar os dados do time
	var team Team

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

func getByCountryTeam(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := connectDB()

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
	var teams []Team

	for rows.Next() {
		// Cria uma variável do tipo Team para cada linha do resultado
		var team Team

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

func getByNameTeam(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := connectDB()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Obtém o valor do parâmetro name da URL da rota
	name := c.Param("name")

	// Executa a consulta SQL que seleciona o time com o nome informado
	row := db.QueryRow("SELECT * FROM Team WHERE name = ?", name)

	// Cria uma variável do tipo Team para armazenar os dados do time
	var team Team

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

func main() {
	// Cria uma nova instância do Echo
	e := echo.New()

	// Define as rotas da aplicação
	e.GET("team/getall", getAllTeams)
	e.GET("team/insert", insertTeam)
	e.GET("team/update", updateTeam)
	e.GET("team/delete", deleteTeam)
	e.GET("team/getbyid/:id", getByIdTeam)
	e.GET("team/getbyname/:name", getByNameTeam)
	e.GET("team/getbycountry/:country", getByCountryTeam)

	// Inicia o servidor na porta 8080
	e.Logger.Fatal(e.Start(":8080"))
}
