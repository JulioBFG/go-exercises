package main

import (
	"fmt"
	"sort"
)

func leituraNumero() {
	fmt.Println("Digite Um numero: ")
	var number int
	fmt.Scanln(&number)
	if number >= 10 {
		fmt.Println("Seu numero é maior ou igual a 10")
	} else {
		fmt.Println("Seu numero é menor que 10")
	}
}

func calcularAreaTerreno() {
	fmt.Println("Digite o tamanho da base do terreno")
	var base int
	fmt.Scanln(&base)

	fmt.Println("Digite o tamanho da altura do terreno")
	var altura int
	fmt.Scanln(&altura)

	var resultado = base * altura
	fmt.Print("A Area do terreno é de: ")
	fmt.Print(resultado)
	fmt.Println("m2")
}

func contadorDiasVida() {
	fmt.Println("Qual seu nome")
	var nome string
	fmt.Scanln(&nome)

	fmt.Println("Qual sua idade")
	var idade int
	fmt.Scanln(&idade)

	var idadeCalculada = idade * 365
	concatenado := fmt.Sprint(nome, ", você viveu o total de: ", idadeCalculada, " dias")
	fmt.Println(concatenado)

}

func ordernarProdutos() {
	produtos := make(map[string]int)

	for i := 0; i <= 4; i++ {
		fmt.Println("Qual o nome do produto")
		var nomeProduto string
		fmt.Scanln(&nomeProduto)
		fmt.Println("Qual o valor do produto")
		var valorProduto int
		fmt.Scanln(&valorProduto)
		produtos[nomeProduto] = valorProduto
	}
	ordenando := map[int]string{}
	chavesOrdenacao := []int{}
	for key, val := range produtos {
		ordenando[val] = key
		chavesOrdenacao = append(chavesOrdenacao, val)
	}
	sort.Ints(chavesOrdenacao)

	for _, val := range chavesOrdenacao {
		fmt.Println(ordenando[val], val)
	}
}

func calculaInss() {
	fmt.Println("qual seu salario?")
	var salario int
	fmt.Scanln(&salario)
	if salario <= 600 {
		fmt.Println("Você está isento de INSS ")
	} else if salario <= 1200 {
		fmt.Println("Você terá que pagar 20% de INSS ")
	} else if salario > 1200 && salario <= 2000 {
		fmt.Println("Você terá que pagar 25% de INSS ")
	} else {
		fmt.Println("Você terá que pagar 30% de INSS ")
	}
}

func main() {
	var idadePerfil = 20
	var hobbiePerfil = "nadar"
	var corPerfil = "azul"
	var corOlhosPerfil = "castanhos"
	var cidadePerfil = "Joinville"

	fmt.Println("Qual sua idade")
	var idade int
	fmt.Scanln(&idade)

	fmt.Println("Qual seu hobbie preferido?")
	var hobbie string
	fmt.Scanln(&hobbie)

	fmt.Println("Qual sua cor preferida?")
	var cor string
	fmt.Scanln(&cor)

	fmt.Println("Qual a cor dos seus  olhos ?")
	var corOlhos string

	fmt.Scanln(&corOlhos)
	fmt.Println("Qual a sua cidade ?")
	var cidade string
	fmt.Scanln(&cidade)

	var counter int
	if idade == idadePerfil {
		counter++
	}
	if hobbie == hobbiePerfil {
		counter++
	}
	if cor == corPerfil {
		counter++
	}
	if corOlhos == corOlhosPerfil {
		counter++
	}
	if cidade == cidadePerfil {
		counter++
	}

	if counter > 3 {
		fmt.Println("MATCH!")
	} else {
		fmt.Println("sem Perfil compativel com o seu")
	}

}
