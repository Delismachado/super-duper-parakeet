package main

import "fmt"

const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const espanhol = "espanhol"
const frances = "frances"
const prefixoOlaFrances = "Bonjour, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}

	if idioma == espanhol {
		return prefixoOlaEspanhol + nome
	}

	if idioma == frances {
		return prefixoOlaFrances + nome
	}

	return prefixoOlaPortugues + nome
}
func main() {
	fmt.Println(Ola("Olá", "Hola"))
}
