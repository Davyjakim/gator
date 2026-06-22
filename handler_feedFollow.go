package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Davyjakim/gator/internal/database"
	"github.com/google/uuid"
)


func handleFollow(s *state, cmd command)error{
	if len(cmd.args)<1{
		return fmt.Errorf("The arguments are not enough: follow <url>")
	}
	feed, err:=s.db.GetAFeed(context.Background(),cmd.args[0])
	if err!=nil{

		return  err
	}
	feedFollow,err:=s.db.CreateFeedFollow(context.Background(),database.CreateFeedFollowParams{
		ID: uuid.New(),
		UserID: s.cfg.CurrentUserId,
		FeedID: feed.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err!=nil{
	
		return err
	}
	fmt.Printf("Feed name: %s, current userName: %s\nnnn", feedFollow.FeedName,s.cfg.CurrentUserName)	
	return nil
}