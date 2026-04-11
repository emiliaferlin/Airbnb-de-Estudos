package router

import (
	"match-dos-estudos/src/controller"
	"match-dos-estudos/src/middleware"
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
	repoAuth := repository.NewUsuarioRepository()

	// Services
	servicePerfil := service.NewPerfilService(repoPerfil)
	serviceSessao := service.NewSessaoService(repoSessao)
	serviceMatch := service.NewMatchService(repoMatch, repoPerfil, repoSessao)
	serviceAuth := service.NewUsuarioService(repoAuth) // ← completar aqui

	// Controllers
	controllerPerfil := controller.NewPerfilController(servicePerfil)
	controllerSessao := controller.NewSessaoController(serviceSessao)
	controllerMatch := controller.NewMatchController(serviceMatch)
	controllerAuth := controller.NewAuthController(serviceAuth) // ← adicionar aqui

	// Rotas públicas
	r.POST("/login", controllerAuth.Login)
	r.POST("/register", controllerAuth.Register)

	// Rotas protegidas
	protegidas := r.Group("/")
	protegidas.Use(middleware.AuthMiddleware())
	{
		protegidas.POST("/perfis", controllerPerfil.CreatePerfil)
		protegidas.PUT("/perfis/:id", controllerPerfil.UpdatePerfil)
		protegidas.DELETE("/perfis/:id", controllerPerfil.DeletePerfil)

		protegidas.POST("/sessoes", controllerSessao.CreateSessao)
		protegidas.PUT("/sessoes/:id", controllerSessao.UpdateSessao)
		protegidas.DELETE("/sessoes/:id", controllerSessao.DeleteSessao)

		protegidas.POST("/matches", controllerMatch.CreateMatch)
	}

	// Rotas de leitura abertas
	r.GET("/perfis", controllerPerfil.GetPerfis)
	r.GET("/sessoes", controllerSessao.GetSessao)
	r.GET("/perfis/:id/matches", controllerMatch.GetMatchesByPerfil)

	return r
}
