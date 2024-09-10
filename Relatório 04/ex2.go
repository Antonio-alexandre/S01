package main

import (
  "fmt"
  "math/rand"
  "time"
)

//Calcula o fatorial do numero
func fatorial(n int) int {
  if n == 0 {
    return 1
  }
  resultado := 1
  for i := 1; i <= n; i++ {
    resultado *= i
  }
  return resultado
}

func main() {
  rand.Seed(time.Now().UnixNano()) //Inicializa o gerador de numeros aleatórios
 
  num := rand.Intn(11) //Gera um numero entre 0 e 10

  fat := fatorial(num) //Chama a função fatorial

  fmt.Printf("Número aleatório gerado: %d\n", num)
  fmt.Printf("Fatorial de %d é: %d\n", num, fat)
}