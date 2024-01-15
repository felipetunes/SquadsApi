package router

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func Initialize() {
	// Cria uma nova instância do Echo
	e := echo.New()

	// Define as rotas da aplicação
	e.GET("team/getall", GetAllTeams)
	e.GET("team/insert", InsertTeam)
	e.GET("team/update", UpdateTeam)
	e.GET("team/delete", DeleteTeam)
	e.GET("team/getbyid/:id", GetByIdTeam)
	e.GET("team/getbyname/:name", GetByNameTeam)
	e.GET("team/getbycountry/:country", GetByCountryTeam)
	e.GET("player/getall", GetAllPlayers)
	e.GET("player/getbyidteam/:idteam", GetByIdTeamPlayer)

	// Inicia o servidor na porta 8080
	e.Logger.Fatal(e.Start(":8080"))
}
