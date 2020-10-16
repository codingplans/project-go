class Person {
    constructor(first, last, age, gender, interests) {
        this.name = {
            first,
            last
        };
        this.age = age;
        this.gender = gender;
        this.interests = interests;
    }

    async greeting() {
        let va = Promise.resolve(`22222Hi! I'm ${this.name.first + this.name.last}`)
        console.log(va)
        return await Promise.resolve(`Hi! I'm ${this.name.first + this.name.last}`);
    };

    farewell() {
        console.log(`${this.name.first} has left the building. Bye for now!`);
    };
}

let han = new Person('Han', 'Solo', 25, 'male', ['Smuggling']);

han.greeting().then(console.log);
han.farewell();
