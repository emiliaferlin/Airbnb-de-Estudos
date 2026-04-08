package controller

import (
	"net/http"

	"match-dos-estudos/src/service"

	"github.com/gin-gonic/gin"
)

type PerfilController struct {
	service *service.PerfilService
}

func NewPerfilController() *PerfilController {
	return &PerfilController{
		service: service.NewPerfilService(),
	}
}

func (pc *PerfilController) GetPerfis(c *gin.Context) {
	perfis := pc.service.GetAll()
	c.JSON(http.StatusOK, perfis)
}
