package character

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//Character : This is the default template for any character, good or bad.
//type Character struct {
//}

//Player : He's the good guy
type Player struct {
	Name      string
	Hp        int
	Mp        int
	Attack    int
	Def       int
	Potions   int
	Speed     int
	Abilities []string
}

//Enemy : He's a bad guy
type Enemy struct {
	Name      string
	Hp        int
	Mp        int
	Attack    int
	Def       int
	Speed     int
	Abilities []string
}

//AttackEnemy : This is a Player Skill. Takes an Enemy object as an input and returns the Player and Enemy objects.
func (player Player) AttackEnemy(enemy Enemy) (Player, Enemy) {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 5
	attackVariance := rand.Intn(max-min) + min
	totalAttack := player.Attack + attackVariance
	playerAttack := strconv.Itoa(totalAttack)
	message := player.Name + " attacked the " + enemy.Name + " for " + playerAttack + " points of damage!"
	fmt.Println(message)
	enemy.Hp = enemy.Hp - totalAttack
	return player, enemy
}

//AttackPlayer : This is a Enemy Skill. Takes a Player object as an input and returns the Enemy and Player objects.
func (enemy Enemy) AttackPlayer(player Player) (Enemy, Player) {
	rand.Seed(time.Now().UnixNano())
	min := 2
	max := 10
	attackVariance := rand.Intn(max-min) + min
	totalAttack := enemy.Attack + attackVariance
	enemyAttack := strconv.Itoa(totalAttack)
	message := "The " + enemy.Name + " attacked " + player.Name + " for " + enemyAttack + " points of damage!"
	fmt.Println(message)
	player.Hp = player.Hp - totalAttack
	return enemy, player
}

//DrinkPotion : This is a Player skill. It is used to drink a potion to restore HP, if there is one available.
func (player Player) DrinkPotion() Player {
	if player.Potions > 0 {
		player.Hp = player.Hp + 50
		player.Potions = player.Potions - 1
		fmt.Println(player.Name + " drank a potion to restore HP!")
	} else {
		fmt.Println(player.Name + " does not have any more potions!")
	}
	return player
}
