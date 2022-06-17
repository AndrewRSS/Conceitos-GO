package main

import (
	"errors"
	"fmt"
)

func main() {

	var numero int64 = 1000000000000000
	fmt.Println(numero)

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	//alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	//Byte = Uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12300000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12300000000.45
	fmt.Println(numeroReal3)

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'a'
	fmt.Println(char)

	var texto int16
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var booleano2 bool
	fmt.Println(booleano2)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
