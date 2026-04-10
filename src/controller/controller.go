package controller

import (
	"net/http"
	"strconv"

	"match-dos-estudos/src/model"
	"match-dos-estudos/src/service"

	"github.com/gin-gonic/gin"
)

type PerfilController struct {
	service *service.PerfilService
}

func NewPerfilController(service *service.PerfilService) *PerfilController {
	return &PerfilController{
		service: service,
	}
}

func (pc *PerfilController) GetPerfis(c *gin.Context) {
	perfis := pc.service.GetAll()
	c.JSON(http.StatusOK, perfis)
}

func (pc *PerfilController) CreatePerfil(c *gin.Context) {
	var perfil model.Perfil

	if err := c.ShouldBindJSON(&perfil); err != nil {
		c.JSON(400, gin.H{"erro": err.Error()})
		return
	}

	resultado := pc.service.Create(perfil)
	if resultado.ID == 0 {
		c.JSON(400, gin.H{"erro": "Não foi possível criar esse perfil!"})
		return
	}

	c.JSON(201, "Perfil criado com sucesso!")
}

func (pc *PerfilController) UpdatePerfil(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"erro": "ID inválido"})
		return
	}

	var perfil model.Perfil

	if err := c.ShouldBindJSON(&perfil); err != nil {
		c.JSON(400, gin.H{"erro": err.Error()})
		return
	}

	resultado := pc.service.Update(id, perfil)
	if resultado.ID == 0 {
		c.JSON(400, gin.H{"erro": "Código não foi encontrado"})
		return
	}

	c.JSON(200, "Perfil atualizado com sucesso!")
}

func (pc *PerfilController) DeletePerfil(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"erro": "ID inválido"})
		return
	}

	resultado := pc.service.Delete(id)
	if len(resultado) == 0 {
		c.JSON(400, gin.H{"erro": "Código não foi encontrado"})
		return
	}

	c.JSON(200, "Perfil removido com sucesso!")
}

type SessaoController struct {
	service *service.SessaoService
}

func NewSessaoController(service *service.SessaoService) *SessaoController {
	return &SessaoController{
		service: service,
	}
}

func (pcs *SessaoController) GetSessao(c *gin.Context) {
	sessoes := pcs.service.GetAll()
	c.JSON(http.StatusOK, sessoes)
}

func (pcs *SessaoController) CreateSessao(c *gin.Context) {
	var sessao model.Sessao

	if err := c.ShouldBindJSON(&sessao); err != nil {
		c.JSON(400, gin.H{"erro": err.Error()})
		return
	}

	resultado := pcs.service.Create(sessao)
	if resultado.ID == 0 {
		c.JSON(400, gin.H{"erro": "Não foi possível criar esse Sessão!"})
		return
	}

	c.JSON(201, "Sessão criado com sucesso!")
}

func (pcs *SessaoController) UpdateSessao(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"erro": "ID inválido"})
		return
	}

	var sessao model.Sessao

	if err := c.ShouldBindJSON(&sessao); err != nil {
		c.JSON(400, gin.H{"erro": err.Error()})
		return
	}

	resultado := pcs.service.Update(id, sessao)
	if resultado.ID == 0 {
		c.JSON(400, gin.H{"erro": "Código não foi encontrado"})
		return
	}

	c.JSON(200, "Sessão atualizado com sucesso!")
}

func (pcs *SessaoController) DeleteSessao(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"erro": "ID inválido"})
		return
	}

	resultado := pcs.service.Delete(id)
	if len(resultado) == 0 {
		c.JSON(400, gin.H{"erro": "Código não foi encontrado"})
		return
	}

	c.JSON(200, "Sessao removido com sucesso!")
}
