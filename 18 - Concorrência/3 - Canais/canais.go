package main

import (
	"fmt"
	"time"
)

func main() {
	//Canal pode fazer envio e recebimento de informações
	canal := make(chan string)
	go escrever("Olá Mundo", canal)

	fmt.Println("Depois da função escrever começar a ser executada!")

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do Programa!")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
