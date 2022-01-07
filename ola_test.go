package main

import "testing"

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("Diz olá para as pessoas", func(t *testing.T) {

		resultado := Ola("Chris", "Hello")
		esperado := "Olá, Chris"
		verificaMensagemCorreta(t, resultado, esperado)

	})

	t.Run("Diz 'Olá, mundo' quando entra uma string vazia", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, mundo"
		verificaMensagemCorreta(t, resultado, esperado)

	})

	t.Run("Em Espanhol", func(t *testing.T) {
		resultado := Ola("Elodie", "espanhol")
		esperado := "Hola, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("Em Frances", func(t *testing.T) {
		resultado := Ola("Elodie", "frances")
		esperado := "Bonjour, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})

}
