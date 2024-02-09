package router

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"net/http"

	"github.com/rs/cors"
)

func Initialize() {
	// Cria uma nova instância do Echo
	e := echo.New()

	// Configura o CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"https://localhost:7299"},
	})

	handler := c.Handler(e)

	// Cria um grupo de rotas com o prefixo api/v1
	apiV1 := e.Group("api/v1")

	// Define as rotas da aplicação
	defineTeamRoutes(apiV1)
	definePlayerRoutes(apiV1)
	defineLiveRoutes(apiV1)
	defineUserRoutes(apiV1)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Inicia o servidor na porta 8080
	e.Logger.Fatal(http.ListenAndServe(":8080", handler))
}

func defineUserRoutes(g *echo.Group) {

	// @Summary Login user
	// @Description Login user
	// @Tags Users
	// @Accept  json
	// @Produce  json
	// @Param username query string true "Username"
	// @Param password query string true "Password"
	// @Success 200 {object} structs.User
	// @Router /api/v1/user/login [post]
	g.POST("/user/login", Login)

	// @Summary Register user
	// @Description Register user
	// @Tags Users
	// @Accept  json
	// @Produce  json
	// @Param username query string true "Username"
	// @Param password query string true "Password"
	// @Success 200 {object} structs.User
	// @Router /api/v1/user/register [post]
	g.POST("/user/register", Register)
}

func defineLiveRoutes(g *echo.Group) {

	// @Summary Obter todos os jogos ao vivo
	// @Description Obter todos os jogos
	// @Tags Lives
	// @Accept  json
	// @Produce  json
	// @Success 200 {array} structs.Live
	// @Router /api/v1/live/getall [get]
	g.GET("/live/getall", GetAllMatches)

	// @Summary Obter todos os jogos do dia
	// @Description Obter todos os jogos do dia
	// @Tags Lives
	// @Accept  json
	// @Produce  json
	// @Success 200 {array} structs.Live
	// @Router /api/v1/live/getalltoday [get]
	g.GET("/live/getalltoday", GetAllLivesToday)

	// @Summary Obter todos os jogos de um time especifico
	// @Description Obter todos os jogos de um time especifico
	// @Tags Lives
	// @Accept  json
	// @Produce  json
	// @Success 200 {array} structs.Live
	// @Router /api/v1/live/getallbyidteam/:id  [get]
	g.GET("/live/getallbyidteam/:id", GetAllByIdTeam)

	// @Summary Inserir uma partida ao vivo
	// @Description Inserir uma partida ao vivo
	// @Tags Lives
	// @Accept  json
	// @Produce  json
	// @Param teamid1 query int true "Team ID 1"
	// @Param teamid2 query int true "Team ID 2"
	// @Param idchampionship query int true "IdChampionship"
	// @Param datematch query string true "Date of Match"
	// @Param stadium query string true "Stadium"
	// @Param teampoints1 query int true "Team Points 1"
	// @Param teampoints2 query int true "Team Points 2"
	// @Success 200 {object} structs.Live
	// @Router /api/v1/live/insert [post]
	g.POST("/live/insert", InsertLive)

	// @Summary Atualizar uma partida ao vivo
	// @Description Atualizar uma partida ao vivo
	// @Tags Lives
	// @Accept  json
	// @Produce  json
	// @Param id query int true "Live ID"
	// @Param teamid1 query int true "Team ID 1"
	// @Param teamid2 query int true "Team ID 2"
	// @Param idchampionship query int true "IdChampionship"
	// @Param datematch query string true "Date of Match"
	// @Param stadium query string true "Stadium"
	// @Param teampoints1 query int true "Team Points 1"
	// @Param teampoints2 query int true "Team Points 2"
	// @Success 200 {object} structs.Live
	// @Router /api/v1/live/update [put]
	g.PUT("/live/update", UpdateLive)

	// @Summary Obter uma partida ao vivo por ID
	// @Description Obter uma partida ao vivo por ID
	// @Tags Lives
	// @Accept  json
	// @Produce  json
	// @Success 200 {object} structs.Live
	// @Router /api/v1/live/getbyid/{id} [get]
	g.GET("/live/getbyid/:id", GetByIdLive)

}

