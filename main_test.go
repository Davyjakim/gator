package main

import (
	"fmt"
	"testing"

	"github.com/Davyjakim/gator/internal/config"
)



func TestLogin(t *testing.T){
	cases := []struct {
		input    command
		expected string
	}{
		{
			input:command{
					name: "login",
					args: []string{"Davy"},
					},
			expected: "Davy",
		},
		{
			input: command{
				name: "login",
				args: []string{"kim"},
			},
			expected: "kim",
		},
	}
	cfg, err:= config.Read()
	st:=state{
		cfg: &cfg,
	}
	if err!=nil{
		t.Errorf("\nFAILED\n")
		fmt.Printf("The test failed due to:%s \n",err)
	}
	
	for _,c:= range cases{
		err:=handleLogin(&st,c.input)
		if err!=nil{
			t.Errorf("\nFAILED\n")
			fmt.Printf("The test failed due to:%s \n",err)
		}
		actual, err:= config.Read()
		if err!=nil{
			t.Errorf("\nFAILED\n")
			fmt.Printf("The test failed due to:%s \n",err)
		}
		if actual.CurrentUserName!=c.expected{
			t.Errorf("\nFAILED\n")
			fmt.Printf("Expected_UserName: %s\n", c.expected)
			fmt.Printf("Actual_UserName: %s\n", actual.CurrentUserName)
		}else{
			fmt.Printf("\nPASSED\n")
			fmt.Printf("Expected_UserName: %s\n", c.expected)
			fmt.Printf("Actual_UserName: %s\n", actual.CurrentUserName)
		}
	}
}