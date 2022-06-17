package main

import "fmt"

func main() {
	fmt.Println("ponteiros")

	//PONTEIRO É UMA REFERENCIA DE MEMORIA
	var variavel3 int
	var ponteiro *int //transformando int em um referencia de memoria

	variavel3 = 100
	//Utilizar & para igualar um ponteiro a um endreço de memoria
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro) //desreferenciando

}
