class Character {
    constructor(_name, _gender, _race, _charClass, _weapon) {
        this.name = _name
        this.gender = _gender
        this.race = _race
        this.charClass = _charClass
        this.weapon = _weapon

        this.skillSlot = 2
        this.exp = 0
        this.lvl = 1
        this.baseMeleeDmg = 10

        this.signatureSkill;
        if(this.charClass === 'Rogue') {
            this.signatureSkill = 'Backstab'
        } else if(this.charClass === 'Wizard') {
            this.signatureSkill = 'FireBall'
        }
    }

    expPoinnts() {
        this.exp += 0.34
        this.levelUp()
    }

    levelUp() {
        if(this.exp >= 1) {
            this.lvl++
            if(this.lvl >= 2) {
                this.baseMeleeDmg += 0.2
            }
            this.exp = 0
        }
    }

    greeting() {
        let something = 0
        // console.log('Hello, I am ' + this.name + '. I am a ' + this.race + ' ' + this.charClass)
        console.log(`Hello, I am ${this.name}. I am a ${this.race} ${this.charClass}`)
    }

    home() {
        let homeLocation = new Location('sunny', 'mountain', 'USA', 'tampa')
        if(this.lvl === 1) {
            console.log(this.name + ' starts at ' + homeLocation.country + ' ' + homeLocation.city)
        }
        else {
            console.log(this.name + ' has moved on...')
        }
    }

    attack() {
        let randomDmg = Math.floor(Math.random() * this.baseMeleeDmg) + 1
        let roll = Math.floor(Math.random() * 20) + 1

        if(roll === 20) {
            randomDmg = this.baseMeleeDmg * 2
            console.log(`${this.name} rolled a natural 20!!  CRITICAL HIT!!!`)
            console.log(`${randomDmg} points of damage to ${this.name}'s target!!`)
        } else if(roll === 1) {
            if(randomDmg > 5) {
                randomDmg = Math.floor(5 / 2)
            } else {
                randomDmg = Math.floor(randomDmg / 2)
            }
            console.log(`${this.name} rolled a natural 1...`)
            console.log(`${randomDmg} points of damage to ${this.name}'s target`)
        } else {
            console.log(`${this.name} rolled a ${roll}.`)
            console.log(`${this.name} dealt ${randomDmg} points of damage to the target.`)
        }

        this.expPoinnts()
    }

    useSkill() {
        if(this.skillSlot > 0) {
            console.log(`${this.name} uses ${this.signatureSkill}!`)
            this.skillSlot--
            this.expPoinnts()

        } else {
            this.skillSlot = 0
            console.log(`You can't use ${this.signatureSkill}, you ran out of skill slots!`)
        }
    }

    stats() {
        console.log('=============================')
        console.log('STATS FOR: ' + this.name)
        console.log('CLASS: ' + this.charClass)
        console.log('RACE: ' + this.race)
        console.log('LEVEL: ' + this.lvl)
        console.log('EXP: ' + this.exp)
        console.log('SKILL SLOTS: ' + this.skillSlot)
        console.log('SPECIAL: ' + this.signatureSkill)
        console.log('BASE DMG: ' + this.baseMeleeDmg)
        console.log('=============================')
    }
    
}

class Location {
    constructor(_climate, _terrain, _country, _city) {
        this.climate = _climate
        this.terrain = _terrain
        this.country = _country
        this.city = _city

    }

}

let theRogue = new Character('Zap Branigan', 'male', 'human', 'Rogue', 'dagger');
let theWizard = new Character('Gandolf', 'male', 'human', 'Wizard', 'wand');

// theRogue.greeting()
// theWizard.greeting()

// console.log(theRogue.signatureSkill)
// console.log(theWizard.signatureSkill)

// theRogue.stats()
// theWizard.stats()

// theRogue.useSkill()
// theRogue.stats()
theRogue.stats()
theRogue.attack()
theRogue.attack()
theRogue.attack()
theRogue.attack()
theRogue.attack()
theRogue.attack()
theRogue.stats()