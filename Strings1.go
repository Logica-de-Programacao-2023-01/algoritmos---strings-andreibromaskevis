package main

import (
	"fmt"
	"strings"
)
func main() {
	var nome string
	fmt.Print("Digite um nome: ")
	fmt.Scanln(&nome)
	LM := strings.ToUpper(nome)
	fmt.Println("String convertida para maiúsculas:", LM)
}
