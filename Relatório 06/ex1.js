class Animal {
    constructor(nome, idade, especie) {
        this.nome = nome;
        this.idade = idade;
        this.especie = especie;
    }

    prinInfo() {
        console.log(`Nome: ${this.nome}, Idade: ${this.idade}, Espécie: ${this.especie}`);
    }
}

class Cachorro extends Animal {
    #raca;

    constructor(nome, idade, especie, raca) {
        super(nome, idade, especie);
        this.#raca = raca;
    }

    prinInfo() {
        super.prinInfo();
        console.log(`Raça: ${this.#raca}`);
    }
}

class Gato extends Animal {
    constructor(nome, idade, especie, cores) {
        super(nome, idade, especie);
        this.cores = cores;
    }

    prinInfo() {
        super.prinInfo();
        console.log(`Cores: ${this.cores.join(', ')}`);
    }
}

const animal = new Animal("Lobo", 10, "Canino");
const cachorro = new Cachorro("Floquinho", 9, "Canino", "Poodle");
const gato = new Gato("Sushi", 3, "Felino", ["laranja", "branco", "preto"]);

animal.prinInfo();
cachorro.prinInfo();
gato.prinInfo();
