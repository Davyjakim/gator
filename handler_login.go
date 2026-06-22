package main

import (
	"context"
	"errors"
	"fmt"
	"os"
)

func handleLogin(s *state, cmd command) error{
	if len(cmd.args)==0{
		return errors.New("The username was not provided")
	}
	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err!=nil{
		os.Exit(1)
	}
	err=s.cfg.SetUser(user.Name)
	if err!=nil{
		return err
	}
	fmt.Printf("User: %s logged in\n", cmd.args[0])
	return nil
}
