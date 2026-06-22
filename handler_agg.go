package main

import (
	"context"
	"fmt"
)

const feedURL = "https://www.wagslane.dev/index.xml"
func handleAgg(s *state, cmd command) error{
	feed, err:= fetchFeed(context.Background(), feedURL)
	if err!=nil{
		return err
	}
	fmt.Println(*feed)
	return nil
}