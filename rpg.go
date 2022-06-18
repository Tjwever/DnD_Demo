package main

import "fmt"

///////////////////////////////////////////////////////////////////////////////
// CLASS OR STRUCT FOR CHARACTERS
///////////////////////////////////////////////////////////////////////////////
type Character struct {
	name string
	gender string
	race string
	charClass string
	weapon string
	
}

type Location struct {
	climate string
	terrain string
	country string
	city string
}

type DefaultAttributes struct {
	skillSlot int
	exp int
	level int
}

///////////////////////////////////////////////////////////////////////////////
// SETTING CONSTRUCTORS FOR CLASSES / INSTANTIATING OBJECTS
///////////////////////////////////////////////////////////////////////////////
func NewCharacter(_name, _gender, _race, _charClass, _weapon string) *Character {
	return &Character{ _name, _gender, _race, _charClass, _weapon }
}

func StartingAttributes(_skillSlot, _exp, _level int) DefaultAttributes {
	return DefaultAttributes{ 2, 0, 1 }
}

func Locale(_climate, _terrain, _country, _city string) *Location {
	return &Location{ _climate, _terrain, _country, _city }
}

///////////////////////////////////////////////////////////////////////////////
// METHODS / FUNCTION
///////////////////////////////////////////////////////////////////////////////
func (c Character) home() {
	home := *Locale("sunny", "mountain", "mordor", "moreDoor")
	fmt.Println(c.name + " starts at " + home.country + " in the great city of " + home.city)
}

func (c Character) applySkill() *string {
	signatureSkill := ""
	if c.charClass == "rogue" {
		signatureSkill = "Backstab"
	} else if c.charClass == "wizard" {
		signatureSkill = "FireBall"
	}
	return &signatureSkill
}

func (c Character) greeting() {
	fmt.Println("Hello, my name is " + c.name)
}

func (c Character) specialAttack() {
	fmt.Println(c.name + " uses " + *c.applySkill() + "!")
}

///////////////////////////////////////////////////////////////////////////////
// MAIN FUNCTION
///////////////////////////////////////////////////////////////////////////////
func main() {
	theRogue := Character{ "Zap Branigan", "male", "human", "rogue", "dagger" }
	theRogue.home()
	// theRogue.greeting()
	// theRogue.specialAttack()

	theWizard := Character{ "Henry Plopper", "male", "human", "wizard", "staff" }
	theWizard.home()
	// theWizard.greeting()
	// theWizard.specialAttack()

	
}