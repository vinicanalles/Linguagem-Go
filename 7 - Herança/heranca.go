package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"João", "Pedro", 20, 178}
	fmt.Println(p1)

	estudante1 := estudante{p1, "Engenharia", "Usp"}
	fmt.Println(estudante1.nome)
}
