package main

import (
	"context"
	"fmt"
	"os"
	"time"
	"github.com/Davyjakim/gator/internal/database"
	"github.com/google/uuid"
)

func handleAddFeed(s *state, cmd command, user database.User)error{
	if len(cmd.args)<2{
		fmt.Println("Not enough argument provided. example: addfeed <name> <url>")
		os.Exit(1)
	}

	feed, err:=s.db.AddFeed(context.Background(),database.AddFeedParams{
		ID: uuid.New(),
		CreatedAt:time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		Name: cmd.args[0],
		Url: cmd.args[1],
	})
	if err!=nil{
		return err
	}
	fmt.Printf("The feed: %s was created\n", feed.Name)
	feedFollow,err:=s.db.CreateFeedFollow(context.Background(),database.CreateFeedFollowParams{
		ID: uuid.New(),
		UserID: user.ID,
		FeedID: feed.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err!=nil{
	
		return err
	}
	fmt.Printf("The user: %s now follows the Feed: %s\n", feedFollow.UserName,feedFollow.FeedName)	
	
	
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