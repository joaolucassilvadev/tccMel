package entites

type Administrador struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"Password"`
}
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Engresso struct {
	ID             uint
	Name           string `json:"name"`
	Cpf            string `json:"Cpf"gorm:"unique"`
	DataNascimento string `json:"data_nascimento"`
}

// é sempre bom fazer a regra de negocios dentro da minha entidade e não decidir o banco de dados que mande na regra de negocios
// Bom sempre que desejarem criar uma category nova eles vão utilizar nossa função não nosso type
func NewEngresso(name, cpf, dataNascimento string) (*Engresso, error) {
	category := &Engresso{
		Name:           name,
		DataNascimento: dataNascimento,
		Cpf:            cpf,
	}

	return category, nil
}

func NewAdmistrador(name, email, senha string) (*Administrador, error) {
	categoryAdmistraodr := &Administrador{
		Name:     name,
		Email:    email,
		Password: senha,
	}

	return categoryAdmistraodr, nil
}
