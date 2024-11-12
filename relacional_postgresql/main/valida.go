package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Função `valida` para verificar se o número de identificação é um CPF ou CNPJ válido
func valida(numeroIdentificador string) string {
	switch len(numeroIdentificador) {
	case 11: // Verifica se o número possui 11 dígitos (CPF)
		if validaCPF(numeroIdentificador) {
			return "CPF válido"
		}
		return "CPF inválido"
	case 14: // Verifica se o número possui 14 dígitos (CNPJ)
		if validaCNPJ(numeroIdentificador) {
			return "CNPJ válido"
		}
		return "CNPJ inválido"
	default:
		return validaNumeroForaDoPadrao(numeroIdentificador)
	}
}

// Função para lidar com números que não estão no padrão esperado de CPF ou CNPJ
func validaNumeroForaDoPadrao(numero string) string {
	return "Número fora do padrão CPF/CNPJ"
}

// Função para validar CPF
func validaCPF(cpf string) bool {
	// Remove caracteres especiais como ponto e hífen
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	// Verifica se o CPF possui 11 dígitos e não é composto por dígitos repetidos
	if len(cpf) != 11 || strings.Count(cpf, string(cpf[0])) == 11 {
		return false
	}

	// Calcula o primeiro dígito verificador
	var sum int
	for i := 0; i < 9; i++ {
		num, _ := strconv.Atoi(string(cpf[i]))
		sum += num * (10 - i)
	}
	digit1 := (sum * 10 % 11) % 10

	// Calcula o segundo dígito verificador
	sum = 0
	for i := 0; i < 10; i++ {
		num, _ := strconv.Atoi(string(cpf[i]))
		sum += num * (11 - i)
	}
	digit2 := (sum * 10 % 11) % 10

	// Verifica se os dígitos calculados correspondem aos dígitos no CPF
	return digit1 == int(cpf[9]-'0') && digit2 == int(cpf[10]-'0')
}

// Função para validar CNPJ
func validaCNPJ(cnpj string) bool {
	// Remove caracteres especiais como ponto, hífen e barra
	cnpj = strings.ReplaceAll(cnpj, ".", "")
	cnpj = strings.ReplaceAll(cnpj, "-", "")
	cnpj = strings.ReplaceAll(cnpj, "/", "")

	// Verifica se o CNPJ possui 14 dígitos e não é composto por dígitos repetidos
	if len(cnpj) != 14 || strings.Count(cnpj, string(cnpj[0])) == 14 {
		return false
	}

	// Define os pesos para cálculo dos dígitos verificadores
	weights1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	weights2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	// Calcula o primeiro dígito verificador
	var sum int
	for i := 0; i < 12; i++ {
		num, _ := strconv.Atoi(string(cnpj[i]))
		sum += num * weights1[i]
	}
	digit1 := (sum % 11)
	if digit1 < 2 {
		digit1 = 0
	} else {
		digit1 = 11 - digit1
	}

	// Calcula o segundo dígito verificador
	sum = 0
	for i := 0; i < 13; i++ {
		num, _ := strconv.Atoi(string(cnpj[i]))
		sum += num * weights2[i]
	}
	digit2 := (sum % 11)
	if digit2 < 2 {
		digit2 = 0
	} else {
		digit2 = 11 - digit2
	}

	// Verifica se os dígitos calculados correspondem aos dígitos no CNPJ
	return digit1 == int(cnpj[13]-'0') && digit2 == int(cnpj[14]-'0')
}

func main() {
	// Configurando o servidor Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Definindo rota para validação
	e.GET("/valida/:numero", func(c echo.Context) error {
		numero := c.Param("numero")
		resultado := valida(numero)
		return c.String(http.StatusOK, resultado)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
