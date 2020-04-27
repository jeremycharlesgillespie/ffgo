package player

//Player : Does things
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

// New : Does things
//func New(name string, hp int, mp int, attack int, def int, potions int, abilities []string) Player {
//	e := Player{name, hp, mp, attack, def, potions, abilities}
//	return e
//}

//AttackEnemy : attacks an enemy
func (p Player) AttackEnemy() string {
	message := p.Name + " attacked!"
	return message
}
