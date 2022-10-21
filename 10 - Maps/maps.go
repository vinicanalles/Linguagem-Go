package main

import "fmt"

func main() {

	//Dentro do colchete é o tipo das chaves
	//Fora do colchete é o tipo dos valores
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campos": "Campus 1",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "nome") //Removendo uma chave de dentro do map

	fmt.Println(usuario2)

	//Inserindo uma nova chave e valores dentro do map
	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}

	fmt.Println(usuario2)
}
