package database

import (
	"errors"
	"goSre/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(email string, senha string) (*entity.User, error) {
	NewUser := entity.User{
		Email: email,
		Senha: senha,
	}
	if err := r.DB.Create(&NewUser).Error; err != nil {
		return nil, err
	}
	return &NewUser, nil
}
func (r *UserRepository) DeleteUser(email string) (*entity.User, error) {
	//conecto ao banco de dados
	var user entity.User // crio como se fosse um construtor para acesssar a struck User
	userFind := r.DB.Where("email = ?", email).First(&user)
	if userFind.Error != nil {
		return nil, errors.New("usuário não encontrado")
	} // faço a busca com o where para achar o email

	if err := r.DB.Delete(&user).Error; err != nil {
		return nil, err
	} // comando para deletar o usuário
	return &user, nil // retorno o usuário deletado
}
func (r *UserRepository) UpdatePassword(email string, senha string) (*entity.User, error) {
	var user entity.User

	userFind := r.DB.Where("email = ?", email).First(&user)
	if userFind.Error != nil {
		return nil, errors.New("usuário não encontrado")
	}
	if err := r.DB.Model(&user).Update("senha", senha).Error; err != nil {
		return nil, err
	} // model server para falar que eu quero a tabela USER
	// e o UPDATE na coluna senha

	user.Senha = senha // atualizando com a nova senha

	return &user, nil
}
