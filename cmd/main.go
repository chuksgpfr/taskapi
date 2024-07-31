package main

import (
	"flag"
	"os"

	"github.com/chuksgpfr/task-api/gin"
	"github.com/chuksgpfr/task-api/pkg"
	"github.com/chuksgpfr/task-api/postgres"
)

func main() {
	os.Getenv("GO_ENV")
	config, err := pkg.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	ginMode := flag.String("mode", config.GIN_MODE, "to set gin mode")
	flag.Parse()

	config.GIN_MODE = *ginMode

	db := postgres.NewDbClient(config.PostgresDSN)
	server := gin.NewServer(config, db)

	err = server.Run(config.ServerAddress)
	if err != nil {
		panic(err)
	}

}
