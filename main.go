package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Davyjakim/gator/internal/config"
	"github.com/Davyjakim/gator/internal/database"
	_ "github.com/lib/pq"
)

func main(){
	
	cfg, err:= config.Read()
	if err!=nil{
		fmt.Println(err)
		return
	}
	db, err := sql.Open("postgres", cfg.DbUrl)
	dbQueries := database.New(db)
	appState:= state{
		cfg: &cfg,
		db : dbQueries,
	}
	commandRegistry:= commands{
		cmdMap: map[string]func(state *state, command command) error{},
	}
	commandRegistry.register("login",handleLogin)
	commandRegistry.register("register",handleRegister)
	commandRegistry.register("reset",ClearUsersTable)
	commandRegistry.register("users", handleGetAllUsers)
	commandRegistry.register("agg", handleAgg)
	commandRegistry.register("addfeed", handleAddFeed)
	commandRegistry.register("feeds",handleGetFeeds)
	commandRegistry.register("follow",handleFollow)
	input := os.Args
	if len(input)<2{
		fmt.Printf("No enough arguments provided ")
		os.Exit(1)
	}
	useCommand:=command{
		name: input[1],
		args: input[2:],
	}
	cmd, ok:= commandRegistry.cmdMap[useCommand.name]
	if !ok{
		fmt.Printf("This command: '%s' is unknown\n", useCommand.name)
		os.Exit(1)
	}
	err= cmd(&appState,useCommand)
	if err != nil {
		log.Fatal(err)
	}
}