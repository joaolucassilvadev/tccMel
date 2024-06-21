package entites

import (
	"time"
)

type LoginSuccessResponse struct {
	Mensagem string `json:"mensagem"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Administrador struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"Password"`
}
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Loginengresso struct {
	Cpf            string `json:"cpf" gorm:"unique"`
	DataNascimento string `json:"data_nascimento"`
}
type Engresso struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `json:"name"`
	Cpf            string `json:"cpf" gorm:"unique"`
	DataNascimento string `json:"data_nascimento"`
}

type DateFormulario struct {
	ID         uint   `gorm:"primaryKey"`
	DateInicio string `json:"date_inicio"`
	DateFim    string `json:"date_fim"`
	create     time.Time
}

type Genero string

const (
	Feminino            Genero = "Feminino"
	Masculino           Genero = "Masculino"
	NAO_BINARIO         Genero = "Não-Binário"
	NAO_DESEJO_INFORMAR Genero = "Não Desejo Informar"
)

type Formulario struct {
	ID                        uint   `gorm:"primaryKey"`
	Genero                    Genero `json:"genero"`
	PaisResidencia            string `json:"pais_residencia"`
	CidadeResidencia          string `json:"cidade_residencia"`
	EstadoResidencia          string `json:"estado_residencia"`
	GraduacaoRealizada        int    `json:"graduacao_realizada"`
	AnoConclusaoGraduacao     string `json:"ano_conclusao_graduacao"`
	LinhaPesquisaPosGraduacao int    `json:"linha_pesquisa_pos_graduacao"`
	AnoInicioPosGraduacao     string `json:"ano_inicio_pos_graduacao"`
	AnoConclusaoPosGraduacao  string `json:"ano_conclusao_pos_graduacao"`
	TrabalhandoAtualmente     bool   `json:"trabalhando_atualmente"`
	InstituicaoTrabalho       string `json:"instituicao_trabalho"`
	Ramo                      string `json:"ramo"`
	FaixaSalarial             string `json:"faixa_salarial"`
	AtuacaoDocencia           bool   `json:"atuacao_docencia"`
	EngressoID                uint   `json:"engresso_id"`
	Engresso                  Engresso
}

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
