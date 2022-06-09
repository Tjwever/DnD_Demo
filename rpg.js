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

        this.signatureSkill;
        if(this.charClass === 'rogue') {
            this.signatureSkill = 'backstab'
        } else if(this.charClass === 'wizard') {
            this.signatureSkill = 'fireball'
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

    stats() {
        console.log('=============================')
        console.log('STATS FOR: ' + this.name)
        console.log('CLASS: ' + this.charClass)
        console.log('RACE: ' + this.race)
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

let theRogue = new Character('Zap Branigan', 'male', 'human', 'rogue', 'dagger');
let theWizard = new Character('Gandolf', 'male', 'human', 'wizard', 'wand');

function something() {

}

theRogue.greeting()
// theWizard.greeting()

// console.log(theRogue.signatureSkill)
// console.log(theWizard.signatureSkill)

theRogue.stats()
theWizard.stats()