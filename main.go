package main

import (
	"fmt"
	"log"
	"os"

	"github.com/YaguarEgor/gator_cli/internal/config"
)

const username = "Egor"

type state struct {
	conf *config.Config
}

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	programState := state{
		conf: &conf,
	}
	cmds := commands{
		registeredCommands: map[string]func(*state, command) error{},
	}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("you have to type a command, when start Gator")
		os.Exit(1)
	}
	name := os.Args[1]
	var args []string
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}
	cmd := command{
		name: name,
		args: args,
	}
	err = cmds.run(&programState, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}