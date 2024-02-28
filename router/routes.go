package router

import (
	"apiSquads/db"
	"apiSquads/structs"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"bytes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// GetAllBetsByUserId godoc
// @Summary Get all bets by user ID
// @Description Get all bets by user ID
// @Tags Bets
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {array} structs.Bet
// @Router /api/v1/bet/getallbyuserid/{id} [get]
func GetAllBetsByUserId(c echo.Context) error {
	// Connect to the database
	db, err := db.ConnectDB()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Make sure the connection will be closed at the end of this function
	defer db.Close()

	// Get the value of the id parameter from the URL route
	id := c.Param("id")

	// Execute the SQL query that selects all bets with the provided user ID
	rows, err := db.Query("SELECT Id, MatchId, SelectedOutcome, BetAmount, PossibleReturn FROM Bet WHERE UserId = ?", id)
	if err != nil {
		// Handle the error
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Create a slice of Bet to store the bets data
	var bets []structs.Bet

	// Loop through the rows and fill the bets slice with the obtained data
	for rows.Next() {
		var bet structs.Bet
		err = rows.Scan(&bet.Id, &bet.MatchId, &bet.SelectedOutcome, &bet.BetAmount, &bet.PossibleReturn)
		if err != nil {
			// Handle the error
			return c.String(http.StatusInternalServerError, err.Error())
		}
		bets = append(bets, bet)
	}

	// Convert the bets slice into JSON
	betsJSON, err := json.Marshal(bets)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Send the response in JSON
	return c.JSONBlob(http.StatusOK, betsJSON)
}

// UpdateBet godoc
// @Summary Update a bet
// @Description Update a bet
// @Tags Bets
// @Accept  json
// @Produce  json
// @Param id query int true "Bet ID"
// @Param userid query int true "User ID"
// @Param matchid query int true "Match ID"
// @Param amount query float64 true "Amount"
// @Param prediction query string true "Prediction"
// @Success 200 {object} structs.Bet
// @Router /api/v1/bet/update [put]
func UpdateBet(c echo.Context) error {
	// Connect to the database
	db, err := db.ConnectDB()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Make sure the connection will be closed at the end of this function
	defer db.Close()

	// Get the values of the parameters from the request
	id, _ := strconv.Atoi(c.QueryParam("id"))
	userid, _ := strconv.Atoi(c.QueryParam("userid"))
	matchid, _ := strconv.Atoi(c.QueryParam("matchid"))
	amount, _ := strconv.ParseFloat(c.QueryParam("amount"), 64)
	prediction := c.QueryParam("prediction")

	// Execute the SQL query that updates the bet with the provided parameters
	_, err = db.Exec("UPDATE Bet SET UserId = ?, MatchId = ?, BetAmount = ?, SelectedOutcome = ? WHERE Id = ?", userid, matchid, amount, prediction, id)
	if err != nil {
		// Handle the error
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// If everything goes well, return a success message
	return c.JSON(http.StatusOK, map[string]string{"message": "Bet successfully updated"})
}

// InsertBet godoc
// @Summary Insert a new bet
// @Description Insert a new bet
// @Tags Bets
// @Accept  json
// @Produce  json
// @Param userid query int true "User ID"
// @Param matchid query int true "Match ID"
// @Param amount query float64 true "Amount"
// @Param prediction query string true "Prediction"
// @Success 200 {object} structs.Bet
// @Router /api/v1/bet/insert [post]
func InsertBet(c echo.Context) error {
	// Connect to the database
	db, err := db.ConnectDB()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Make sure the connection will be closed at the end of this function
	defer db.Close()

	// Get the values of the parameters from the request
	userid, _ := strconv.Atoi(c.QueryParam("userid"))
	matchid, _ := strconv.Atoi(c.QueryParam("matchid"))
	amount, _ := strconv.ParseFloat(c.QueryParam("amount"), 64)
	prediction := c.QueryParam("prediction")

	// Execute the SQL query that inserts a new bet with the provided parameters
	_, err = db.Exec("INSERT INTO Bet (UserId, MatchId, BetAmount, SelectedOutcome) VALUES (?, ?, ?, ?)", userid, matchid, amount, prediction)
	if err != nil {
		// Handle the error
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// If everything goes well, return a success message
	return c.JSON(http.StatusOK, map[string]string{"message": "Bet successfully inserted"})
}

// GetByIdBet godoc
// @Summary Get a bet by ID
// @Description Get a bet by ID
// @Tags Bets
// @Accept  json
// @Produce  json
// @Param id path string true "Bet ID"
// @Success 200 {object} structs.Bet
// @Router /api/v1/bet/getbyid/{id} [get]
func GetBetById(c echo.Context) error {
	// Connect to the database
	db, err := db.ConnectDB()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Make sure the connection will be closed at the end of this function
	defer db.Close()

	// Get the value of the id parameter from the URL route
	id := c.Param("id")

	// Execute the SQL query that selects the bet with the provided ID
	row := db.QueryRow("SELECT MatchId, UserId, SelectedOutcome, BetAmount, PossibleReturn FROM Bet WHERE id = ?", id)

	// Create a variable of type Bet to store the bet data
	var bet structs.Bet

	// Read the result of the query and fill the bet structure with the obtained data
	err = row.Scan(&bet.MatchId, &bet.UserId, &bet.SelectedOutcome, &bet.BetAmount, &bet.PossibleReturn)
	if err != nil {
		// Handle the error
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Convert the bet structure into JSON
	betJSON, err := json.Marshal(bet)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Send the response in JSON
	return c.JSONBlob(http.StatusOK, betJSON)
}

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
	rows, err := db.Query("SELECT Name, City, Country, ID FROM Team ORDER BY Name")
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
		err = rows.Scan(&team.Name, &team.City, &team.Country, &team.ID)
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
		err = rows.Scan(&live.ID, &live.HomeTeam, &live.VisitingTeam, &live.Championship, &live.DateMatch, &live.Stadium, &live.StatusMatch, &live.TeamPoints1, &live.TeamPoints2, &live.HomeTeamWins, &live.VisitingTeamWins, &live.Draws, &live.HomeTeamRecentPerformance, &live.VisitingTeamRecentPerformance)
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

func generateSecretKey() string {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatalf("Erro ao gerar a chave secreta: %v", err)
	}
	return base64.StdEncoding.EncodeToString(key)
}

// Login godoc
// @Summary Login a new user
// @Description Login a new user with username and password
// @Tags Users
// @Accept  multipart/form-data
// @Produce  json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {object} structs.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/user/login [post]
func Login(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		log.Printf("Erro ao conectar ao banco de dados: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro interno do servidor"})
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Pega o nome de usuário e a senha do corpo da solicitação
	u := new(structs.User)
	if err := c.Bind(u); err != nil {
		log.Printf("Erro ao fazer bind dos dados do usuário: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Dados inválidos fornecidos"})
	}

	// Verifica se o nome de usuário e a senha foram fornecidos
	if u.Username == "" || u.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Nome de usuário e senha devem ser fornecidos"})
	}

	// Executa a consulta SQL
	row := db.QueryRow("SELECT ID, Username, Password, Photo, Cash FROM User WHERE Username = ?", u.Username)

	// Lê o resultado
	userInDb := structs.User{}

	err = row.Scan(&userInDb.ID, &userInDb.Username, &userInDb.Password, &userInDb.Photo, &userInDb.Cash)
	if err == sql.ErrNoRows {
		// Nenhum usuário com o nome de usuário fornecido foi encontrado
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Nome de usuário ou senha inválidos"})
	} else if err != nil {
		// Um erro diferente ocorreu
		log.Printf("Erro ao executar a consulta SQL: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro interno do servidor"})
	}

	// Verifica a senha
	err = bcrypt.CompareHashAndPassword([]byte(userInDb.Password), []byte(u.Password))
	if err == nil {
		// A senha está correta, então faça o login do usuário

		// Gera um token JWT
		token := jwt.New(jwt.SigningMethodHS256)

		// Armazena as reivindicações no token
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = userInDb.Username
		claims["admin"] = true // ou qualquer outra reivindicação que você queira

		// Assina o token com uma chave secreta
		secretKey := generateSecretKey()
		tokenString, _ := token.SignedString([]byte(secretKey))

		// Remova a senha do objeto userInDb antes de retorná-lo
		userInDb.Password = ""
		// Retorna os detalhes do usuário e o token
		return c.JSON(http.StatusOK, map[string]interface{}{
			"id":       userInDb.ID,
			"username": userInDb.Username,
			"photo":    base64.StdEncoding.EncodeToString(userInDb.Photo),
			"cash":     userInDb.Cash,
			"token":    tokenString,
		})
	}

	// Se chegamos até aqui, o login falhou
	return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Nome de usuário ou senha inválidos"})
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with username, password and an optional photo
// @Tags Users
// @Accept  multipart/form-data
// @Produce  json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Param photo formData file false "User Photo"
// @Success 200 {object} structs.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
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

	// Pega a foto do corpo da solicitação, se houver
	var photoBytes []byte
	photo, err := c.FormFile("photo")
	if err == nil {
		src, err := photo.Open()
		if err != nil {
			return c.String(http.StatusBadRequest, "Failed to open photo")
		}
		defer src.Close()

		// Lê a foto em um array de bytes
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, src); err != nil {
			return c.String(http.StatusBadRequest, "Failed to read photo")
		}
		photoBytes = buf.Bytes()
	}

	// Executa a consulta SQL para inserir o novo usuário
	_, err = db.Exec("INSERT INTO User (Username, Password, Photo) VALUES (?, ?, ?)", u.Username, hashedPassword, photoBytes)
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
// @Param team body structs.Team true "Team"
// @Success 200 {array} structs.Live
// @Router /api/v1/live/getallbyidteam [get]
func GetAllByIdTeam(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Obtém os dados da equipe do corpo da solicitação
	team := new(structs.Team)
	if err := c.Bind(team); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Prepara a consulta SQL
	stmt, err := db.Prepare("SELECT * FROM Live WHERE HomeTeam = ? OR VisitingTeam = ?")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Executa a consulta SQL
	rows, err := stmt.Query(team, team)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Cria um slice de partidas ao vivo
	lives := []structs.Live{}

	// Lê os resultados
	for rows.Next() {
		live := structs.Live{}
		err = rows.Scan(&live.ID, &live.HomeTeam, &live.VisitingTeam, &live.Championship, &live.DateMatch, &live.Stadium, &live.StatusMatch, &live.TeamPoints1, &live.TeamPoints2, &live.HomeTeamWins, &live.VisitingTeamWins, &live.Draws, &live.HomeTeamRecentPerformance, &live.VisitingTeamRecentPerformance)
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

	// Executa a consulta SQL
	rows, err := db.Query("SELECT * FROM Live WHERE DATE(DateMatch) = CURDATE() ORDER BY DateMatch")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Cria um slice de partidas ao vivo
	lives := []structs.Live{}

	// Lê os resultados
	for rows.Next() {
		// Cria uma partida ao vivo vazia
		live := structs.Live{}

		// Cria variáveis para armazenar os IDs das equipes, do campeonato e a data da partida
		var idTeam1, idTeam2, idChampionship int
		var dateMatch string

		// Preenche as variáveis com os dados da linha
		err = rows.Scan(&live.ID, &idTeam1, &idTeam2, &idChampionship, &dateMatch, &live.Stadium, &live.StatusMatch, &live.TeamPoints1, &live.TeamPoints2, &live.HomeTeamWins, &live.VisitingTeamWins, &live.Draws, &live.HomeTeamRecentPerformance, &live.VisitingTeamRecentPerformance)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// Converte a string de data/hora para um time.Time
		live.DateMatch, err = time.Parse("2006-01-02 15:04:05", dateMatch)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// Busca as equipes e o campeonato correspondentes no banco de dados
		live.HomeTeam, err = FetchTeamByID(strconv.Itoa(idTeam1))
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		live.VisitingTeam, err = FetchTeamByID(strconv.Itoa(idTeam2))
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		live.Championship, err = FetchChampionshipByID(strconv.Itoa(idChampionship))
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

// FetchChampionshipByID godoc
// @Summary Fetch a championship by ID
// @Description Fetch a championship by ID
// @Tags Championships
// @Accept  json
// @Produce  json
// @Param id path string true "Championship ID"
// @Success 200 {object} structs.Championship
// @Router /api/v1/championship/fetchbyid/{id} [get]
func FetchChampionshipByID(id string) (structs.Championship, error) {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		return structs.Championship{}, err
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Executa a consulta SQL que seleciona o campeonato com o ID informado
	row := db.QueryRow("SELECT Id, Name, COALESCE(Year, 0) FROM Championship WHERE Id = ?", id)

	// Cria uma variável do tipo Championship para armazenar os dados do campeonato
	var championship structs.Championship

	// Lê o resultado da consulta e preenche a estrutura do campeonato com os dados obtidos
	err = row.Scan(&championship.Id, &championship.Name, &championship.Year)
	if err != nil {
		// Lida com o erro
		return championship, err
	}

	return championship, nil
}

// FetchTeamByID godoc
// @Summary Fetch a team by ID
// @Description Fetch a team by ID
// @Tags Teams
// @Accept  json
// @Produce  json
// @Param id path string true "Team ID"
// @Success 200 {object} structs.Team
// @Router /api/v1/team/fetchbyid/{id} [get]
func FetchTeamByID(id string) (structs.Team, error) {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		return structs.Team{}, err
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Executa a consulta SQL que seleciona o time com o ID informado
	row := db.QueryRow("SELECT Name, City, Country, ID FROM Team WHERE id = ?", id)

	// Cria uma variável do tipo Team para armazenar os dados do time
	var team structs.Team

	// Lê o resultado da consulta e preenche a estrutura do time com os dados obtidos
	err = row.Scan(&team.Name, &team.City, &team.Country, &team.ID)
	if err != nil {
		// Lida com o erro
		return team, err
	}

	return team, nil
}

// InsertLive godoc
// @Summary Insert a live match
// @Description Insert a live match
// @Tags Lives
// @Accept  json
// @Produce  json
// @Param live body structs.Live true "Live Match"
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

	// Obtém os dados da partida ao vivo do corpo da solicitação
	live := new(structs.Live)
	if err := c.Bind(live); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Executa a consulta SQL
	_, err = db.Exec("INSERT INTO Live (HomeTeam, VisitingTeam, Championship, DateMatch, Stadium, TeamPoints1, TeamPoints2) VALUES (?, ?, ?, ?, ?, ?, ?)", live.HomeTeam, live.VisitingTeam, live.Championship, live.DateMatch, live.Stadium, live.TeamPoints1, live.TeamPoints2)
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
// @Param live body structs.Live true "Live Match"
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

	// Obtém os dados da partida ao vivo do corpo da solicitação
	live := new(structs.Live)
	if err := c.Bind(live); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Verifica se a partida ao vivo existe
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM Live WHERE id = ?", live.ID).Scan(&count)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if count == 0 {
		return c.String(http.StatusNotFound, "Partida ao vivo não encontrada")
	}

	// Executa a consulta SQL
	_, err = db.Exec("UPDATE Live SET HomeTeam = ?, VisitingTeam = ?, Championship = ?, DateMatch = ?, Stadium = ?, TeamPoints1 = ?, TeamPoints2 = ? WHERE id = ?", live.HomeTeam, live.VisitingTeam, live.Championship, live.DateMatch, live.Stadium, live.TeamPoints1, live.TeamPoints2, live.ID)
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

	// Executa a consulta SQL
	result, err := db.Exec("INSERT INTO Team (name, city, country", name, city, country)
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

	// Verifica se os campos obrigatórios estão presentes
	if name == "" || city == "" || country == "" {
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
	row := db.QueryRow("SELECT Name, City, Country, ID FROM Team WHERE id = ?", id)

	// Cria uma variável do tipo Team para armazenar os dados do time
	var team structs.Team

	// Lê o resultado da consulta e preenche a estrutura do time com os dados obtidos
	err = row.Scan(&team.Name, &team.City, &team.Country, &team.ID)
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
// @Param live body structs.Live true "Live Match"
// @Success 200 {object} structs.Live
// @Router /api/v1/live/getbyid [get]
func GetByIdLive(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Obtém os dados da partida ao vivo do corpo da solicitação
	live := new(structs.Live)
	if err := c.Bind(live); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Executa a consulta SQL que seleciona a partida ao vivo com o ID informado
	row := db.QueryRow("SELECT * FROM Live WHERE id = ?", live.ID)

	// Lê o resultado da consulta e preenche a estrutura da partida ao vivo com os dados obtidos
	err = row.Scan(&live.ID, &live.HomeTeam, &live.VisitingTeam, &live.Championship, &live.DateMatch, &live.Stadium, &live.StatusMatch, &live.TeamPoints1, &live.TeamPoints2, &live.HomeTeamWins, &live.VisitingTeamWins, &live.Draws, &live.HomeTeamRecentPerformance, &live.VisitingTeamRecentPerformance)
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
	row := db.QueryRow("SELECT Name, City, Country, Birth, IdTeam, ID, COALESCE(Height, '0.00') as Height, COALESCE(Position, '') as Position, COALESCE(ImagePath, '') as ImagePath, COALESCE(ShirtNumber, '0') as ShirtNumber FROM Player WHERE ID = ?", id)

	// Create a variable of type Player to store the player data
	var player structs.Player

	// Read the result of the query and fill the player structure with the obtained data
	var birthStr string
	err = row.Scan(&player.Name, &player.City, &player.Country, &birthStr, &player.IdTeam, &player.ID, &player.Height, &player.Position, &player.ImagePath, &player.ShirtNumber)
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
	// Connect to the database
	db, err := db.ConnectDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	idteam := c.Param("idteam")

	rows, err := db.Query("SELECT Name, City, Country, COALESCE(Birth, '') as Birth, IdTeam, ID, COALESCE(Height, '0.00') as Height, COALESCE(Position, '') as Position FROM Player WHERE IdTeam = ?", idteam)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	players := []structs.Player{}

	for rows.Next() {
		var player structs.Player
		var birthStr string
		err = rows.Scan(&player.Name, &player.City, &player.Country, &birthStr, &player.IdTeam, &player.ID, &player.Height, &player.Position)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		if birthStr != "0000-00-00" && birthStr != "" {
			player.Birth, err = time.Parse("2006-01-02", birthStr)
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
		}

		players = append(players, player)
	}

	rows.Close()

	playersJSON, err := json.Marshal(players)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

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
	rows, err := db.Query("SELECT Name, City, Country, Id FROM Team WHERE name LIKE ?", "%"+name+"%")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	// Cria uma lista para armazenar os times
	var teams []structs.Team

	// Lê o resultado da consulta e preenche a lista de times com os dados obtidos
	for rows.Next() {
		var team structs.Team
		err = rows.Scan(&team.Name, &team.City, &team.Country, &team.ID)
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
	rows, err := db.Query("SELECT Name, City, Country, Birth, IdTeam, ID, Height, COALESCE(Position, '') as Position, COALESCE(ImagePath, '') as ImagePath, COALESCE(ShirtNumber, '0') as ShirtNumber FROM Player WHERE name LIKE ?", "%"+name+"%")
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
		err = rows.Scan(&player.Name, &player.City, &player.Country, &birthStr, &player.IdTeam, &player.ID, &player.Height, &player.Position, &player.ImagePath, &player.ShirtNumber)
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
// @Param idteam query int false "Id Team"
// @Param city query string false "City"
// @Param country query string false "Country"
// @Param birth query string false "Birth" example="DD-MM-AAAA"
// @Param height query string false "Height"
// @Param position query string false "Position"
// @Param position query string false "ImagePath"
// @Param position query string false "ShirtNumber"
// @Success 200 {object} structs.Player
// @Router /api/v1/player/update [put]
func UpdatePlayer(c echo.Context) error {
	// Conecta ao banco de dados
	db, err := db.ConnectDB()
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
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
	imagepath := c.QueryParam("imagepath")
	shirtnumber := c.QueryParam("shirtnumber")

	// Converte a string de data de nascimento para o tipo date
	var birthDate time.Time
	if birth != "" {
		birthDate, err = time.Parse("02/01/2006", birth)
		if err != nil {
			fmt.Println("Erro ao converter a data de nascimento:", err)
			return c.String(http.StatusInternalServerError, err.Error())
		}
	}

	// Formata a data de nascimento para o formato aceito pelo MySQL
	formattedBirthDate := birthDate.Format("2006-01-02")

	// Executa a consulta SQL
	result, err := db.Exec("UPDATE Player SET name = ?, city = ?, country = ?, birth = ?, idteam = ?, height = ?, position = ?, imagepath = ?, shirtnumber = ? WHERE id = ?", name, city, country, formattedBirthDate, idteam, height, position, imagepath, shirtnumber, id)
	if err != nil {
		fmt.Println("Erro ao executar a consulta SQL:", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Certifique-se de que a conexão será fechada no final desta função
	defer db.Close()

	// Obtém o número de linhas afetadas
	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Erro ao obter o número de linhas afetadas:", err)
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
		err = db.QueryRow("SELECT Name, City, Country, ID FROM Team WHERE id = ?", idTeam).Scan(&team.Name, &team.City, &team.Country, &team.ID)
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
