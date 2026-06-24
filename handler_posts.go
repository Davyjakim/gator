package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Davyjakim/gator/internal/database"
)

func handleBrowse(s *state, cmd command, user database.User)error{
	limit :=2
	if len(cmd.args)>0{
		num, err := strconv.Atoi(cmd.args[0])
		if err!=nil{
			return err
		}
		limit = num
	}
	posts, err:=s.db.GetPostsForUser(context.Background(),database.GetPostsForUserParams{
		ID: user.ID,
		Limit: int32(limit),
	})
	if err!=nil{
		return  err
	}
	for i,p :=range posts{
		fmt.Printf("\n\n----------Posts(%d)-----------\n",i+1)
		fmt.Printf("\u2010 Title: %s\n",p.Title)
		fmt.Printf("\u2010 Description:\n %s\n",p.Description.String)
		fmt.Printf("\u2010 Published date: %s\n",p.PublishedAt.Time)
	}
	return nil
}