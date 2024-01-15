package main

import (
	"apiSquads/router"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router.Initialize()
}
