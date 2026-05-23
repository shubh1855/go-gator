package main

import (
	"log"
	"os"

	"github.com/shubh1855/Gator/internal/cli"
	"github.com/shubh1855/Gator/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("not enough arguments provided")
	}

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	s := cli.NewState(&cfg)

	cmds := cli.NewCommands()
	cmds.Register("login", cli.HandlerLogin)

	cmd := cli.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	if err := cmds.Run(s, cmd); err != nil {
		log.Fatal(err)
	}
}
