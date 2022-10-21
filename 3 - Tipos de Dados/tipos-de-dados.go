package main

import (
	"errors"
	"fmt"
)

func main() {
	//INÍCIO NÚMEROS INTEIROS
	numero := -1000000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias
	//INT32 = RUNE
	var numero3 rune = 12456 //rune se refere ao int32
	fmt.Println(numero3)

	//UINT8 = BYTE
	var numero4 byte = 123
	fmt.Println(numero4)
	// FIM NÚMEROS INTEIROS

	//INÍCIO NÚMEROS REAIS
	var numeroReal float32 = 123.45
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 123000000000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)
	//FIM NÚMEROS REAIS

	// INÍCIO STRINGS
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)
	//FIM STRINGS

	texto := 5
	fmt.Println(texto)

	var booleano1 bool //inicia em false
	fmt.Println(booleano1)

	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)
}
