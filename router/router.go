package router

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Initialize() {
	// Cria uma nova instância do Echo
	e := echo.New()

	// Cria um grupo de rotas com o prefixo api/v1
	apiV1 := e.Group("api/v1")

	// Define as rotas da aplicação
	defineTeamRoutes(apiV1)
	definePlayerRoutes(apiV1)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// Inicia o servidor na porta 8080
	e.Logger.Fatal(e.Start(":8080"))
}

func defineTeamRoutes(g *echo.Group) {
	// @Summary Obter todos os times
	// @Description Obter todos os times
	// @Tags times
	// @Accept  json
	// @Produce  json
	// @Success 200 {array} structs.Team
	// @Router /api/v1/team/getall [get]
	g.GET("/team/getall", GetAllTeams)

	// @Summary Inserir um time
	// @Description Inserir um time
	// @Tags times
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
	// @Tags times
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

	// @Summary Delete a player by ID
	// @Description Delete a player by ID
	// @Tags Players
	// @Accept  json
	// @Produce  json
	// @Param id path string true "Player ID"
	// @Success 200 {object} string
	// @Router /api/v1/player/delete/{id} [delete]
	g.DELETE("/player/delete/:id", DeletePlayer)

	// @Summary Delete a team by ID
	// @Description Delete a team by ID
	// @Tags Team
	// @Accept  json
	// @Produce  json
	// @Param id path string true "ID Team"
	// @Success 200 {object} string
	// @Router /api/v1/team/delete/{id} [delete]
	g.DELETE("/team/delete/:id", DeletePlayer)

	// @Summary Obter um time por ID
	// @Description Obter um time por ID
	// @Tags times
	// @Accept  json
	// @Produce  json
	// @Success 200 {object} structs.Team
	// @Router /api/v1/team/getbyid/{id} [get]
	g.GET("/team/getbyid/:id", GetByIdTeam)

	// @Summary Obter um time por nome
	// @Description Obter um time por nome
	// @Tags times
	// @Accept  json
	// @Produce  json
	// @Success 200 {object} structs.Team
	// @Router api/v1/team/getbyname/{name} [get]
	g.GET("/team/getbyname/:name", GetTeamsByName)

	// @Summary Obter um time por país
	// @Description Obter um time por país
	// @Tags times
	// @Accept  json
	// @Produce  json
	// @Success 200 {object} structs.Team
	// @Router /api/v1/team/getbycountry/{country} [get]
	g.GET("/team/getbycountry/:country", GetByCountryTeam)
}

// @title API de Times e Jogadores
// @description Esta é uma API simples para gerenciar times e jogadores.
// @schemes http
// @host localhost:8080
// @BasePath /api/v1
func definePlayerRoutes(g *echo.Group) {
	// @Summary Obter todos os jogadores
	// @Description Obter todos os jogadores
	// @Tags players
	// @Accept  json
	// @Produce  json
	// @Success 200 {array} structs.Player
	// @Router /api/v1/player/getall [get]
	g.GET("/player/getall", GetAllPlayers)

	// @Summary Obter um jogador por nome
	// @Description Obter um player por nome
	// @Tags players
	// @Accept  json
	// @Produce  json
	// @Success 200 {object} structs.Player
	// @Router api/v1/player/getbyname/{name} [get]
	g.GET("/player/getbyname/:name", GetPlayersByName)

	// @Summary Obter jogadores por ID do time
	// @Description Obter jogadores por ID do time
	// @Tags players
	// @Accept  json
	// @Produce  json
	// @Success 200 {array} structs.Player
	// @Router /api/v1/player/getbyidteam/{idteam} [get]
	g.GET("/player/getbyidteam/:idteam", GetByIdTeamPlayer)
	// @Summary Inserir um jogador
	// @Description Inserir um jogador
	// @Tags jogadores
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
	// @Tags jogadores
	// @Accept  json
	// @Produce  json
	// @Success 200 {object} structs.Player
	// @Router /api/v1/player/getbyid/{id} [get]
	g.GET("/player/getbyid/:id", GetByIdPlayer)
	// @Summary Obter um jogador por país
	// @Description Obter um jogador por país
	// @Tags jogadores
	// @Accept  json
	// @Produce  json
	// @Success 200 {object} structs.Player
	// @Router /api/v1/player/getbycountry/{country} [get]
	g.GET("/player/getbycountry/:country", GetByCountryPlayer)
	// @Summary Atualizar um jogador
	// @Description Atualizar um jogador
	// @Tags times
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
