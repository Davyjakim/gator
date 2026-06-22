package main

import (
	"context"
	"fmt"
	"os"

)


func handleGetAllUsers(s *state, cmd command) error{
	users, err := s.db.GetUsers(context.Background())
	if err!=nil{
		return err
	}
	for _, user:= range users{
		if user.Name == s.cfg.CurrentUserName{
			fmt.Printf("* %s (current)\n", user.Name)
		}else{
			fmt.Printf("* %s \n", user.Name)
		}
	}
	return nil
}


func ClearUsersTable (s *state, cmd command) error{
	err:=s.db.ResetUsersTable(context.Background())
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Successfully deleted every rows in the users table")
	return nil
}