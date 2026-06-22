package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"
	"github.com/Davyjakim/gator/internal/database"
	"github.com/google/uuid"
)

func handleRegister(s *state, cmd command) error{
	if len(cmd.args)==0{
		return errors.New("The username was not provided")
	}
	user, err:=s.db.CreateUser(context.Background(),database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: cmd.args[0],
	})
	if err!=nil{
		
		os.Exit(1)
		return err
	}
	err=s.cfg.SetUser(user.Name, user.ID)
	if err!=nil{
		
		return err
	}
	fmt.Printf("The user: %s was created\n", user.Name)
	return  nil
}