package main

import "fmt"

var n int

func init() {
	//A ordem não afeta na execução
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}
