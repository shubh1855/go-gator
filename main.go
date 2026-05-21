package main

import (
	"fmt"
	"log"

	"github.com/shubh1855/Gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}
	fmt.Printf("Read config: %v\n", cfg)

	err = cfg.SetUser("shubh")
	if err != nil {
		log.Fatalf("could not set current user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	fmt.Printf("%+v\n", cfg)
}
