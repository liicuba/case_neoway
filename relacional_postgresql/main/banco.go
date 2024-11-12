package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Função para salvar dados no banco de dados PostgreSQL
func salvaBanco(listaUsuario []Usuario, listaEmpresa []Empresa) error {
	// String de conexão para o PostgreSQL
	dsn := "host=localhost user=gorm password=gorm dbname=aline port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("falha ao conectar ao banco de dados: %v", err)
	}

	// Migra o esquema para as tabelas Usuario e Empresa
	db.AutoMigrate(&Usuario{})
	db.AutoMigrate(&Empresa{})

	// Salva os dados nas tabelas
	for _, usuario := range listaUsuario {
		db.Create(&usuario)
	}
	for _, empresa := range listaEmpresa {
		db.Create(&empresa)
	}

	return nil
}
