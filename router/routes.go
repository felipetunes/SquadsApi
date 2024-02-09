package router

import (
	"apiSquads/db"
	"apiSquads/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
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

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Executa a consulta SQL
	rows, err := db.Query("SELECT Name, City, Country, ID, COALESCE(Color1, '') as Color1 FROM Team ORDER BY Name")
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
		err = rows.Scan(&team.Name, &team.City, &team.Country, &team.ID, &team.Color1)
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

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Executa a consulta SQL
	rows, err := db.Query("SELECT Name, City, Country, Birth, IdTeam, ID, COALESCE(Height, '0.00') as Height, COALESCE(Position, '') as Position FROM Player ORDER BY Name")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Cria um slice de jogadores
	players := []structs.Player{}

	// Lê os resultados
	for rows.Next() {
		// Cria um jogador vazio
		player := structs.Player{}

		// Preenche o jogador com os dados da linha
		var birthStr string
		err = rows.Scan(&player.Name, &player.City, &player.Country, &birthStr, &player.IdTeam, &player.ID, &player.Height, &player.Position)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		if birthStr != "0000-00-00" {
			player.Birth, err = time.Parse("2006-01-02", birthStr)
			if err != nil {
				// Handle the error
				return c.String(http.StatusInternalServerError, err.Error())
			}
		}

		// Adiciona o jogador ao slice de jogadores
		players = append(players, player)
	}

	// Fecha o conjunto de resultados
	rows.Close()

	// Converte os jogadores em JSON
	playersJSON, err := json.Marshal(players)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Envia a resposta em JSON
	return c.JSONBlob(http.StatusOK, playersJSON)
}

// GetAllLives godoc
// @Summary Get all matches
// @Description Get all matches
// @Tags Lives
// @Accept  json
// @Produce  json
// @Success 200 {array} structs.Live
// @Router /api/v1/live/getall [get]
func GetAllMatches(c echo.Context) error {
	var dateMatch string
	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Executa a consulta SQL
	rows, err := db.Query("SELECT * FROM Live ORDER BY DateMatch")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Cria um slice de partidas ao vivo
	lives := []structs.Live{}

	// Lê os resultados
	for rows.Next() {
		live := structs.Live{}
		err = rows.Scan(&live.ID, &live.IdTeam1, &live.IdTeam2, &live.IdChampionship, &dateMatch, &live.Stadium, &live.TeamPoints1, &live.TeamPoints2)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		live.DateMatch, err = time.Parse("2006-01-02 15:04:05", dateMatch)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		lives = append(lives, live)
	}
	// Fecha o conjunto de resultados
	rows.Close()

	// Converte as partidas ao vivo em JSON
	livesJSON, err := json.Marshal(lives)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Envia a resposta em JSON
	return c.JSONBlob(http.StatusOK, livesJSON)
}

