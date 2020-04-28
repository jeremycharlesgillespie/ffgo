package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	character "github.com/metalgear386/ffgo/src/character"
)

func main() {
	clearScreen()
	showTitleScreen()
	enemy := createEnemy()
	player := createPlayer()
	battleLoop(player, enemy)
	os.Exit(0)
}

func playerDied(player character.Player) bool {
	var playerDied bool
	if player.Hp <= 0 {
		playerDied = true
	}
	return playerDied
}

func enemyDied(enemy character.Enemy) bool {
	var enemyDied bool
	if enemy.Hp <= 0 {
		enemyDied = true
	}
	return enemyDied
}

func enemyDeathScreen() {
	clearScreen()
	fmt.Println("You have defeated the enemy!\nTry again?")
	os.Exit(0)
}

func playerDeathScreen() {
	clearScreen()
	fmt.Println("The player has died while in battle.\nTry again?")
	os.Exit(0)
}

func battleLoop(player character.Player, enemy character.Enemy) {
	for !enemyDied(enemy) && !playerDied(player) {
		showStatusScreen(player, enemy)
		player, enemy = battleMenu(player, enemy)
		time.Sleep(2 * time.Second)
		if enemyDied(enemy) {
			enemyDeathScreen()
		}
		enemy, player = enemy.AttackPlayer(player)
		time.Sleep(2 * time.Second)
		if playerDied(player) {
			playerDeathScreen()
		}
		clearScreen()
	}
}

func showStatusScreen(player character.Player, enemy character.Enemy) {
	clearScreen()

	fmt.Printf("Player HP: %d \nEnemy HP: %d\n", player.Hp, enemy.Hp)
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	showHeader()
}

type errorString struct {
	s string
}

func takeUserInput(availableChoices []string) (string, error) {
	var userInput string
	fmt.Scanln(&userInput)
	var validInput bool
	for _, item := range availableChoices {
		if userInput == item {
			validInput = true
		}
	}
	if validInput {
		fmt.Printf("")
	} else {
		return "", errors.New("This is not valid input")
	}
	return userInput, nil
}

func battleMenu(player character.Player, enemy character.Enemy) (character.Player, character.Enemy) {
	fmt.Printf("\n\n")
	for index, ability := range player.Abilities {
		fmt.Println(index+1, ability)
	}
	fmt.Printf("\n")
	possibleChoices := []string{"1", "2"}
	userInput, err := takeUserInput(possibleChoices)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	clearScreen()
	if userInput == "1" {
		player, enemy = player.AttackEnemy(enemy)
	}
	if userInput == "2" {
		player = player.DrinkPotion()
	}
	return player, enemy
}

func createPlayer() character.Player {
	abilities := []string{"Attack", "Drink Potion"}
	hero := character.Player{Name: "Jeremy", Hp: 100, Mp: 50, Attack: 20, Def: 10, Potions: 2, Speed: 5, Abilities: abilities}
	fmt.Println(hero.Name + " Enters the Fight!")
	return hero
}

func createEnemy() character.Enemy {
	abilities := []string{"Attack"}
	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 30
	attack := rand.Intn(max-min) + min
	enemy := character.Enemy{Name: "Dragon", Hp: 100, Mp: 50, Attack: attack, Def: 10, Speed: 1, Abilities: abilities}
	fmt.Println(enemy.Name + " Enters the Fight!")
	return enemy
}

func determineFirstPlay(player character.Player, enemy character.Enemy) string {
	var whoGoesFirst string
	if player.Speed > enemy.Speed {
		whoGoesFirst = "player"
	} else {
		whoGoesFirst = "enemy"
	}
	return whoGoesFirst
}

func showHeader() {
	fmt.Println("------Final Fantasy Battle Simulator------")
	fmt.Printf("\n")
}

func showTitleScreen() {
	clearScreen()
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
