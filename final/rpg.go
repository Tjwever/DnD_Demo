package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/////////////////////////////////////////////////////////////////////////////////////
// CHARACTER
/////////////////////////////////////////////////////////////////////////////////////

// FAUX CLASS DATA TYPE CHARACTER
type Character struct {
	name                  string
	gender                string
	race                  string
	charClass             string
	weapon                string
	newStartingAttributes DefaultAttributes
	backpack          	  []string
}

// FAUX CONSTRUCTOR FOR STRUCT DATA TYPE  //  FOR CHARACTER
func NewCharacter(_name, _gender, _race, _charClass, _weapon string) *Character {
	return &Character{_name, _gender, _race, _charClass, _weapon, *NewStartingAttributes(), *backPack(_charClass, _weapon)}
}

// FUNCTIONS OR METHODS THAT BELONG TO THE CHARACTER DATA TYPE
func (c *Character) home() {
	home := *NewLocation("sunny", "mountain", "mordor", "moreDoor")
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

func (c *Character) attack() {
	rand.Seed(time.Now().UnixNano())
	randomDmg := rand.Float64()*c.newStartingAttributes.baseMeleeDmg
	roll := rand.Intn(20) + 1

	if roll == 20 {
		randomDmg = c.newStartingAttributes.baseMeleeDmg * 2
		fmt.Println(c.name, "rolled a natural 20!! CRITIAL HIT!!!" )
		fmt.Println(randomDmg, "points of damage to", c.name + "'s target!!")
	} else if roll == 1 {
		// you're either doing 2 point of damage, or 0 points...rough...
		if randomDmg > 5 {
			randomDmg = math.Ceil(rand.Float64()* (5 / 2))
		} else {
			randomDmg = math.Floor(rand.Float64()* (randomDmg / 2))
		}
		fmt.Println("Oof..", c.name, "rolled a 1..." )
		fmt.Println(randomDmg, "points of damage to", c.name + "'s target.")
	} else {
		fmt.Println(c.name, "rolled a", roll)
		fmt.Println(c.name, "dealt", randomDmg, "'s of damange")

	}

	c.expPoints()
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

func (c *Character) dmg() {
	c.newStartingAttributes.life = 50
	fmt.Println(c.name, "took damage!", "Health:", c.newStartingAttributes.life)
}

// Since we can't use 'this.name' or anything that belongs to Character, we have to 
// think about the way we create this differently.  We also can't bind Character to this
// function to make it a method and call it in our 'constructor'.  So, we'll have to pass
// in some values that will take in the parameters that we use for our constructor
// When charClass is called, it'll turn into whatever we pass in from the constructor.
// Example: charClass = 'rogue' and weapon = 'dagger'
func backPack(charClass, weapon string) *[]string {
	pack := []string{}
	if charClass == "rogue" {
		for i := 0; i < 5; i++ {
			pack = append(pack, "health potion")
		}
		pack = append(pack, weapon)
		pack = append(pack, "leather tunic")
		pack = append(pack, "leather cap")
		pack = append(pack, "leather leggings")
	} else {
		for i := 0; i < 2; i++ {
			pack = append(pack, "health potion")
		}
		for i := 0; i < 3; i++ {
			pack = append(pack, "mana potion")
		}
		pack = append(pack, weapon)
		pack = append(pack, "cloth robe")
		pack = append(pack, "cloth cap")
		pack = append(pack, "cloth leggings")
	}
	return &pack
}

func (c *Character) stats() {
	fmt.Println("('=============================')")
	fmt.Println("STATS FOR:", c.name)
	fmt.Println("CLASS:", c.charClass)
	fmt.Println("RACE:", c.race)
	fmt.Println("LEVEL:", c.newStartingAttributes.level)
	fmt.Println("EXP:", c.newStartingAttributes.exp)
	fmt.Println("SKILL SLOTS:", c.newStartingAttributes.skillSlot)
	fmt.Println("SPECIAL:", *c.applySkill())
	fmt.Println("BASE DMG:", c.newStartingAttributes.baseMeleeDmg)
	fmt.Println("('=============================')")
}

/////////////////////////////////////////////////////////////////////////////////////
// CHARACTER END
/////////////////////////////////////////////////////////////////////////////////////


/////////////////////////////////////////////////////////////////////////////////////
// DEFAULT ATTRIBUTES
/////////////////////////////////////////////////////////////////////////////////////

// FAUX CLASS DATA TYPE ATTRIBUTES
type DefaultAttributes struct {
	life         float64
	skillSlot    int
	exp          float64
	level        int
	baseMeleeDmg float64
}

// FAUX CONSTRUCTOR FOR STRUCT DATA TYPE  //  FOR ATTRIBUTES
func NewStartingAttributes() *DefaultAttributes {
	return &DefaultAttributes{100, 2, 0.0, 1, 10.0}
}

/////////////////////////////////////////////////////////////////////////////////////
// DEFAULT ATTRIBUTES END
/////////////////////////////////////////////////////////////////////////////////////


/////////////////////////////////////////////////////////////////////////////////////
// LOCATION
/////////////////////////////////////////////////////////////////////////////////////

// FAUX CLASS DATA TYPE LOCATION
type Location struct {
	climate string
	terrain string
	country string
	city    string
}

// FAUX CONSTRUCTOR FOR STRUCT DATA TYPE  //  FOR LOCATION
func NewLocation(_climate, _terrain, _country, _city string) *Location {
	return &Location{_climate, _terrain, _country, _city}
}

/////////////////////////////////////////////////////////////////////////////////////
// LOCATION END
/////////////////////////////////////////////////////////////////////////////////////


/////////////////////////////////////////////////////////////////////////////////////
// TAVERN 
/////////////////////////////////////////////////////////////////////////////////////

// FAUX CLASS DATA TYPE TAVERN
type Tavern struct {
	name          string
	size          string
	hasInn        bool
	hasQuestBoard bool
}

// FAUX CONSTRUCTOR FOR STRUCT DATA TYPE  //  FOR TAVERN
func NewTavern(_name string, _size string, _hasInn bool, _hasQuestBoard bool) *Tavern {
	return &Tavern{_name, _size, _hasInn, _hasQuestBoard}
}

// FUNCTIONS OR METHODS THAT BELONG TO THE TAVERN DATA TYPE
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
		character.newStartingAttributes.life = 100
		character.newStartingAttributes.skillSlot = 2
		fmt.Println(character.name, "goes to sleep at", t.name, "Inn and is fully rested!")
		} else {
			fmt.Println(character.name, "is attempting to go to sleep at", t.name, "Inn, but there is no Inn here...")
		}
	}

