package main

import (
	"log"

	"github.com/federicodosantos/hello-world-app/config"
	
)

func Run() error {
	conf, err := config.LoadConfig(".env")
	if err != nil {
		return err
	}

	app := config.NewApp(conf)

	app.SetupRoutes()

	return app.Start()
}

func main() {
	if err := Run(); err != nil {
		log.Fatalf("cannot listen to port due to : %s", err)
	}
}