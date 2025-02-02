package main

import (
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