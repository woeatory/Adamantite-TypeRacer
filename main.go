package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/woeatory/Adamantite-TypeRacer/game"
)

func main() {
	for true {
		prompt := promptui.Select{
			Label: "Welcome to Adamantite-TypeRacer - a CLI-based typing game",
			Items: []string{
				"Log In",
				"Sign Up",
				"Test your skill solo",
				"Create room",
				"Join public room",
				"Exit",
			},
		}

		num, _, err := prompt.Run()

		if err != nil {

			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		switch num {
		case 0: // Log In
			// todo implement Log In
		case 1: // Sign Up
			// todo implement Sign Up
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
