package router

import (
	"match-dos-estudos/src/controller"
	"match-dos-estudos/src/repository"
	"match-dos-estudos/src/service"

	"github.com/gin-gonic/gin"
)

//Router → Controller → Service → DAO (ou Repository)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	repo := repository.NewPerfilRepository()
	service := service.NewPerfilService(repo)
	controller := controller.NewPerfilController(service)

	//PERFIL
	r.GET("/perfis", controller.GetPerfis)
	r.POST("/perfis", controller.CreatePerfil)
	r.PUT("/perfis/:id", controller.UpdatePerfil)
	r.DELETE("/perfis/:id", controller.DeletePerfil)

	// //SESSOES
	// r.GET("/perfis", controller.GetPerfis)
	// r.POST("/perfis", controller.GetPerfis)
	// r.PUT("/perfis/:id", controller.GetPerfis)
	// r.DELETE("/perfis/:id", controller.GetPerfis)

	// //MATCH
	// r.GET("/matches", controller.GetPerfis)

	return r
}
