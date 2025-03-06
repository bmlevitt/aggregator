package main

import (
	"fmt"
	"log"

	"github.com/bmlevitt/aggregator/internal/config"
)

func main() {
	fmt.Println("Loading config file...")

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	err = cfg.SetUser("ben levitt")
	if err != nil {
		log.Fatal(err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("db url: %v | current user: %v\n", cfg.DBURL, cfg.CurrentUserName)

}