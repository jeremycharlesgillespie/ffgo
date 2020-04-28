package character

import (
	"fmt"
)

//Character : This is the default template for any character, good or bad.
type Character struct {
}

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

//AttackEnemy : This is a Player Skill. Attacks an enemy and returns both the Player and Enemy objects.
func (player Player) AttackEnemy(enemy Enemy) (Player, Enemy) {
	message := player.Name + " attacked the " + enemy.Name + "!"
	fmt.Println(message)
	enemy.Hp = enemy.Hp - player.Attack
	return player, enemy
}

//AttackPlayer : This is a Enemy Skill. Attacks the player and returns both the Enemy and Player objects.
func (enemy Enemy) AttackPlayer(player Player) (Enemy, Player) {
	message := enemy.Name + " attacked " + player.Name + "!"
	fmt.Println(message)
	player.Hp = player.Hp - enemy.Attack
	return enemy, player
}