// User godoc
// @Summary Login user
// @Description Login user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param   username     query    string     true        "Username"
// @Param   password     query    string     true        "Password"
// @Success 200 {object} structs.User
// @Router /api/v1/user/login [post]
func Login(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Pega o nome de usuário e a senha do corpo da solicitação
	u := new(structs.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	// Verifica se o nome de usuário e a senha foram fornecidos
	if u.Username == "" || u.Password == "" {
		return c.String(http.StatusBadRequest, "Username and password must be provided")
	}

	// Executa a consulta SQL
	row := db.QueryRow("SELECT * FROM Users WHERE Username = ?", u.Username)

	// Lê o resultado
	userInDb := structs.User{}
	err = row.Scan(&userInDb.ID, &userInDb.Username, &userInDb.Password)
	if err == sql.ErrNoRows {
		// Nenhum usuário com o nome de usuário fornecido foi encontrado
		return echo.ErrUnauthorized
	} else if err != nil {
		// Um erro diferente ocorreu
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Verifica a senha
	err = bcrypt.CompareHashAndPassword([]byte(userInDb.Password), []byte(u.Password))
	if err == nil {
		// A senha está correta, então faça o login do usuário
		return c.JSON(http.StatusOK, userInDb)
	}

	// Se chegamos até aqui, o login falhou
	return echo.ErrUnauthorized
}

// User godoc
// @Summary Register user
// @Description Register user
// @Tags Users
// @Accept json
// @Produce json
// @Param username query string true "Username"
// @Param password query string true "Password"
// @Success 200 {object} structs.User
// @Router /api/v1/user/register [post]
func Register(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Pega o nome de usuário e a senha do corpo da solicitação
	u := new(structs.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	// Verifica se o nome de usuário e a senha foram fornecidos
	if u.Username == "" || u.Password == "" {
		return c.String(http.StatusBadRequest, "Username and password must be provided")
	}

	// Gera o hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Executa a consulta SQL para inserir o novo usuário
	_, err = db.Exec("INSERT INTO Users (Username, Password) VALUES (?, ?)", u.Username, hashedPassword)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Retorna uma mensagem de sucesso
	return c.JSON(http.StatusOK, map[string]string{
		"message": "User registered successfully",
	})
}

// GetAllByIdTeam godoc
// @Summary Get all matches by team id
// @Description Get all matches by team id
// @Tags Lives
// @Accept  json
// @Produce  json
// @Param id path string true "Team ID"
// @Success 200 {array} structs.Live
// @Router /api/v1/live/getallbyidteam/{id} [get]
func GetAllByIdTeam(c echo.Context) error {
	var dateMatch string

	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Obtém o IdTeam do parâmetro de rota
	idTeam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Prepara a consulta SQL
	stmt, err := db.Prepare("SELECT * FROM Live WHERE IdTeam1 = ? OR IdTeam2 = ?")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Executa a consulta SQL
	rows, err := stmt.Query(idTeam, idTeam)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Cria um slice de partidas ao vivo
	lives := []structs.Live{}

	// Lê os resultados
	for rows.Next() {
		live := structs.Live{}
		err = rows.Scan(&live.ID, &live.IdTeam1, &live.IdTeam2, &live.IdChampionship, &dateMatch, &live.Stadium, &live.TeamPoints1, &live.TeamPoints2)
		live.DateMatch, err = time.Parse("2006-01-02 15:04:05", dateMatch)

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		lives = append(lives, live)
	}
	// Fecha o conjunto de resultados
	rows.Close()

	// Verifica se o slice de partidas ao vivo está vazio
	if len(lives) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Não há jogos encontrados")
	}

	// Converte as partidas ao vivo em JSON
	livesJSON, err := json.Marshal(lives)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Envia a resposta em JSON
	return c.JSONBlob(http.StatusOK, livesJSON)
}

// GetAllLivesToday godoc
// @Summary Get all matches today
// @Description Get all matches today
// @Tags Lives
// @Accept  json
// @Produce  json
// @Success 200 {array} structs.Live
// @Router /api/v1/live/getalltoday [get]
func GetAllLivesToday(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Obtém o valor do CURDATE()
	var curDate string
	err = db.QueryRow("SELECT CURDATE()").Scan(&curDate)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Executa a consulta SQL
	rows, err := db.Query("SELECT ID, IdTeam1, IdTeam2, IdChampionship, DATE_FORMAT(DateMatch, '%Y-%m-%d %H:%i:%s') as DateMatch, Stadium, TeamPoints1, TeamPoints2 FROM Live WHERE DATE(DateMatch) = CURDATE() ORDER BY DateMatch")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Cria um slice de partidas ao vivo
	lives := []structs.Live{}

	// Lê os resultados
	for rows.Next() {
		// Cria uma partida ao vivo vazia
		live := structs.Live{}

		// Cria uma variável para armazenar a data e hora como string
		var dateMatch string

		// Preenche a partida ao vivo com os dados da linha
		err = rows.Scan(&live.ID, &live.IdTeam1, &live.IdTeam2, &live.IdChampionship, &dateMatch, &live.Stadium, &live.TeamPoints1, &live.TeamPoints2)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// Converte a string de data e hora para time.Time
		live.DateMatch, err = time.Parse("2006-01-02 15:04:05", dateMatch)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// Adiciona a partida ao vivo ao slice de partidas ao vivo
		lives = append(lives, live)
	}
	// Fecha o conjunto de resultados
	rows.Close()

	// Converte as partidas ao vivo em JSON
	livesJSON, err := json.Marshal(lives)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Envia a resposta em JSON
	return c.JSONBlob(http.StatusOK, livesJSON)
}

// InsertLive godoc
// @Summary Insert a live match
// @Description Insert a live match
// @Tags Lives
// @Accept  json
// @Produce  json
// @Param idteam1 query int true "Team ID 1"
// @Param idteam2 query int true "Team ID 2"
// @Param idchampionship query int true "IdChampionship"
// @Param datematch query string true "Date of Match"
// @Param stadium query string true "Stadium"
// @Param teampoints1 query int true "Team Points 1"
// @Param teampoints2 query int true "Team Points 2"
// @Success 200 {object} structs.Live
// @Router /api/v1/live/insert [post]
func InsertLive(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Obtém os dados da partida ao vivo da URL
	idteam1, _ := strconv.Atoi(c.QueryParam("idteam1"))
	idteam2, _ := strconv.Atoi(c.QueryParam("idteam2"))
	idchampionship, _ := strconv.Atoi(c.QueryParam("idchampionship"))
	datematch := c.QueryParam("datematch")
	stadium := c.QueryParam("stadium")
	teampoints1, _ := strconv.Atoi(c.QueryParam("teampoints1"))
	teampoints2, _ := strconv.Atoi(c.QueryParam("teampoints2"))

	// Converte a string datematch para o tipo time.Time
	layout := "02/01/2006 15:04"
	parsedDate, err := time.Parse(layout, datematch)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Cria um novo struct Live com os dados obtidos
	live := structs.Live{
		ID:             0,
		IdTeam1:        idteam1,
		IdTeam2:        idteam2,
		IdChampionship: idchampionship,
		DateMatch:      parsedDate,
		Stadium:        stadium,
		TeamPoints1:    teampoints1,
		TeamPoints2:    teampoints2,
	}

	// Executa a consulta SQL
	_, err = db.Exec("INSERT INTO Live (id, idteam1, idteam2, idchampionship, datematch, stadium, teampoints1, teampoints2) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", live.ID, live.IdTeam1, live.IdTeam2, live.IdChampionship, live.DateMatch, live.Stadium, live.TeamPoints1, live.TeamPoints2)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Envia a resposta
	return c.String(http.StatusOK, "Partida ao vivo inserida com sucesso.")
}

// UpdateLive godoc
// @Summary Update a live match
// @Description Update a live match
// @Tags Lives
// @Accept  json
// @Produce  json
// @Param id query int true "Match ID"
// @Param idteam1 query int true "Team ID 1"
// @Param idteam2 query int true "Team ID 2"
// @Param idchampionship query int true "Championship ID"
// @Param datematch query string true "Date of Match"
// @Param stadium query string true "Stadium"
// @Param teampoints1 query int true "Team Points 1"
// @Param teampoints2 query int true "Team Points 2"
// @Success 200 {string} string "Partida ao vivo atualizada com sucesso."
// @Router /api/v1/live/update [put]
func UpdateLive(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Obtém os dados da partida ao vivo da URL
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "ID inválido")
	}
	idteam1, err := strconv.Atoi(c.QueryParam("idteam1"))
	if err != nil {
		return c.String(http.StatusBadRequest, "ID da equipe 1 inválido")
	}
	idteam2, err := strconv.Atoi(c.QueryParam("idteam2"))
	if err != nil {
		return c.String(http.StatusBadRequest, "ID da equipe 2 inválido")
	}
	idchampionship, err := strconv.Atoi(c.QueryParam("idchampionship"))
	if err != nil {
		return c.String(http.StatusBadRequest, "ID do campeonato inválido")
	}
	datematch := c.QueryParam("datematch")
	stadium := c.QueryParam("stadium")
	teampoints1, err := strconv.Atoi(c.QueryParam("teampoints1"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Pontos da equipe 1 inválidos")
	}
	teampoints2, err := strconv.Atoi(c.QueryParam("teampoints2"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Pontos da equipe 2 inválidos")
	}

	// Converte a string datematch para o tipo time.Time
	layout := "02/01/2006 15:04"
	parsedDate, err := time.Parse(layout, datematch)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Verifica se a partida ao vivo existe
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM Live WHERE id = ?", id).Scan(&count)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if count == 0 {
		return c.String(http.StatusNotFound, "Partida ao vivo não encontrada")
	}

	// Cria um novo struct Live com os dados obtidos
	live := structs.Live{
		ID:             id,
		IdTeam1:        idteam1,
		IdTeam2:        idteam2,
		IdChampionship: idchampionship,
		DateMatch:      parsedDate,
		Stadium:        stadium,
		TeamPoints1:    teampoints1,
		TeamPoints2:    teampoints2,
	}

	// Executa a consulta SQL
	_, err = db.Exec("UPDATE Live SET idteam1 = ?, idteam2 = ?, idchampionship = ?, datematch = ?, stadium = ?, teampoints1 = ?, teampoints2 = ? WHERE id = ?", live.IdTeam1, live.IdTeam2, live.IdChampionship, live.DateMatch, live.Stadium, live.TeamPoints1, live.TeamPoints2, live.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Envia a resposta
	return c.String(http.StatusOK, "Partida ao vivo atualizada com sucesso.")
}

// InsertTeam godoc
// @Summary Insert a team
// @Description Insert a team
// @Tags Teams
// @Accept  json
// @Produce  json
// @Param name query string true "Team Name"
// @Param city query string true "City"
// @Param country query string true "Country"
// @Param color1 query string true "Color1"
// @Success 200 {object} structs.Team
// @Router /api/v1/team/insert [post]
func InsertTeam(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Obtém os dados do time da URL
	name := c.QueryParam("name")
	city := c.QueryParam("city")
	country := c.QueryParam("country")
	color1 := c.QueryParam("color1")

	// Executa a consulta SQL
	result, err := db.Exec("INSERT INTO Team (name, city, country, Color1 VALUES (?, ?, ?, ?)", name, city, country, color1)
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

func UpdateTeam(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Obtém os dados do time da URL
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "ID inválido")
	}
	name := c.QueryParam("name")
	city := c.QueryParam("city")
	country := c.QueryParam("country")
	color1 := c.QueryParam("color1")

	// Verifica se os campos obrigatórios estão presentes
	if name == "" || city == "" || country == "" || color1 == "" {
		return c.String(http.StatusBadRequest, "Nome, cidade, país e cor1 são obrigatórios")
	}

	// Verifica se o time existe
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM Team WHERE id = ?", id).Scan(&count)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if count == 0 {
		return c.String(http.StatusNotFound, "Time não encontrado")
	}

	// Executa a consulta SQL
	result, err := db.Exec("UPDATE Team SET name = ?, city = ?, country = ?, color1 = ? WHERE id = ?", name, city, country, color1, id)
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

// DeletePlayer godoc
// @Summary Delete a player by ID
// @Description Delete a player by ID
// @Tags Players
// @Accept  json
// @Produce  json
// @Param id path string true "Player ID"
// @Success 200 {object} string
// @Router /api/v1/player/delete/{id} [delete]
func DeletePlayer(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Obtém o ID do jogador da URL
	id := c.Param("id")

	// Executa a consulta SQL
	result, err := db.Exec("DELETE FROM Player WHERE id = ?", id)
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

// DeleteTeam godoc
// @Summary Delete a team by ID
// @Description Delete a team by ID
// @Tags Teams
// @Accept  json
// @Produce  json
// @Param id path string true "Team ID"
// @Success 200 {object} string
// @Router /api/v1/team/delete/{id} [delete]
func DeleteTeam(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Obtém o ID do time da URL
	id := c.Param("id")

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

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Obtém o valor do parâmetro id da URL da rota
	id := c.Param("id")

	// Executa a consulta SQL que seleciona o time com o ID informado
	row := db.QueryRow("SELECT Name, City, Country, ID, COALESCE(Color1, '') as Color1 FROM Team WHERE id = ?", id)

	// Cria uma variável do tipo Team para armazenar os dados do time
	var team structs.Team

	// Lê o resultado da consulta e preenche a estrutura do time com os dados obtidos
	err = row.Scan(&team.Name, &team.City, &team.Country, &team.ID, &team.Color1)
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

// GetByIdLive godoc
// @Summary Get a live match by ID
// @Description Get a live match by ID
// @Tags Lives
// @Accept  json
// @Produce  json
// @Param id path string true "Live ID"
// @Success 200 {object} structs.Live
// @Router /api/v1/live/getbyid/{id} [get]
func GetByIdLive(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Obtém o valor do parâmetro id da URL da rota
	id := c.Param("id")

	// Executa a consulta SQL que seleciona a partida ao vivo com o ID informado
	row := db.QueryRow("SELECT * FROM Live WHERE id = ?", id)

	// Cria uma variável do tipo Live para armazenar os dados da partida ao vivo
	var live structs.Live

	// Lê o resultado da consulta e preenche a estrutura da partida ao vivo com os dados obtidos
	err = row.Scan(&live.ID, &live.IdTeam1, &live.IdTeam2, &live.IdChampionship, &live.DateMatch, &live.Stadium, &live.StatusMatch, &live.GameTime, &live.TeamPoints1, &live.TeamPoints2, &live.HomeTeamWins, &live.VisitingTeamWins, &live.Draws, &live.HomeTeamRecentPerformance, &live.VisitingTeamRecentPerformance, &live.HomeTeamOdds, &live.VisitingTeamOdds, &live.DrawOdds)
	if err != nil {
		// Lida com o erro
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Converte a estrutura da partida ao vivo em JSON
	liveJSON, err := json.Marshal(live)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Envia a resposta em JSON
	return c.JSONBlob(http.StatusOK, liveJSON)
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

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Get the value of the id parameter from the route URL
	id := c.Param("id")

	// Execute the SQL query that selects the player with the informed ID
	row := db.QueryRow("SELECT Name, City, Country, Birth, IdTeam, ID, COALESCE(Height, '0.00') as Height, COALESCE(Position, '') as Position FROM Player WHERE ID = ?", id)

	// Create a variable of type Player to store the player data
	var player structs.Player

	// Read the result of the query and fill the player structure with the obtained data
	var birthStr string
	err = row.Scan(&player.Name, &player.City, &player.Country, &birthStr, &player.IdTeam, &player.ID, &player.Height, &player.Position)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if birthStr != "0000-00-00" {
		player.Birth, err = time.Parse("2006-01-02", birthStr)
		if err != nil {
			// Handle the error
			return c.String(http.StatusInternalServerError, err.Error())
		}
	} else {
		// Handle '0000-00-00' birth date here
		// For example, you can leave player.Birth as zero value (which is '0001-01-01 00:00:00 +0000 UTC' for time.Time)
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
// @Param idteam path string true "IdTeam"
// @Success 200 {array} structs.Player
// @Router /api/v1/player/getbyidteam/{idteam} [get]
func GetByIdTeamPlayer(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Obtém o valor do parâmetro idteam da URL da rota
	idteam := c.Param("idteam")

	// Executa a consulta SQL que seleciona todos os jogadores do time informado
	rows, err := db.Query("SELECT Name, City, Country, Birth, IdTeam, ID, COALESCE(Height, '0.00') as Height, COALESCE(Position, '') as Position FROM Player WHERE IdTeam = ?", idteam)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Cria um slice para armazenar os jogadores
	players := []structs.Player{}

	// Lê os resultados
	for rows.Next() {
		// Cria um jogador vazio
		var player structs.Player

		// Preenche o jogador com os dados da linha
		var birthStr string
		err = rows.Scan(&player.Name, &player.City, &player.Country, &birthStr, &player.IdTeam, &player.ID, &player.Height, &player.Position)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		if birthStr != "0000-00-00" {
			player.Birth, err = time.Parse("2006-01-02", birthStr)
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
		}

		// Adiciona o jogador ao slice de jogadores
		players = append(players, player)
	}

	// Fecha o conjunto de resultados
	rows.Close()

	// Converte o slice de jogadores em JSON
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

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

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

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Get the value of the country parameter from the route URL
	country := c.Param("country")

	// Execute the SQL query that selects all players from the informed country
	rows, err := db.Query("SELECT Name, City, Country, Birth, IdTeam, ID, COALESCE(Height, '0.00') as Height, COALESCE(Position, '') as Position FROM Player WHERE Country = ?", country)
	if err != nil {
		// Handle the error
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	// Create a slice of Player to store the player data
	var players []structs.Player

	for rows.Next() {
		// Cria um jogador vazio
		player := structs.Player{}

		// Preenche o jogador com os dados da linha
		var birthStr string
		err = rows.Scan(&player.Name, &player.City, &player.Country, &birthStr, &player.IdTeam, &player.ID, &player.Height, &player.Position)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		if birthStr != "0000-00-00" {
			player.Birth, err = time.Parse("2006-01-02", birthStr)
			if err != nil {
				// Handle the error
				return c.String(http.StatusInternalServerError, err.Error())
			}
		}

		// Adiciona o jogador ao slice de jogadores
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
// @Summary Get teams by name
// @Description Get teams by name
// @Tags Teams
// @Accept  json
// @Produce  json
// @Param name path string true "Team Name"
// @Success 200 {array} structs.Team
// @Router /api/v1/team/getbyname/{name} [get]
func GetTeamsByName(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Obtém o valor do parâmetro name da URL da rota
	name := c.Param("name")

	// Executa a consulta SQL que seleciona os times com o nome informado
	rows, err := db.Query("SELECT Name, City, Country, Id, COALESCE(Color1, '') as Color1 FROM Team WHERE name LIKE ?", "%"+name+"%")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	// Cria uma lista para armazenar os times
	var teams []structs.Team

	// Lê o resultado da consulta e preenche a lista de times com os dados obtidos
	for rows.Next() {
		var team structs.Team
		err = rows.Scan(&team.Name, &team.City, &team.Country, &team.ID, &team.Color1)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		teams = append(teams, team)
	}

	// Converte a lista de times em JSON
	teamsJSON, err := json.Marshal(teams)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Envia a resposta em JSON
	return c.JSONBlob(http.StatusOK, teamsJSON)
}

// GetPlayersByName godoc
// @Summary Get players by name
// @Description Get players by name
// @Tags Players
// @Accept  json
// @Produce  json
// @Param name path string true "Player Name"
// @Success 200 {array} structs.Player
// @Router /api/v1/player/getbyname/{name} [get]
func GetPlayersByName(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Obtém o valor do parâmetro name da URL da rota
	name := c.Param("name")

	// Executa a consulta SQL que seleciona os players cujo nome contém o texto informado
	rows, err := db.Query("SELECT Name, City, Country, Birth, IdTeam, ID, COALESCE(Height, '0.00') as Height, COALESCE(Position, '') as Position FROM Player WHERE name LIKE ?", "%"+name+"%")
	if err != nil {
		// Lida com o erro
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	// Cria uma lista para armazenar os jogadores
	var players []structs.Player

	for rows.Next() {
		// Cria uma variável do tipo player para armazenar os dados do jogador
		var player structs.Player

		// Lê o resultado da consulta e preenche a estrutura do jogador com os dados obtidos
		var birthStr string
		err = rows.Scan(&player.Name, &player.City, &player.Country, &birthStr, &player.IdTeam, &player.ID, &player.Height, &player.Position)
		if err != nil {
			// Lida com o erro
			return c.String(http.StatusInternalServerError, err.Error())
		}
		if birthStr != "0000-00-00" {
			player.Birth, err = time.Parse("2006-01-02", birthStr)
			if err != nil {
				// Handle the error
				return c.String(http.StatusInternalServerError, err.Error())
			}
		} else {
			// Handle '0000-00-00' birth date here
			// For example, you can leave player.Birth as zero value (which is '0001-01-01 00:00:00 +0000 UTC' for time.Time)
		}

		// Adiciona o jogador à lista
		players = append(players, player)
	}

	// Converte a lista de jogadores em JSON
	playersJSON, err := json.Marshal(players)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Envia a resposta em JSON
	return c.JSONBlob(http.StatusOK, playersJSON)
}

// InsertPlayer godoc
// @Summary Insert a player
// @Description Insert a player
// @Tags Players
// @Accept  json
// @Produce  json
// @Param name query string true "Player Name"
// @Param idteam query string true "Id Team"
// @Param city query string true "City"
// @Param country query string true "Country"
// @Param birth query string true "Birth" example="DD-MM-YYYY"
// @Param height query string true "Height"
// @Param position query string true "Position"
// @Success 200 {object} structs.Player
// @Router /api/v1/player/insert [post]
func InsertPlayer(c echo.Context) error {
	// Connect to the database
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Get the player data from the URL
	name := c.QueryParam("name")
	idteam := c.QueryParam("idteam")
	city := c.QueryParam("city")
	country := c.QueryParam("country")
	birth := c.QueryParam("birth")
	height := c.QueryParam("height")
	position := c.QueryParam("position")

	// Converte a string de data de nascimento para o tipo date
	birthDate, err := time.Parse("02/01/2006", birth)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Formata a data de nascimento para o formato aceito pelo MySQL
	formattedBirthDate := birthDate.Format("2006-01-02")

	// Execute the SQL query
	result, err := db.Exec("INSERT INTO Player (Name, IdTeam, City, Country, Birth, Height, Position) VALUES (?, ?, ?, ?, ?, ?, ?)", name, idteam, city, country, formattedBirthDate, height, position)
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

// UpdatePlayer godoc
// @Summary Update a player
// @Description Update a player
// @Tags Players
// @Accept  json
// @Produce  json
// @Param id query int true "ID Player"
// @Param name query string true "Player Name"
// @Param idteam query int true "Id Team"
// @Param city query string true "City"
// @Param country query string true "Country"
// @Param birth query string true "Birth" example="DD-MM-AAAA"
// @Param height query string true "Height"
// @Param position query string true "Position"
// @Success 200 {object} structs.Player
// @Router /api/v1/player/update [put]
func UpdatePlayer(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Obtém os dados do jogador da URL
	id := c.QueryParam("id")
	name := c.QueryParam("name")
	idteam := c.QueryParam("idteam")
	city := c.QueryParam("city")
	country := c.QueryParam("country")
	birth := c.QueryParam("birth")
	height := c.QueryParam("height")
	position := c.QueryParam("position")

	// Converte a string de data de nascimento para o tipo date
	var birthDate time.Time
	if birth != "" {
		birthDate, err = time.Parse("02/01/2006", birth)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
	}

	// Formata a data de nascimento para o formato aceito pelo MySQL
	formattedBirthDate := birthDate.Format("2006-01-02")

	// Executa a consulta SQL
	result, err := db.Exec("UPDATE Player SET name = ?, city = ?, country = ?, birth = ?, idteam = ?, height = ?, position = ? WHERE id = ?", name, city, country, formattedBirthDate, idteam, height, position, id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Obtém o número de linhas afetadas
	rows, err := result.RowsAffected()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Envia a resposta com o número de linhas afetadas
	return c.String(http.StatusOK, fmt.Sprintf("%d linha(s) afetada(s)", rows))
}

// GetByChampionship godoc
// @Summary Get teams by championship
// @Description Get teams by a given championship ID
// @Tags Teams
// @Accept  json
// @Produce  json
// @Param id query int true "ID Championship"
// @Success 200 {array} structs.Team
// @Router /api/v1/team/getbychampionship/{id} [get]
func GetByChampionship(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Obtém o IdChampionship da URL
	idChampionship := c.Param("id")
	// Consulta para obter todos os IdTeam associados ao IdChampionship
	rows, err := db.Query("SELECT IdTeam FROM TeamChampionships WHERE IdChampionship = ?", idChampionship)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	var teams []structs.Team
	for rows.Next() {
		var idTeam int
		err := rows.Scan(&idTeam)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// Consulta para obter a equipe com o IdTeam
		var team structs.Team
		err = db.QueryRow("SELECT Name, City, Country, ID, COALESCE(Color1, '') FROM Team WHERE id = ?", idTeam).Scan(&team.Name, &team.City, &team.Country, &team.ID, &team.Color1)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("No team found with the given id.")
			} else {
				return c.String(http.StatusInternalServerError, err.Error())
			}
		}

		teams = append(teams, team)
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Envia a resposta com as equipes encontradas
	return c.JSON(http.StatusOK, teams)
}
