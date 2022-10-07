class Character {
    constructor(_name, _charClass, _weapon, _race) {
        this.name = _name
        this.charClass = _charClass
        this.weapon = _weapon
        this.race = _race

        this.health = 100
        this.specialSkill

        if(this.charClass === 'warlock') {
            this.specialSkill = 'fireball'
        } else if(this.charClass === 'rogue') {
            this.specialSkill = 'backstab'
        } else {
            this.specialSkill = 'head punch'
            
        }
        
    }
    
    greeting() {
        // console.log('Hello, my name is ' + this.name + 'I am a ' + this.race + ' ' + this.charClass)
        console.log(`Hello, my name is ${this.name} I am a ${this.race} ${this.charClass}`);
    }
    
}

class Tavern {
    constructor(_tavernName, _hasInn, _questBoard, _city) {
        this.tavernName = _tavernName;
        this.hasInn = _hasInn;
        this.questBoard = _questBoard;
        this.city = _city;
    }

    rest(character) {
        if(this.hasInn === true) {
            character.health += 50
        }
    }
}

let ryan = new Character('ryan', 'warlock', 'staff', 'muggle');
let tavern = new Tavern('The Tavern', true, false, 'mordor')

// console.log(ryan)

// ryan.greeting()

// console.log(`${ryan.name} uses ${ryan.specialSkill}`)

ryan.health = 25

console.log(ryan.health)

tavern.rest(ryan)

console.log(ryan.health)

// let person = {
//     name: 'tim',
//     age: 25,
//     hair_color: 'brown'
// }

// // console.log(person)

// class Car {
//     constructor(_make, _model, _color) {
//         this.make = _make
//         this.model = _model
//         this.color = _color
//     }

// }
// let factory = []

// let truck1 = new Car('chevy', 'silverado', 'white')
// console.log(truck1.model)

