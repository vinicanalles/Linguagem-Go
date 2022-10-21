package main

import "fmt"

func main() {
	//Canal com buffer só bloqueia quando atingir a capacidade máxima dele
	//Nesse caso só bloqueia quando atingir 2, logo não teremos deadlock
	//Canal sem buffer sempre irá bloquear com deadlock

	canal := make(chan string, 2) //declaração de canal com buffer

	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
