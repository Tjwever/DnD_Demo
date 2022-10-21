package main

import "fmt"

/////////////////////////////////////////////////////////////////////////////////////
// CHARACTER
/////////////////////////////////////////////////////////////////////////////////////

// FAUX CLASS DATA TYPE CHARACTER
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

// FAUX CONSTRUCTOR FOR STRUCT DATA TYPE  //  FOR CHARACTER
func NewCharacter(_name, _charClass, _weapon, _race string) *Character {
	// return Character{_name, _charClass, _weapon, _race, 100, 2, 1, items()}
	return &Character{_name, _charClass, _weapon, _race, 100, 2, 1}
}

// FUNCTIONS OR METHODS THAT BELONG TO THE CHARACTER DATA TYPE
func (c Character) greeting() {
	fmt.Println("Hello", c.name)
}

func (c *Character) dmg() {
	fmt.Println(c.name, "takes 50 damage")
	c.health -= 50
}

/////////////////////////////////////////////////////////////////////////////////////
// CHARACTER END
/////////////////////////////////////////////////////////////////////////////////////


/////////////////////////////////////////////////////////////////////////////////////
// TAVERN
/////////////////////////////////////////////////////////////////////////////////////

// FAUX CLASS DATA TYPE TAVERN
type Tavern struct {
	tavernName string
	hasInn bool
	questBoard bool
	city string
}

// FAUX CONSTRUCTOR FOR STRUCT DATA TYPE  //  FOR TAVERN
func NewTavern(tavernName string, hasInn bool, questBoard bool, city string) Tavern {
	return Tavern{tavernName, hasInn, questBoard, city}
}

// FUNCTIONS OR METHODS THAT BELONG TO THE CHARACTER DATA TYPE
func (t Tavern) rest(character *Character) {
	if t.hasInn == true {
		character.health += 25
	}
}

/////////////////////////////////////////////////////////////////////////////////////
// TAVERN END
/////////////////////////////////////////////////////////////////////////////////////

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