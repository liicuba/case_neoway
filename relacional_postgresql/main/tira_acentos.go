package main

import (
	"os"
	"golang.org/x/text/unicode/norm"
	"unicode"
	"fmt"
)

// Função para remover acentos
func RemoveAcentos(input string) string {
	// Normaliza a string para decompor os caracteres acentuados
	// e depois remove os caracteres de acento.
	result := norm.NFD.String(input)
	var output []rune
	for _, r := range result {
		if unicode.Is(unicode.M, r) { // Ignore caracteres de acento
			continue
		}
		output = append(output, r)
	}
	return string(output)
}

// Função que processa e remove acentos de todos os campos
func RemoveAcentosDosCampos() []string {
	// Obtemos os campos através da função PegaCampos()
	campos := PegaCampos()
	var camposSemAcento []string

	// Para cada campo, removemos os acentos
	for _, campo := range campos {
		camposSemAcento = append(camposSemAcento, RemoveAcentos(campo))
	}

	return camposSemAcento
}
