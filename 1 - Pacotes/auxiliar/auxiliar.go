package auxiliar

import "fmt"

/*
- Funções escritas com letra minúscula só podem ser usadas dentro do próprio pacote que estão inseridas
- Funções escritas com letra maiúscula podem ser usadas e exportadas para outros pacotes
- Escrever registra uma mensagem na tela
*/
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2()
}
