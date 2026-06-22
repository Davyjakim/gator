package config

import (
	"encoding/json"
	"os"
)

func Read() (Config, error){
	var cfg Config
	path,err := getConfigFilePath()
	if err!=nil{
		return Config{}, err
	}
	file, err:=os.Open(path)
	
	if err!=nil{
		return Config{}, err
	}
	defer file.Close()
	decoder :=json.NewDecoder(file)
	err= decoder.Decode(&cfg)
	if err!=nil{
		return Config{}, err
	}
	
	return  cfg, nil
}