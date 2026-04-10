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

	repoPerfil := repository.NewPerfilRepository()
	servicePerfil := service.NewPerfilService(repoPerfil)
	controllerPerfin := controller.NewPerfilController(servicePerfil)

	repoSessao := repository.NewSessaoRepository()
	serviceSessao := service.NewSessaoService(repoSessao)
	controllerSessao := controller.NewSessaoController(serviceSessao)

	repoMatch := repository.NewMatchRepository()
	serviceMatch := service.NewMatchService(repoMatch)
	controllerMatch := controller.NewMatchController(serviceMatch)

	//PERFIL
	r.GET("/perfis", controllerPerfin.GetPerfis)
	r.POST("/perfis", controllerPerfin.CreatePerfil)
	r.PUT("/perfis/:id", controllerPerfin.UpdatePerfil)
	r.DELETE("/perfis/:id", controllerPerfin.DeletePerfil)

	// //SESSOES
	r.GET("/sessoes", controllerSessao.GetSessao)
	r.POST("/sessoes", controllerSessao.CreateSessao)
	r.PUT("/sessoes/:id", controllerSessao.UpdateSessao)
	r.DELETE("/sessoes/:id", controllerSessao.DeleteSessao)

	// //MATCH
	r.POST("/matches", controllerMatch.CreateMatch)
	r.GET("perfis/:id/matches", controllerMatch.GetMatch)

	return r
}
