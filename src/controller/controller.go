package controller

import (
	"net/http"
	"strconv"

	"match-dos-estudos/src/model"
	"match-dos-estudos/src/service"

	"github.com/gin-gonic/gin"
)

// ---------- Perfil ----------

type PerfilController struct {
	service *service.PerfilService
}

func NewPerfilController(service *service.PerfilService) *PerfilController {
	return &PerfilController{service: service}
}

func (pc *PerfilController) GetPerfis(c *gin.Context) {
	perfis := pc.service.GetAll()
	c.JSON(http.StatusOK, perfis)
}

func (pc *PerfilController) CreatePerfil(c *gin.Context) {
	var perfil model.Perfil
	if err := c.ShouldBindJSON(&perfil); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	resultado := pc.service.Create(perfil)
	if resultado.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Não foi possível criar esse perfil!"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"mensagem": "Perfil criado com sucesso!", "perfil": resultado})
}

func (pc *PerfilController) UpdatePerfil(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	var perfil model.Perfil
	if err := c.ShouldBindJSON(&perfil); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	resultado := pc.service.Update(id, perfil)
	if resultado.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Perfil não encontrado"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensagem": "Perfil atualizado com sucesso!"})
}

func (pc *PerfilController) DeletePerfil(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	resultado := pc.service.Delete(id)
	if len(resultado) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Perfil não encontrado"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensagem": "Perfil removido com sucesso!"})
}

// ---------- Sessao ----------

type SessaoController struct {
	service *service.SessaoService
}

func NewSessaoController(service *service.SessaoService) *SessaoController {
	return &SessaoController{service: service}
}

func (pcs *SessaoController) GetSessao(c *gin.Context) {
	sessoes := pcs.service.GetAll()
	c.JSON(http.StatusOK, sessoes)
}

func (pcs *SessaoController) CreateSessao(c *gin.Context) {
	var sessao model.Sessao
	if err := c.ShouldBindJSON(&sessao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	resultado := pcs.service.Create(sessao)
	if resultado.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Não foi possível criar essa sessão!"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"mensagem": "Sessão criada com sucesso!", "sessao": resultado})
}

func (pcs *SessaoController) UpdateSessao(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	var sessao model.Sessao
	if err := c.ShouldBindJSON(&sessao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	resultado := pcs.service.Update(id, sessao)
	if resultado.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Sessão não encontrada"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensagem": "Sessão atualizada com sucesso!"})
}

func (pcs *SessaoController) DeleteSessao(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	resultado := pcs.service.Delete(id)
	if len(resultado) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Sessão não encontrada"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensagem": "Sessão removida com sucesso!"})
}

// ---------- Match ----------

type MatchController struct {
	service *service.MatchService
}

func NewMatchController(service *service.MatchService) *MatchController {
	return &MatchController{service: service}
}

// POST /matches
// Body: { "perfilId": 1, "sessaoId": 2 }
// Calcula o score entre o perfil e a sessão e retorna se o match foi aprovado.
func (mc *MatchController) CreateMatch(c *gin.Context) {
	var match model.Match
	if err := c.ShouldBindJSON(&match); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	resultado, err := mc.service.Create(match)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	status := http.StatusCreated
	c.JSON(status, resultado)
}

// GET /perfis/:id/matches
// Retorna todos os matches aprovados do perfil com o id informado.
func (mc *MatchController) GetMatchesByPerfil(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	matches, err := mc.service.GetByPerfilID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matches)
}
