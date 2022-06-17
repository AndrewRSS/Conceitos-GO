package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")

	i := 0
	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		//time.Sleep(time.Second)
	}

	for j := 0; j < 10; j += 2 {
		fmt.Println("Incrementando i", j)
		//	time.Sleep(time.Second)
	}

	nomes := [3]string{"joão", "Davi", "Lucas"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, value := range usuario {
		fmt.Println(chave, value)
	}

}
