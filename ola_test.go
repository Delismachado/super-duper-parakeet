package main

import "testing"

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	}

	t.Run("Diz olá para as pessoas", func(t *testing.T) {

		result := Ola("Chris")
		expected := "Olá, Chris"
		verificaMensagemCorreta(t, result, expected)

	})

	t.Run("Diz 'Olá, mundo' quando entra uma string vazia", func(t *testing.T) {
		result := Ola("")
		expected := "Olá, mundo"
		verificaMensagemCorreta(t, result, expected)

	})

}
