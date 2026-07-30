package main

import(
	"fmt"
)

func main()  {

	var opcao int
	var primeiroValor int
	var segundoValor int

	fmt.Println("===============")
	fmt.Println("  Calculadora ")
	fmt.Println("===============")
	fmt.Println("Escolha uma opção:")
	fmt.Println("[1] - Soma")
	fmt.Println("[2] - Subtração")
	fmt.Println("[3] - Multiplicação")
	fmt.Println("[4] - Divisão")
	fmt.Println("Você digitou: ")
	fmt.Scan(&opcao)	

	switch opcao {
	case 1:
		fmt.Println("Digite o primeiro valor: ")
		fmt.Scan(&primeiroValor)
		fmt.Println("Digite o segundo valor: ")
		fmt.Scan(&segundoValor)
		fmt.Println("Resultado: ")
		fmt.Println(soma(primeiroValor,segundoValor))
	
	case 2:
		fmt.Println("Digite o primeiro valor: ")
		fmt.Scan(&primeiroValor)
		fmt.Println("Digite o segundo valor: ")
		fmt.Scan(&segundoValor)
		fmt.Println("Resultado: ")
		fmt.Println(subtracao(primeiroValor,segundoValor))
		
	case 3:
		fmt.Println("Digite o primeiro valor: ")
		fmt.Scan(&primeiroValor)
		fmt.Println("Digite o segundo valor: ")
		fmt.Scan(&segundoValor)
		fmt.Println("Resultado: ") 
		fmt.Println(multiplicacao(primeiroValor,segundoValor))
	
	case 4:
		fmt.Println("Digite o primeiro valor: ")
		fmt.Scan(&primeiroValor)
		fmt.Println("Digite o segundo valor: ")
		fmt.Scan(&segundoValor)
		fmt.Println("Resultado: ")
		fmt.Println(divisao(primeiroValor,segundoValor))

	default:
		fmt.Println("Número digitado é inválido. Deve ser de 1 a 4.")
	}
}

func soma (x int, y int) int{
	return x + y
}

func subtracao (x int, y int) int{
	return x - y
}

func multiplicacao (x int, y int) int{
	return x*y
}

func divisao (x int, y int) int{
	return x/y
}