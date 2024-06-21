package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"microservicos.com/cmd/config"
	entites "microservicos.com/internal/entities"
)

type FormularioDTO struct {
	Genero                    entites.Genero `json:"genero"`
	PaisResidencia            string         `json:"pais_residencia"`
	CidadeResidencia          string         `json:"cidade_residencia"`
	EstadoResidencia          string         `json:"estado_residencia"`
	GraduacaoRealizada        int            `json:"graduacao_realizada"`
	AnoConclusaoGraduacao     string         `json:"ano_conclusao_graduacao"`
	LinhaPesquisaPosGraduacao int            `json:"linha_pesquisa_pos_graduacao"`
	AnoInicioPosGraduacao     string         `json:"ano_inicio_pos_graduacao"`
	AnoConclusaoPosGraduacao  string         `json:"ano_conclusao_pos_graduacao"`
	TrabalhandoAtualmente     bool           `json:"trabalhando_atualmente"`
	InstituicaoTrabalho       string         `json:"instituicao_trabalho"`
	Ramo                      string         `json:"ramo"`
	FaixaSalarial             string         `json:"faixa_salarial"`
	AtuacaoDocencia           bool           `json:"atuacao_docencia"`
	EngressoID                uint           `json:"engresso_id"`
	DateInicio                time.Time      `json:"date_inicio"`
	DateFim                   time.Time      `json:"date_fim"`
}

// CreateFormularioController godoc
// @Summary Criar um novo formulário
// @Description Cria um novo formulário com base nos dados fornecidos do engresso.
// @Tags formulario
// @Accept json
// @Produce json
// @Param formulario body entites.Formulario true "Dados do formulário a ser criado"
// @Router /formulario [post]
func CreateFormularioController(ctx *gin.Context) {
	var formDTO FormularioDTO

	// Bind JSON request body to DTO
	if err := ctx.ShouldBindJSON(&formDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Por favor, forneça dados válidos"})
		return
	}

	// Buscar o Engresso associado
	var engresso entites.Engresso
	if err := config.DB.First(&engresso, formDTO.EngressoID).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Engresso não encontrado"})
		return
	}

	// Criar entidade de formulário
	formEntity := entites.Formulario{
		Genero:                    formDTO.Genero,
		PaisResidencia:            formDTO.PaisResidencia,
		CidadeResidencia:          formDTO.CidadeResidencia,
		EstadoResidencia:          formDTO.EstadoResidencia,
		GraduacaoRealizada:        formDTO.GraduacaoRealizada,
		AnoConclusaoGraduacao:     formDTO.AnoConclusaoGraduacao,
		LinhaPesquisaPosGraduacao: formDTO.LinhaPesquisaPosGraduacao,
		AnoInicioPosGraduacao:     formDTO.AnoInicioPosGraduacao,
		AnoConclusaoPosGraduacao:  formDTO.AnoConclusaoPosGraduacao,
		TrabalhandoAtualmente:     formDTO.TrabalhandoAtualmente,
		InstituicaoTrabalho:       formDTO.InstituicaoTrabalho,
		Ramo:                      formDTO.Ramo,
		FaixaSalarial:             formDTO.FaixaSalarial,
		AtuacaoDocencia:           formDTO.AtuacaoDocencia,

		EngressoID: formDTO.EngressoID,
		Engresso:   engresso,
	}

	// Salvar entidade no banco de dados
	if err := config.DB.Create(&formEntity).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao salvar no banco de dados", "details": err.Error()})
		return
	}

	// Responder com sucesso
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Formulário criado com sucesso",

		"nome_engresso": engresso.Name,
		"formulario":    formEntity,
		"engresso":      engresso,
	})
}
