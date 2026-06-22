package main

import (
	"context"
	"fmt"
	"os"
	"time"
	"github.com/Davyjakim/gator/internal/database"
	"github.com/google/uuid"
)

func handleAddFeed(s *state, cmd command)error{
	if len(cmd.args)<2{
		fmt.Println("Not enough argument provided. example: addfeed <name> <url>")
		os.Exit(1)
	}
	currentUser, err:= s.db.GetUser(context.Background(),s.cfg.CurrentUserName)
	if err!=nil{
		return err
	}
	feed, err:=s.db.AddFeed(context.Background(),database.AddFeedParams{
		ID: uuid.New(),
		CreatedAt:time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: currentUser.ID,
		Name: cmd.args[0],
		Url: cmd.args[1],
	})
	if err!=nil{
		return err
	}
	fmt.Printf("The feed: %s was created\n", feed.Name)
	return nil

}

func handleGetFeeds(s *state, cmd command)error{
	
	feeds,err:= s.db.GetFeeds(context.Background())
	if err!=nil{
		return err
	}
	for _,f :=range feeds{
		fmt.Printf("\u2022 %s\n", f.Name)
		fmt.Printf("\u2022 %s\n", f.Username.String)
	}
	return nil
}