package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	clearScreen()
	showTitleScreen()
	os.Exit(0)
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

type errorString struct {
	s string
}

//func (e *errorString) Error() string {
//	return e.s
//}

func takeUserInput(availableChoices []string) (string, error) {
	//consoleReader := bufio.NewReader(os.Stdin)
	//userInput, err := consoleReader.ReadString('\n')
	var userInput string
	fmt.Scanln(&userInput)
	var validInput bool
	for _, item := range availableChoices {
		if userInput == item {
			validInput = true
		}
	}
	if validInput {
		fmt.Println("This is valid input!")
	} else {
		return "", errors.New("This is not valid input")
	}
	return userInput, nil
}

func showTitleScreen() {
	clearScreen()
	fmt.Println("------Final Fantasy Battle Simulator------")
	fmt.Println("")
	fmt.Println("1. Start New Game")
	possibleChoices := []string{"1"}
	userInput, err := takeUserInput(possibleChoices)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if userInput == "1" {
		clearScreen()
		fmt.Println("Let's go!")
		time.Sleep(2 * time.Second)
		clearScreen()
	}
}

func createPlayer() {
	abilities := []string{"Attack", "Drink Potion"}
	hero := player{"Jeremy", 100, 50, 20, 10, 2, abilities}
	fmt.Println(hero)
}

func createEnemy() {
}

func determineFirstPlay() {

}

func showEndScreen() {

}
