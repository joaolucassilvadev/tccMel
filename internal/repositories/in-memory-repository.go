package repositories

import (
	"gorm.io/gorm"
	"microservicos.com/cmd/config"
	entites "microservicos.com/internal/entities"
)

type CategoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository cria uma nova instância do repositório usando a conexão GORM
func NewCategoryRepository() *CategoryRepository {
	db := config.GetDatabase()
	db.AutoMigrate(&entites.Engresso{}) // Supondo que Engresso seja sua entidade de ingresso
	return &CategoryRepository{
		db: db,
	}
}

// NewAdministradorRepository cria uma nova instância do repositório de administradores usando a conexão GORM
func NewAdministradorRepository() *CategoryRepository {
	db := config.GetDatabase()
	db.AutoMigrate(&entites.Administrador{})
	return &CategoryRepository{
		db: db,
	}
}

// Save salva um novo Engresso no banco de dados
func (r *CategoryRepository) Save(category *entites.Engresso) error {
	return r.db.Create(category).Error
}

// Savee salva um novo Administrador no banco de dados
func (r *CategoryRepository) Savee(administrador *entites.Administrador) error {
	return r.db.Create(administrador).Error
}
