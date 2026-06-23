package main

import (
	"context"
	"fmt"
	"html"
	"time"
)

func handleAgg(s *state, cmd command) error{
	if len(cmd.args)<1{
		return fmt.Errorf("Not enough arguments provided the time between request is required")
	}
	timeBetweenRequests,err:= time.ParseDuration(cmd.args[0])
	if err!=nil{
		return err
	}
	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()
	fmt.Printf("Collecting feeds every %s\n",timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
		fmt.Printf("\n")
	}
	
}

func scrapeFeeds(s *state) error{
	
	nextFeed,err:= s.db.GetNextFeedToFetch(context.Background())
	if err!=nil{
		return err
	}
	fmt.Println(nextFeed.Name)
	err=s.db.MarkFeedFetched(context.Background(),nextFeed.ID)
	if err!=nil{
		return err
	}
	rssfeed,err := fetchFeed(context.Background(),nextFeed.Url)
	if err!=nil{
		return err
	}
	for _,rssI:=range rssfeed.Channel.Item{
		fmt.Printf("(current: %s) \u2022 %s \n ",nextFeed.Name,html.UnescapeString(rssI.Title))
	}
	return nil
}