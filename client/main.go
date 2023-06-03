package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/woeatory/Adamantite-TypeRacer/client/game"
	"github.com/woeatory/Adamantite-TypeRacer/client/handlers"
)

func main() {
	selects := promptui.Select{
		Label: "Welcome to Adamantite-TypeRacer - a CLI-based typing game",
		Items: []string{
			"Log In",
			"Sign Up",
			"Test your skill solo",
			"Create room",
			"Join public room",
			"Exit",
		},
		HideSelected: true,
	}
	for {
		num, _, err := selects.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		switch num {
		case 0: // Log In
			err := handlers.LogIn()
			if err != nil {
				fmt.Println(err)
				continue
			}
		case 1: // Sign Up
			err := handlers.SignUp()
			if err != nil {
				fmt.Println(err)
				continue
			}
		case 2: // Play solo
			game.SoloTyper()
		case 3: // Create room
			// todo implement create room
		case 4:
			// todo implement join room
		case 5: // exit
			return
		}

	}

}
