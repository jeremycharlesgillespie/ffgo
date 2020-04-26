package player

//Player : Does things
type Player struct {
	Name      string
	Hp        int
	Mp        int
	Attack    int
	Def       int
	Potions   int
	Abilities []string
}

// New : Does things
//func New(name string, hp int, mp int, attack int, def int, potions int, abilities []string) Player {
//	e := Player{name, hp, mp, attack, def, potions, abilities}
//	return e
//}

//Taco : is a taco
var Taco = "asdf"

func attackEnemy() string {
	return "Attacked!"
}
