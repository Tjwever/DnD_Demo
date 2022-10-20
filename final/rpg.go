package main

import "fmt"

///////////////////////////////////////////////////////////////////////////////
// CLASS OR STRUCT FOR CHARACTERS
///////////////////////////////////////////////////////////////////////////////
type Character struct {
	name                  string
	gender                string
	race                  string
	charClass             string
	weapon                string
	newStartingAttributes DefaultAttributes
}

type DefaultAttributes struct {
	life         float64
	skillSlot    int
	exp          float64
	level        int
	baseMeleeDmg float64
}

type Location struct {
	climate string
	terrain string
	country string
	city    string
}

type Tavern struct {
	name          string
	size          string
	hasInn        bool
	hasQuestBoard bool
}

// var defaultAttributes = NewStartingAttributes( 2, 0, 1 )

///////////////////////////////////////////////////////////////////////////////
// SETTING CONSTRUCTORS FOR CLASSES / INSTANTIATING OBJECTS
///////////////////////////////////////////////////////////////////////////////
func NewCharacter(_name, _gender, _race, _charClass, _weapon string) *Character {
	return &Character{_name, _gender, _race, _charClass, _weapon, *NewStartingAttributes()}
}

func NewStartingAttributes() *DefaultAttributes {
	return &DefaultAttributes{100, 2, 0.0, 1, 10.0}
}

func NewTavern(_name string, _size string, _hasInn bool, _hasQuestBoard bool) *Tavern {
	return &Tavern{_name, _size, _hasInn, _hasQuestBoard}
}

func Locale(_climate, _terrain, _country, _city string) *Location {
	return &Location{_climate, _terrain, _country, _city}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS / FUNCTION
///////////////////////////////////////////////////////////////////////////////
func (c *Character) home() {
	home := *Locale("sunny", "mountain", "mordor", "moreDoor")
	fmt.Println(c.name + " starts at " + home.country + " in the great city of " + home.city)
}

func (c *Character) applySkill() *string {
	signatureSkill := ""
	if c.charClass == "rogue" {
		signatureSkill = "Backstab"
	} else if c.charClass == "wizard" {
		signatureSkill = "FireBall"
	}
	return &signatureSkill
}

func (c *Character) expPoints() {
	c.newStartingAttributes.exp += 0.34
	c.levelUp()
}

func (c *Character) levelUp() {
	if c.newStartingAttributes.exp >= 1 {
		c.newStartingAttributes.level += 1
		if c.newStartingAttributes.level >= 2 {
			c.newStartingAttributes.baseMeleeDmg += 0.2
		}
		c.newStartingAttributes.exp = 0
	}
}

func (c *Character) greeting() {
	fmt.Println("Hello, my name is " + c.name + ". I am a " + c.race + " " + c.charClass)
}

func (c *Character) useSkill() {
	skill := c.newStartingAttributes.skillSlot
	if skill > 0 {
		fmt.Println(c.name + " uses " + *c.applySkill() + "!")
		c.newStartingAttributes.skillSlot = skill - 1
	} else {
		fmt.Println("You can't use", *c.applySkill(), ", you ran out of skill slots!")
		c.newStartingAttributes.skillSlot = 0
	}

}

func (c *Character) stats() {
	fmt.Println("('=============================')")
	fmt.Println("STATS FOR:", c.name)
	fmt.Println("CLASS:", c.charClass)
	fmt.Println("RACE:", c.race)
	fmt.Println("LEVEL:", c.newStartingAttributes.level)
	fmt.Println("EXP:", c.newStartingAttributes.exp)
	fmt.Println("SKILL SLOTS:", &c.newStartingAttributes.skillSlot)
	fmt.Println("SPECIAL:", *c.applySkill())
	fmt.Println("BASE DMG:", c.newStartingAttributes.baseMeleeDmg)
	fmt.Println("('=============================')")
}

func (t *Tavern) drink(character *Character) {
	if t.size == "small" {
		character.newStartingAttributes.life = character.newStartingAttributes.life * 0.25 + character.newStartingAttributes.life
		fmt.Println(character.name, "takes a drink at", t.name, "and replenishes their health by 25%")
	} else if t.size == "medium" {
		character.newStartingAttributes.life = character.newStartingAttributes.life * 0.50 + character.newStartingAttributes.life
		fmt.Println(character.name, "takes a drink at", t.name, "and replenishes their health by 50%")
	}
}

func (t *Tavern) sleep(character *Character) {
	if t.hasInn == true {
		fmt.Println("*********sleep method block start*********")
		fmt.Println(character)
		fmt.Println(character.name, "goes to sleep at", t.name, "Inn and is fully rested!")
		character.newStartingAttributes.life = 100
		character.newStartingAttributes.skillSlot = 2
		// fmt.Println(character.newStartingAttributes.life)
		// fmt.Println(character.newStartingAttributes.skillSlot)
		fmt.Println(character)
		fmt.Println("*********sleep method block end*********")
		} else {
			fmt.Println(character.name, "is attempting to go to sleep at", t.name, "Inn, but there is no Inn here...")
		}
}

func (c *Character) dmg() {
	c.newStartingAttributes.life = 50
	fmt.Println(c.name, "took damage!", "Health:", c.newStartingAttributes.life)
}

///////////////////////////////////////////////////////////////////////////////
// MAIN FUNCTION
///////////////////////////////////////////////////////////////////////////////
func main() {
	theRogue := *NewCharacter("Zap Branigan", "male", "human", "rogue", "dagger")
	theTavern := *NewTavern("The BullFrog Creek", "small", true, true)
	// theRogue.greeting()
	// theRogue.home()
	// theRogue.stats()
	// theRogue.stats()
	// fmt.Println("before using skill slot", theRogue.newStartingAttributes.skillSlot)
	// theRogue.useSkill()
	// fmt.Println("after using skill slot", theRogue.newStartingAttributes.skillSlot)
	// theRogue.useSkill()
	// theRogue.useSkill()
	// fmt.Println("after using skill slot", theRogue.newStartingAttributes.skillSlot)
	fmt.Println(theRogue.name, "Health:", theRogue.newStartingAttributes.life)
	fmt.Println(theRogue.name, "Skill Slot:", theRogue.newStartingAttributes.skillSlot)
	theRogue.useSkill()
	theRogue.dmg()
	fmt.Println("before resting skill slot is now", theRogue.newStartingAttributes.skillSlot)
	fmt.Println("before resting skill slot is now", theRogue.newStartingAttributes.life)
	theTavern.drink(&theRogue)
	fmt.Println("after resting skill slot is now", theRogue.newStartingAttributes.skillSlot)
	fmt.Println("after resting skill slot is now", theRogue.newStartingAttributes.life)

	
	// theRogue.stats()

	// theWizard := *NewCharacter("Henry Plopper", "male", "human", "wizard", "staff")
	// theWizard.stats()
	// theWizard.greeting()
	// theWizard.home()
	// theWizard.useSkill()

}

// Something to add.  Want to demo how to deal with lists/arrays.  Was thinking about implementing a storage cache, or stash.
// Maybe every character can start with a stash with a few items inside, potions, armor and such.  Stash may end up being it's own
// class in the future, but we can use a simple property on the Character class for now.
