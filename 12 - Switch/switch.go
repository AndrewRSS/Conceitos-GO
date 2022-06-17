package main

import "fmt"

func diasDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terceira-feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sabado"
	default:
		return "Numero invalido"
	}

}

func diaDaSemana2(numero int) string {

	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough
	case numero == 2:
		diaDaSemana = "Segunda-Feira"
	case numero == 3:
		diaDaSemana = "Terça-feira"
	case numero == 4:
		diaDaSemana = "Quarta-Feira"
	case numero == 5:
		diaDaSemana = "Quinta-Feira"
	case numero == 6:
		diaDaSemana = "Sexta-Feira"
	case numero == 7:
		diaDaSemana = "Sabado"
	default:
		diaDaSemana = "Numero Invalido"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia := diasDaSemana(10)
	fmt.Println(dia)
	fmt.Println("-------------------------------------")

	dia2 := diaDaSemana2(3)
	fmt.Println(dia2)

}
