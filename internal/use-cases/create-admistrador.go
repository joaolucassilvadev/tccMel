package usecases

import (
	"log"

	entites "microservicos.com/internal/entities"
	"microservicos.com/internal/repositories"
)

type createAdmistrador struct {
	// Db aqui recebemos uma instancia do banco de dados

	repositoryy repositories.IRepositoriess
}

func NewcreateAdmistradorUseCase(repository repositories.IRepositoriess) *createAdmistrador {
	return &createAdmistrador{repository}
}

func (u *createAdmistrador) Execute(name, email, senha string) error {
	categoryAdmistraodr, err := entites.NewAdmistrador(name, email, senha)
	if err != nil {
		return err
	}

	log.Println(categoryAdmistraodr)

	err = u.repositoryy.Savee(categoryAdmistraodr)
	if err != nil {
		return err
	}
	return nil
}
