package main

import "fmt"

type Character struct {
	name string
	charClass string
	weapon string
	race string

	health int
	skillSlot int
	level int

	// backpack []string
}

func NewCharacter(_name, _charClass, _weapon, _race string) *Character {
	// return Character{_name, _charClass, _weapon, _race, 100, 2, 1, items()}
	return &Character{_name, _charClass, _weapon, _race, 100, 2, 1}
}

func (c Character) greeting() {
	fmt.Println("Hello", c.name)
}

func (c *Character) dmg() {
	fmt.Println(c.name, "takes 50 damage")
	c.health -= 50
}

// func items() []string {

// }

type Tavern struct {
	tavernName string
	hasInn bool
	questBoard bool
	city string
}

func NewTavern(tavernName string, hasInn bool, questBoard bool, city string) Tavern {
	return Tavern{tavernName, hasInn, questBoard, city}
}

func (t Tavern) rest(character *Character) {
	if t.hasInn == true {
		character.health += 25
	}
}

func main() {
	ryanDavis := NewCharacter("ryan", "warlock", "staff", "muggle")
	tavern := NewTavern("The Dodgey Tavern", true, true, "mordor")
	ryanDavis.greeting()
	fmt.Println(ryanDavis.health)
	ryanDavis.dmg()
	fmt.Println(ryanDavis.health)
	tavern.rest(ryanDavis)
	fmt.Println(ryanDavis.health)
}