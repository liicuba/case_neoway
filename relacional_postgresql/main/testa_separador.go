package main

import (
	"fmt"
	"strings"
)

var separadores = []string{" ", ";", ","}

// Estruturas para representar os dados de Usuario e Empresa
type Usuario struct {
	gorm.Model
	CPF                 string
	Private             int
	Incompleto          int
	UltimaCompra        string
	Compra_Ticket_Medio string
	Ticket_Ultima_Compra string
	Loja_Mais_Frequente string
	Loja_Ultima_Compra  string
}

type Empresa struct {
	gorm.Model
	CNPJ                string
	Private             int
	Incompleto          int
	UltimaCompra        string
	Compra_Ticket_Medio string
	Ticket_Ultima_Compra string
	Loja_Mais_Frequente string
	Loja_Ultima_Compra  string
}

// Função pegaCampos para processar o arquivo
func pegaCampos(arquivo *File) ([]Usuario, []Empresa, []string) {
	var listaUsuarios []Usuario
	var listaEmpresas []Empresa
	var listaErrado []string

	for _, linha := range arquivo {
		for _, separador := range separadores {
			if strings.Contains(linha, separador) {
				linhaSeparada := strings.Split(linha, separador)

				numeroID := linhaSeparada[0]
				if valida(numeroID) == "cpf" {
					usuario := Usuario{
						CPF:                 numeroID,
						Private:             1,
						Incompleto:          2,
						UltimaCompra:        "3",
						Compra_Ticket_Medio: "4",
						Ticket_Ultima_Compra: "5",
						Loja_Mais_Frequente: "6",
						Loja_Ultima_Compra:  "7",
					}
					listaUsuarios = append(listaUsuarios, usuario)
				} else if valida(numeroID) == "cnpj" {
					empresa := Empresa{
						CNPJ:                numeroID,
						Private:             1,
						Incompleto:          2,
						UltimaCompra:        "3",
						Compra_Ticket_Medio: "4",
						Ticket_Ultima_Compra: "5",
						Loja_Mais_Frequente: "6",
						Loja_Ultima_Compra:  "7",
					}
					listaEmpresas = append(listaEmpresas, empresa)
				} else {
					listaErrado = append(listaErrado, numeroID)
				}
			}
		}
	}
	return listaUsuarios, listaEmpresas, listaErrado
}

func main() {
	// Exemplo de arquivo sendo passado para `pegaCampos`
	usuarios, empresas, _ := pegaCampos(arquivo)

	// Chama `salvaBanco` para salvar os dados no banco de dados PostgreSQL
	err := salvaBanco(usuarios, empresas)
	if err != nil {
		fmt.Printf("Erro ao salvar no banco de dados: %v\n", err)
	} else {
		fmt.Println("Dados salvos com sucesso no banco de dados.")
	}
}
