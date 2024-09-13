package main

import (
	"log"
	"s3service/pkg/config"
	"s3service/pkg/di"
)

func main() {
	cnf, err := config.LoadConfig()

	if err != nil {
		log.Fatal("failed to load environements")
	}
	server := di.Init(cnf)
	server.Start()
}
