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
func CreateUser(email string, senha string) *entity.User {
	db := UserDatabase()
	NewUser := entity.User{
		Email: email,
		Senha: senha,
	}
	db.Create(&NewUser)
	return &NewUser
}
func DeleteUser(email string) *entity.User {
	db := UserDatabase()                                  //conecto ao banco de dados
	var user entity.User                                  // crio como se fosse um construtor para acesssar a struck User
	userFind := db.Where("email = ?", email).First(&user) // faço a busca com o where para achar o email

	if userFind.Error != nil {
		panic("usuário não encontrado")
	}
	db.Delete(&user) // comando para deletar o usuário
	return &user     // retorno o usuário deletado
}
func UpdatePassword(email string, senha string) *entity.User {
	db := UserDatabase()
	var user entity.User

	userFind := db.Where("email = ?", email).First(&user)
	if userFind.Error != nil {
		panic("usuário não encontrado")
	}
	db.Model(&user).Update("senha", senha) // model server para falar que eu quero a tabela USER
	// e o UPDATE na coluna senha

	user.Senha = senha // atualizando com a nova senha

	return &user
}
