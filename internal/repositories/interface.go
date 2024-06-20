package repositories

import entites "microservicos.com/internal/entities"

// bom aqui vamos definir os metodos que nosso repositorio deve ter
type IRepositories interface {
	Save(category *entites.Engresso) error
	// Savee(categoryAdmistraodr *entites.Administrador) error
}

type IRepositoriess interface {
	Savee(categoryAdmistraodr *entites.Administrador) error
}
