package database

import (
    "desafioGo/model"
    "github.com/glebarez/sqlite"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func Init() {
	var err error

	DB, err = gorm.Open(sqlite.Open("fruits.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar no banco de dados: ", err)
	}

	err = DB.AutoMigrate(&model.Fruit{})
	if err != nil {
		log.Fatal("Erro ao migrar o modelo Fruit: ", err)
	}

	log.Println("Conexão com o banco realizada com sucesso e migração executada.")
}
