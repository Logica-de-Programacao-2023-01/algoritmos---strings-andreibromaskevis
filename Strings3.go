package main

import (
	"fmt"
	"strings"
)
func main() {
	var nome string
	var Al, Nl string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&nome)
	fmt.Print("Digite o caractere a ser substituído: ")
	fmt.Scanln(&Al)
	fmt.Print("Digite o novo caractere: ")
	fmt.Scanln(&Nl)
	NovaString := strings.ReplaceAll(nome, Al, Nl)
	fmt.Println("String atualizada:", NovaString)
}
