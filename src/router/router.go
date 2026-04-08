package router

import (
	"match-dos-estudos/src/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	perfilController := controller.NewPerfilController()

	//PERFIL
	r.GET("/perfis", perfilController.GetPerfis)
	r.POST("/perfis", perfilController.GetPerfis)
	r.PUT("/perfis/:id", perfilController.GetPerfis)
	r.DELETE("/perfis/:id", perfilController.GetPerfis)

	//SESSOES
	r.GET("/perfis", perfilController.GetPerfis)
	r.POST("/perfis", perfilController.GetPerfis)
	r.PUT("/perfis/:id", perfilController.GetPerfis)
	r.DELETE("/perfis/:id", perfilController.GetPerfis)

	//MATCH
	r.GET("/matches", perfilController.GetPerfis)

	return r
}