/////////////////////////////////////////////////////////////////////////////////////
// TAVERN END
/////////////////////////////////////////////////////////////////////////////////////
	

///////////////////////////////////////////////////////////////////////////////
// MAIN FUNCTION
///////////////////////////////////////////////////////////////////////////////
func main() {
	theRogue := *NewCharacter("Zap Branigan", "male", "human", "rogue", "dagger")
	// theTavern := *NewTavern("The BullFrog Creek", "small", true, true)
	theRogue.greeting()
	// theRogue.home()
	// theRogue.stats()
	// theRogue.stats()
	// fmt.Println("before using skill slot", theRogue.newStartingAttributes.skillSlot)
	// theRogue.useSkill()
	// fmt.Println("after using skill slot", theRogue.newStartingAttributes.skillSlot)
	// theRogue.useSkill()
	// theRogue.useSkill()
	// fmt.Println("after using skill slot", theRogue.newStartingAttributes.skillSlot)
	// fmt.Println(theRogue.name, "Health:", theRogue.newStartingAttributes.life)
	// fmt.Println(theRogue.name, "Skill Slot:", theRogue.newStartingAttributes.skillSlot)
	// theRogue.useSkill()
	// theRogue.dmg()
	// fmt.Println("before resting skill slot is now", theRogue.newStartingAttributes.skillSlot)
	// fmt.Println("before resting skill slot is now", theRogue.newStartingAttributes.life)
	// theTavern.drink(&theRogue)
	// fmt.Println("after resting skill slot is now", theRogue.newStartingAttributes.skillSlot)
	// fmt.Println("after resting skill slot is now", theRogue.newStartingAttributes.life)

	// fmt.Println(theRogue.backpack)
	
	// theRogue.stats()
	
	// theWizard := *NewCharacter("Henry Plopper", "male", "human", "wizard", "staff")
	// theWizard.stats()
	// theWizard.greeting()
	// theWizard.home()
	// theWizard.useSkill()
	// fmt.Println(theWizard.backpack)
	
}

