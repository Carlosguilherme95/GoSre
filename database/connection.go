package database

import (
	"fmt"
	"goSre/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func UserDatabase() *gorm.DB {

	dsn := "root:root123@tcp(127.0.0.1:3306)/meu_banco?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("erro ao conectar ao banco de dados", err))
	}
	err = db.AutoMigrate(&entity.User{})

	if err != nil {
		panic("erro ao criar a tabela")
	}

	fmt.Println("Conectado ao banco de dados")
	return db
}
