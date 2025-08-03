package main

import (
	"fmt"
	"goSre/database"
)

func main() {
	// Cria conexão com o banco
	db := database.UserDatabase() // deve retornar *gorm.DB

	// Usa a interface em vez da implementação concreta
	var userRepo database.UserInterface = database.NewUserRepository(db)

	// Criar usuário
	user, err := userRepo.CreateUser("teste@email.com", "123456")
	if err != nil {
		fmt.Println("Erro ao criar usuário:", err)
		return
	}
	fmt.Println("Usuário criado:", user)

	// Atualizar senha
	updatedUser, err := userRepo.UpdatePassword(user.Email, "novaSenha")
	if err != nil {
		fmt.Println("Erro ao atualizar senha:", err)
		return
	}
	fmt.Println("Usuário atualizado:", updatedUser)

	// Deletar usuário
	deletedUser, err := userRepo.DeleteUser(user.Email)
	if err != nil {
		fmt.Println("Erro ao deletar usuário:", err)
		return
	}
	fmt.Println("Usuário deletado:", deletedUser)
}
