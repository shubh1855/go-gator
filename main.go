package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/shubh1855/go-gator/internal/cli"
	"github.com/shubh1855/go-gator/internal/config"
	"github.com/shubh1855/go-gator/internal/database"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("not enough arguments provided")
	}

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	state := cli.NewState(&cfg, dbQueries)

	cmds := cli.NewCommands()
	cmds.Register("login", cli.HandlerLogin)
	cmds.Register("register", cli.HandlerRegister)
	cmds.Register("reset", cli.HandlerReset)
	cmds.Register("users", cli.HandleUsers)
	cmds.Register("agg", cli.HandlerAgg)
	cmds.Register("addfeed", cli.MiddlewareLoggedIn(cli.HandlerAddFeed))
	cmds.Register("feeds", cli.HandlerFeeds)
	cmds.Register("follow", cli.MiddlewareLoggedIn(cli.HandlerFollow))
	cmds.Register("following", cli.MiddlewareLoggedIn(cli.HandlerFollowing))
	cmds.Register("unfollow", cli.MiddlewareLoggedIn(cli.HandlerUnfollow))
	cmds.Register("browse", cli.MiddlewareLoggedIn(cli.HandlerBrowse))

	cmd := cli.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	if err := cmds.Run(state, cmd); err != nil {
		log.Fatal(err)
	}
}
