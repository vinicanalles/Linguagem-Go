package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main") //escrevendo uma mensagem
	auxiliar.Escrever()                       //usando uma função de outro pacote para escrever uma mensagem

	erro := checkmail.ValidateFormat("133") //exemplo de utilização de uma biblioteca externa importada dentro do projeto
	fmt.Println(erro)
}
