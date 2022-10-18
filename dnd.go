package main

import "fmt"

// Character Class /////////////////////////////////
type Character struct {
	name string
	charClass string
	weapon string
	race string
	newAttributes DefaultAttributes
}

func NewCharacter(name, charClass, weapon, race string) Character {
	return Character{name, charClass, weapon, race, NewAttributes()}
}

type DefaultAttributes struct {
	health int
	skillSlot int
	lvl int
}

func NewAttributes() DefaultAttributes {
	return DefaultAttributes{100, 2, 1}
}

func (c Character) greeting() {
	fmt.Println("Hello, my name is", c.name, "I am a", c.charClass, c.race)
}
// Character Class /////////////////////////////////

// Tavern Class /////////////////////////////////
type Tavern struct {
	tavernName string
	hasInn bool
	hasQuestBoard bool
	city string
}

func NewTavern(tavernName string, hasInn bool, hasQuestBoard bool, city string) Tavern {
	return Tavern{tavernName, hasInn, hasQuestBoard, city}
}

func (t Tavern) rest(character *Character) {
	if t.hasInn == true {
		character.newAttributes.health = 150
	}
}
// Tavern Class /////////////////////////////////

func main() {
	ryanDavis := NewCharacter("ryan", "warlock", "staff", "muggle")
	tavern := NewTavern("The Tavern", true, false, "Mordor")
	ryanDavis.greeting()
	fmt.Println(tavern)
	fmt.Println(ryanDavis.newAttributes.health)
	tavern.rest(&ryanDavis)
	fmt.Println(ryanDavis.newAttributes.health)
}