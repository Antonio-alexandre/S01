package main

import (
"fmt"
"math"
)

func main() {
var a, b, c float64

//Entra com os valores de a, b e c
fmt.Print("Entre com o valor de a: ")
fmt.Scan(&a)
fmt.Print("Entre com o valor de b: ")
fmt.Scan(&b)
fmt.Print("Entre com o valor de c: ")
fmt.Scan(&c)

delta := b*b - 4*a*c

if delta > 0 {

x1 := (-b + math.Sqrt(delta)) / (2 * a)
x2 := (-b - math.Sqrt(delta)) / (2 * a)
fmt.Printf("As raízes são: x1 = %.2f e x2 = %.2f\n", x1, x2)
} else if delta == 0 {

x := -b / (2 * a)
fmt.Printf("A raiz é: x = %.2f\n", x)
} else {

pReal := -b / (2 * a)
pImag := math.Sqrt(-delta) / (2 * a)
fmt.Printf("As raízes são: x1 = %.2f + %.2fi e x2 = %.2f - %.2fi\n", pReal, pImag, pReal, pImag)
}
}