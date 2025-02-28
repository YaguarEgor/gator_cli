package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/YaguarEgor/gator_cli/internal/config"
	"github.com/YaguarEgor/gator_cli/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	conf *config.Config
}

func main() {
	
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	db, err := sql.Open("postgres", conf.DB_url)
	if err != nil {
		log.Fatalf("error connecting to DB: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := state{
		db: dbQueries,
		conf: &conf,
	}

	cmds := commands{
		registeredCommands: map[string]func(*state, command) error{},
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))

	if len(os.Args) < 2 {
		fmt.Println("you have to type a command, when start Gator")
		os.Exit(1)
	}
	cmd_name := os.Args[1]
	var cmd_args []string
	if len(os.Args) > 2 {
		cmd_args = os.Args[2:]
	}
	cmd := command{
		name: cmd_name,
		args: cmd_args,
	}
	err = cmds.run(&programState, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		user, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
		if err != nil {
			return fmt.Errorf("couldn't find user: %w", err)
		}
		return handler(s, c, user)
	}
}