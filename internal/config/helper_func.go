package config

import (
	"encoding/json"
	"os"
)

func getConfigFilePath() (string, error){
	homedir,err := os.UserHomeDir()
	if err!=nil{
		return "",err
	}
	return  homedir+"/"+ConfigFileName, nil
}

func write(cfg Config) error{
	path, err := getConfigFilePath()
	if err!=nil{
		return err
	}
	file,err:=os.Create(path)
	if err!=nil{
		return err
	}
	defer file.Close()
	enconder:=json.NewEncoder(file)
	err= enconder.Encode(cfg)
	if err!=nil{
		return err
	}
	return  nil
}