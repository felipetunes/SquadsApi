package router

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func Initialize() {
	// Cria uma nova instância do Echo
	e := echo.New()

	// Cria um grupo de rotas com o prefixo api/v1
	V1 := e.Group("api/v1")

	// Define as rotas da aplicação
	V1.GET("/team/getall", GetAllTeams)
	V1.GET("/team/insert", InsertTeam)
	V1.GET("/team/update", UpdateTeam)
	V1.GET("/team/delete", DeleteTeam)
	V1.GET("/team/getbyid/:id", GetByIdTeam)
	V1.GET("/team/getbyname/:name", GetByNameTeam)
	V1.GET("/team/getbycountry/:country", GetByCountryTeam)
	V1.GET("/player/getall", GetAllPlayers)
	V1.GET("/player/getbyidteam/:idteam", GetByIdTeamPlayer)

	// Inicia o servidor na porta 8080
	e.Logger.Fatal(e.Start(":8080"))
}
