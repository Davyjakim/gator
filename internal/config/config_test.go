package config

import (
	"fmt"

	"testing"
)


func TestSetUser(t *testing.T){
	cases:= []struct{
		input string
		expected string
	}{
		{
			input:"jakim",
			expected:"jakim",
		},
		{
			input:"kim",
			expected:"kim",
		},
		{
			input:"leon",
			expected:"leon",
		},
	}
	for _, c :=range cases{
		cfg, err:= Read()
		if err!=nil{
			t.Errorf("Test failed due to: %s\n", err)
		}
		err= cfg.SetUser(c.input)
		if err!=nil{
			t.Errorf("Test fail due to:%s\n",err)
		}
		actual, err:= Read()
		if err!=nil{
			t.Errorf("Test failed due to: %s\n", err)
		}
		if actual.CurrentUserName!=c.expected{
			t.Errorf("\nFail")
			fmt.Printf("\n------------\nFAILED\n")
			fmt.Printf("Expected_UserName: %s\n", c.expected)
			fmt.Printf("Actual_UserName: %s\n", actual.CurrentUserName)
		}else{
			fmt.Printf("\n------------\nPASSED\n")
			fmt.Printf("Expected_UserName: %s\n", c.expected)
			fmt.Printf("Actual_UserName: %s\n", actual.CurrentUserName)
		}
	}
	

}