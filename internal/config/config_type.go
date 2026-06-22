package config

import "errors"

type Config struct {
	DbUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(userName string)error{
	if len(userName)==0{
		return errors.New("The username was not provided")
	}
	
	c.CurrentUserName = userName
	err:= write(*c)
	if err!=nil{
		return  err
	}
	return nil
}
