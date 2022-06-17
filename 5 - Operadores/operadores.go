package main

import "fmt"

func main() {
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	somat := numero1 + numero2
	fmt.Println(somat)

	// Atibruição

	var variavel string = "String"
	variavel2 := "String2"
	fmt.Println(variavel, variavel2)

	// Operadores relacionais

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	fmt.Println("--------------------------------")

	// operadores logicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	fmt.Println("--------------------------------")

	// Operadores unarios

	numero := 10
	numero++
	numero += 15
	numero--
	numero -= 20
	numero *= 3
	numero /= 10
	numero %= 3

	fmt.Println(numero)

	// Operador ternario

	var texto string

	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)
}
