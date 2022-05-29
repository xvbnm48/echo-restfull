package main

import (
	"github.com/xvbnm48/echo-restfull/db"
	"github.com/xvbnm48/echo-restfull/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
