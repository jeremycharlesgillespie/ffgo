package enemy

//Enemy : is a bad guy
type Enemy struct {
	Name      string
	Hp        int
	Mp        int
	Attack    int
	Def       int
	Speed     int
	Abilities []string
}

//AttackPlayer : attacks the player
//func (p Enemy) AttackPlayer(player player) string {
//	message := p.Name + " attacked " + player.Name + "!"
//	fmt.Println(message)
//}
