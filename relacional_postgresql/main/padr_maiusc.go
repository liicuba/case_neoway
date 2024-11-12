
package main

import (
	"strings"
)

// Padroniza os títulos das colunas para maiúsculas
func padr_maiusc(dados []map[string]string) []map[string]string {
	for i := range dados {
		for chave, valor := range dados[i] {
			// Transforma a chave em maiúsculas
			novaChave := strings.ToUpper(chave)
			// Atualiza o mapa com a nova chave e valor
			dados[i][novaChave] = valor
			// Remove a chave antiga
			if novaChave != chave {
				delete(dados[i], chave)
			}
		}
	}
	return dados
}
