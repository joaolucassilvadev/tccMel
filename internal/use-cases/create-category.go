package usecases

import (
	"log"

	entites "microservicos.com/internal/entities"
	"microservicos.com/internal/repositories"
)

type createCategoryUseCase struct {
	// Db aqui recebemos uma instancia do banco de dados

	repository repositories.IRepositories
}

func NewcreateCategoryUseCase(repository repositories.IRepositories) *createCategoryUseCase {
	return &createCategoryUseCase{repository}
}

func (u *createCategoryUseCase) Execute(name, cpf, dataNascimento string) error {
	category, err := entites.NewEngresso(name, cpf, dataNascimento)
	if err != nil {
		return err
	}

	log.Println(category)
	// bom aqui quando botamos esse u.repository ele não sabe com qual banco de dados ele está lidando
	// não importa qual é o banco ou se é nosql ou sql, pois ele só tem o trabalho de salvar
	err = u.repository.Save(category)
	if err != nil {
		return err
	}
	return nil
}
