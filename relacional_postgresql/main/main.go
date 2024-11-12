package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Função de upload para lidar com o envio do arquivo
func upload(c echo.Context) error {
	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	arquivo, err := file.Open()
	if err != nil {
		return err
	}
	defer arquivo.Close()

	// Chamada para as funções em ordem, passando o arquivo ou dados relevantes entre elas
	// Testa separadores no arquivo
	linhasSeparadas := testa_separador(arquivo)

	// Testa se algum campo é nulo ou vazio (ex.: CPF)
	usuariosSemNulos := testa_nulos(linhasSeparadas)

	// Testa duplicados nos campos
	usuariosUnicos := testa_duplicados(usuariosSemNulos)

	// Remove acentos nos títulos das colunas
	usuariosSemAcentos := tira_acentos(usuariosUnicos)

	// Padroniza os títulos das colunas para maiúsculas
	usuariosComMaiusculas := padr_maiusc(usuariosSemAcentos)

	// Valida CPF ou CNPJ
	usuariosValidos := valida(usuariosComMaiusculas)

	// Chama a função banco para salvar os dados no banco de dados PostgreSQL
	err = banco(usuariosValidos) // Substituindo a chamada do Salva_Postgres
	if err != nil {
		return fmt.Errorf("erro ao salvar no banco de dados: %v", err)
	}

}

// Função de rota para a página principal
func home(c echo.Context) error {
	return c.HTML(http.StatusOK, "<p>Welcome to the upload service</p>")
}

func main() {
	e := echo.New()

	// Trata automaticamente panics (erros críticos que causariam a interrupção do servidor)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Rota para a página principal
	e.GET("/", home)

	// Rotas para outras tabelas e upload
	e.GET("/usuario", usuario) // Rota para a tabela usuario
	e.GET("/empresa", empresa) // Rota para a tabela empresa
	e.POST("/upload", upload)  // Rota de upload de arquivo

	// Inicia o servidor na porta 1323
	e.Logger.Fatal(e.Start(":1323"))
}
