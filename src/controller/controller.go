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

// GetPerfis godoc
// @Summary      Lista todos os perfis
// @Tags         perfis
// @Produce      json
// @Success      200  {array}   model.Perfil
// @Router       /perfis [get]
func (pc *PerfilController) GetPerfis(c *gin.Context) {
	perfis := pc.service.GetAll()
	c.JSON(http.StatusOK, perfis)
}

// CreatePerfil godoc
// @Summary      Cria um novo perfil
// @Tags         perfis
// @Accept       json
// @Produce      json
// @Param        perfil  body      model.Perfil  true  "Dados do perfil"
// @Success      201  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Security     BearerAuth
// @Router       /perfis [post]
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

// UpdatePerfil godoc
// @Summary      Atualiza um perfil existente
// @Tags         perfis
// @Accept       json
// @Produce      json
// @Param        id      path      int           true  "ID do perfil"
// @Param        perfil  body      model.Perfil  true  "Dados atualizados"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Security     BearerAuth
// @Router       /perfis/{id} [put]
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

// DeletePerfil godoc
// @Summary      Remove um perfil
// @Tags         perfis
// @Produce      json
// @Param        id  path      int  true  "ID do perfil"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Security     BearerAuth
// @Router       /perfis/{id} [delete]
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

// GetSessao godoc
// @Summary      Lista todas as sessões
// @Tags         sessoes
// @Produce      json
// @Success      200  {array}   model.Sessao
// @Router       /sessoes [get]
func (pcs *SessaoController) GetSessao(c *gin.Context) {
	sessoes := pcs.service.GetAll()
	c.JSON(http.StatusOK, sessoes)
}

// CreateSessao godoc
// @Summary      Cria uma nova sessão
// @Tags         sessoes
// @Accept       json
// @Produce      json
// @Param        sessao  body      model.Sessao  true  "Dados da sessão"
// @Success      201  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Security     BearerAuth
// @Router       /sessoes [post]
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

// UpdateSessao godoc
// @Summary      Atualiza uma sessão existente
// @Tags         sessoes
// @Accept       json
// @Produce      json
// @Param        id      path      int           true  "ID da sessão"
// @Param        sessao  body      model.Sessao  true  "Dados atualizados"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Security     BearerAuth
// @Router       /sessoes/{id} [put]
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

// DeleteSessao godoc
// @Summary      Remove uma sessão
// @Tags         sessoes
// @Produce      json
// @Param        id  path      int  true  "ID da sessão"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Security     BearerAuth
// @Router       /sessoes/{id} [delete]
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

// CreateMatch godoc
// @Summary      Calcula o match entre um perfil e uma sessão
// @Tags         matches
// @Accept       json
// @Produce      json
// @Param        match  body      model.Match  true  "perfilId e sessaoId"
// @Success      201  {object}  model.Match
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Security     BearerAuth
// @Router       /matches [post]
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
	c.JSON(http.StatusCreated, resultado)
}

// GetMatchesByPerfil godoc
// @Summary      Lista os matches aprovados de um perfil
// @Tags         matches
// @Produce      json
// @Param        id  path      int  true  "ID do perfil"
// @Success      200  {array}   model.Match
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /perfis/{id}/matches [get]
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

// ---------- Auth ----------

type AuthController struct {
	service *service.UsuarioService
}

func NewAuthController(service *service.UsuarioService) *AuthController {
	return &AuthController{service: service}
}

// Register godoc
// @Summary      Cria um novo usuário
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        usuario  body      model.LoginRequest  true  "Email e senha"
// @Success      201  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Router       /register [post]
func (ac *AuthController) Register(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	usuario, err := ac.service.Register(req.Email, req.Senha)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"mensagem": "Usuário criado!", "id": usuario.ID})
}

// Login godoc
// @Summary      Autentica um usuário e retorna o token JWT
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credenciais  body      model.LoginRequest   true  "Email e senha"
// @Success      200          {object}  model.LoginResponse
// @Failure      400          {object}  map[string]string
// @Failure      401          {object}  map[string]string
// @Router       /login [post]
func (ac *AuthController) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	token, err := ac.service.Login(req.Email, req.Senha)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.LoginResponse{Token: token})
}
