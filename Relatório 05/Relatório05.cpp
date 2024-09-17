#include <iostream>
using namespace std;

class Pessoa {
public:
  string nome;
  int idade;
  Pessoa(int idade, string nome) {
    this->idade = idade;
    this->nome = nome;
  }

  void imprimirNome() { cout << "O nome é: " << nome << endl; }
};

class Professor : public Pessoa {
public:
  Professor(int idade, string nome) : Pessoa(idade, nome){};
  void imprimirIdade() { cout << "A idade do professor é: " << idade << endl; }
};

class Aluno : public Pessoa {
public:
  Aluno(int idade, string nome) : Pessoa(idade, nome){};
  void imprimirIdade() { cout << "A idade do aluno é: " << idade << endl; }

private:
  void matricula() { cout << "Aluno matriculado" << endl; }
};

int main() {
  Pessoa pessoa(18, "Antonio");
  Professor professor(29, "Marcelo");
  Aluno aluno(22, "João");

  pessoa.imprimirNome();
  professor.imprimirNome();
  professor.imprimirIdade();
  aluno.imprimirNome();
  aluno.imprimirIdade();
}