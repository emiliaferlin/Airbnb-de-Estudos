package router

import (
	"match-dos-estudos/src/controller"
	"match-dos-estudos/src/repository"
	"match-dos-estudos/src/service"

	"github.com/gin-gonic/gin"
)

// Router → Controller → Service → Repository

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Repositórios
	repoPerfil := repository.NewPerfilRepository()
	repoSessao := repository.NewSessaoRepository()
	repoMatch := repository.NewMatchRepository()

	// Services
	servicePerfil := service.NewPerfilService(repoPerfil)
	serviceSessao := service.NewSessaoService(repoSessao)
	// MatchService recebe os três repos para calcular o score internamente
	serviceMatch := service.NewMatchService(repoMatch, repoPerfil, repoSessao)

	// Controllers
	controllerPerfil := controller.NewPerfilController(servicePerfil)
	controllerSessao := controller.NewSessaoController(serviceSessao)
	controllerMatch := controller.NewMatchController(serviceMatch)

	// PERFIL
	r.GET("/perfis", controllerPerfil.GetPerfis)
	r.POST("/perfis", controllerPerfil.CreatePerfil)
	r.PUT("/perfis/:id", controllerPerfil.UpdatePerfil)
	r.DELETE("/perfis/:id", controllerPerfil.DeletePerfil)

	// SESSOES
	r.GET("/sessoes", controllerSessao.GetSessao)
	r.POST("/sessoes", controllerSessao.CreateSessao)
	r.PUT("/sessoes/:id", controllerSessao.UpdateSessao)
	r.DELETE("/sessoes/:id", controllerSessao.DeleteSessao)

	// MATCH
	r.POST("/matches", controllerMatch.CreateMatch)
	r.GET("/perfis/:id/matches", controllerMatch.GetMatchesByPerfil)

	return r
}
