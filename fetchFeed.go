package main

import (
	"context"
	"encoding/xml"
	"net/http"
)


func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error){
	var rssFeed RSSFeed
	client := &http.Client{}

	req,err:= http.NewRequestWithContext(ctx, "GET",feedURL, nil)
	if err!=nil{
		return &rssFeed, err
	}
	req.Header.Set("User-Agent", "gator")
	resp,err:=client.Do(req)
	if err!=nil{
		return &rssFeed, err
	}
	defer resp.Body.Close()
	decoder:= xml.NewDecoder(resp.Body)
	err= decoder.Decode(&rssFeed)
	if err!=nil{
		return &rssFeed, err
	}
	return &rssFeed, nil
}