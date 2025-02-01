package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("you need to type a username to login a new user")
	}
	err := s.conf.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User %s has been set\n", cmd.args[0])
	return nil
}