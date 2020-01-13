package main

import (
	"fmt"
	"github.com/alexaxms/TopSecret/config"
	"github.com/alexaxms/TopSecret/postgres"
	"github.com/alexaxms/TopSecret/registry"
	"github.com/alexaxms/TopSecret/router"
	"github.com/labstack/echo"
	"log"
)

func main() {
	config.ReadConfig()

	db := postgres.NewDB()
	db.LogMode(true)
	defer db.Close()

	r := registry.NewRegistry(db)
	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
