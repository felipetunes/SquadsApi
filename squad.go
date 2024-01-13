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
	db, err := sql.Open("mysql", "root:Tunes1313#@tcp(localhost:3306)/TeamDB")
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
func getAll(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := connectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Executa a consulta SQL
	rows, err := db.Query("SELECT * FROM team")
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
func insert(c echo.Context) error {
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
	result, err := db.Exec("INSERT INTO team (name, city, country) VALUES (?, ?, ?)", name, city, country)
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
func update(c echo.Context) error {
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
	result, err := db.Exec("UPDATE team SET name = ?, city = ?, country = ? WHERE id = ?", name, city, country, id)
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
func delete(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := connectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Obtém o ID do time da URL
	id := c.QueryParam("id")

	// Executa a consulta SQL
	result, err := db.Exec("DELETE FROM team WHERE id = ?", id)
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
func getById(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := connectDB()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Obtém o valor do parâmetro id da URL da rota
	id := c.Param("id")

	fmt.Printf(id)

	// Executa a consulta SQL que seleciona o time com o ID informado
	row := db.QueryRow("SELECT * FROM team WHERE id = ?", id)

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
	e.GET("/getall", getAll)
	e.GET("/insert", insert)
	e.GET("/update", update)
	e.GET("/delete", delete)
	e.GET("/getbyid/:id", getById)

	// Inicia o servidor na porta 8080
	e.Logger.Fatal(e.Start(":8080"))
}
