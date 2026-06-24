package main

import (
	"context"
	"database/sql"
	"fmt"
	"html"
	"log"
	"strings"
	"time"

	"github.com/Davyjakim/gator/internal/database"
	"github.com/google/uuid"
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
		publishedDate, err := time.Parse(time.RFC1123Z, strings.TrimSpace(rssI.PubDate))
		if err!=nil {
			return err
		}
		post,err:= s.db.CreatePost(context.Background(),database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title: html.UnescapeString(rssI.Title),
			Url: rssI.Link,
			Description: sql.NullString{String: html.UnescapeString(rssI.Description),Valid: true},
			PublishedAt: sql.NullTime{Time: publishedDate, Valid: true},
			FeedID: nextFeed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
		fmt.Printf("\u2022 %s\n",post.Title)
	}
	return nil
}