func defineTeamRoutes(g *echo.Group) {
	// @Summary Obter todos os times
	// @Description Obter todos os times
	// @Tags Teams
	// @Accept  json
	// @Produce  json
	// @Success 200 {array} structs.Team
	// @Router /api/v1/team/getall [get]
	g.GET("/team/getall", GetAllTeams)

	// @Summary Inserir um time
	// @Description Inserir um time
	// @Tags Teams
	// @Accept  json
	// @Produce  json
	// @Param name query string true "Team Name"
	// @Param city query string true "City"
	// @Param country query string true "Country"
	// @Param color1 query string true "Color1"
	// @Success 200 {object} structs.Team
	// @Router /api/v1/team/insert [post]
	g.POST("/team/insert", InsertTeam)

	// @Summary Atualizar um time
	// @Description Atualizar um time
	// @Tags Teams
	// @Accept  json
	// @Produce  json
	// @Param id team query int true "ID Team"
	// @Param name query string true "Team Name"
	// @Param city query string true "City"
	// @Param country query string true "Country"
	// @Param color1 query string true "Color1"
	// @Success 200 {object} structs.Team
	// @Router /api/v1/team/update [put]
	g.PUT("/team/update", UpdateTeam)

	// @Summary Delete a team by ID
	// @Description Delete a team by ID
	// @Tags Teams
	// @Accept  json
	// @Produce  json
	// @Param id path string true "ID Team"
	// @Success 200 {object} string
	// @Router /api/v1/team/delete/{id} [delete]
	g.DELETE("/team/delete/:id", DeletePlayer)

	// @Summary Obter um time por ID
	// @Description Obter um time por ID
	// @Tags Teams
	// @Accept  json
	// @Produce  json
	// @Success 200 {object} structs.Team
	// @Router /api/v1/team/getbyid/{id} [get]
	g.GET("/team/getbyid/:id", GetByIdTeam)

	// @Summary Obter um time por nome
	// @Description Obter um time por nome
	// @Tags Teams
	// @Accept  json
	// @Produce  json
	// @Success 200 {object} structs.Team
	// @Router api/v1/team/getbyname/{name} [get]
	g.GET("/team/getbyname/:name", GetTeamsByName)

	// @Summary Obter um time por país
	// @Description Obter um time por país
	// @Tags Teams
	// @Accept  json
	// @Produce  json
	// @Success 200 {object} structs.Team
	// @Router /api/v1/team/getbycountry/{country} [get]
	g.GET("/team/getbycountry/:country", GetByCountryTeam)

	// @Summary Get teams by championship
	// @Description Get teams by a given championship ID
	// @Tags Teams
	// @Accept  json
	// @Produce  json
	// @Param id path int true "ID Championship"
	// @Success 200 {array} structs.Team
	// @Router /api/v1/team/getbychampionship/{id} [get]
	g.GET("/team/getbychampionship/:id", GetByChampionship)

}

// @title API de Times e Jogadores
// @description Esta é uma API simples para gerenciar times e jogadores.
// @schemes http
// @host localhost:8080
// @BasePath /api/v1
func definePlayerRoutes(g *echo.Group) {
	// @Summary Obter todos os jogadores
	// @Description Obter todos os jogadores
	// @Tags Players
	// @Accept  json
	// @Produce  json
	// @Success 200 {array} structs.Player
	// @Router /api/v1/player/getall [get]
	g.GET("/player/getall", GetAllPlayers)

	// @Summary Delete a player by ID
	// @Description Delete a player by ID
	// @Tags Players
	// @Accept  json
	// @Produce  json
	// @Param id path string true "Player ID"
	// @Success 200 {object} string
	// @Router /api/v1/player/delete/{id} [delete]
	g.DELETE("/player/delete/:id", DeletePlayer)

	// @Summary Obter um jogador por nome
	// @Description Obter um player por nome
	// @Tags Players
	// @Accept  json
	// @Produce  json
	// @Success 200 {object} structs.Player
	// @Router api/v1/player/getbyname/{name} [get]
	g.GET("/player/getbyname/:name", GetPlayersByName)

	// @Summary Obter jogadores por ID do time
	// @Description Obter jogadores por ID do time
	// @Tags Players
	// @Accept  json
	// @Produce  json
	// @Success 200 {array} structs.Player
	// @Router /api/v1/player/getbyidteam/{idteam} [get]
	g.GET("/player/getbyidteam/:idteam", GetByIdTeamPlayer)

	// @Summary Inserir um jogador
	// @Description Inserir um jogador
	// @Tags Players
	// @Accept  json
	// @Produce  json
	// @Param name query string true "Player Name"
	// @Param idteam query string true "Id Team"
	// @Param city query string true "City"
	// @Param country query string true "Country"
	// @Param birth query string true "Birth" example="AAAA-MM-DD"
	// @Param height query float true "Height"
	// @Param position query string true "Position"
	// @Success 200 {object} structs.Player
	// @Router /api/v1/player/insert [post]
	g.POST("/player/insert", InsertPlayer)

	// @Summary Obter um jogador por ID
	// @Description Obter um jogador por ID
	// @Tags Players
	// @Accept  json
	// @Produce  json
	// @Success 200 {object} structs.Player
	// @Router /api/v1/player/getbyid/{id} [get]
	g.GET("/player/getbyid/:id", GetByIdPlayer)

	// @Summary Obter um jogador por país
	// @Description Obter um jogador por país
	// @Tags Players
	// @Accept  json
	// @Produce  json
	// @Success 200 {object} structs.Player
	// @Router /api/v1/player/getbycountry/{country} [get]
	g.GET("/player/getbycountry/:country", GetByCountryPlayer)

	// @Summary Atualizar um jogador
	// @Description Atualizar um jogador
	// @Tags Players
	// @Accept  json
	// @Produce  json
	// @Param id player query int true "ID Team"
	// @Param name query string true "Team Name"
	// @Param city query string true "City"
	// @Param country query string true "Country"
	// @Param height query float true "Height"
	// @Success 200 {object} structs.Player
	// @Router /api/v1/player/update [put]
	g.PUT("/player/update", UpdatePlayer)

}
