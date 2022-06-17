package main

import "fmt"

func função1() {
	fmt.Println("Executando Função 1")
}

func função2() {
	fmt.Println("Executando Função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculada. Resultado será retornado")

	fmt.Println("Entrando na função para verificar se o aluno esta aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {
	fmt.Println(alunoEstaAprovado(7, 8))
}
