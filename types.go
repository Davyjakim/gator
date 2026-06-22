package main

import (
	"fmt"

	"github.com/Davyjakim/gator/internal/config"
	"github.com/Davyjakim/gator/internal/database"
)

type CliCommand struct{
	name string
	description string
	callback func(state *state, arg ...string) error
}

type state struct{
	cfg *config.Config
	db *database.Queries
}
type command struct{
	name string
	args []string
}

type commands struct{
	cmdMap map[string]func(state *state,command command) error
}
type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func (c *commands)run(s *state, cmd command) error{
	if c, ok:= c.cmdMap[cmd.name]; ok{
		c(s,cmd)
		return nil
	}else{
		return fmt.Errorf("This command '%s' is unkown", cmd.name)
	}
}
func (c *commands) register(name string,f func(*state,command)error){
	c.cmdMap[name]= f
}

