package service

import (
	"match-dos-estudos/src/model"
	"match-dos-estudos/src/repository"
)

type PerfilService struct {
	repo *repository.PerfilRepository
}

func NewPerfilService(repo *repository.PerfilRepository) *PerfilService {
	return &PerfilService{
		repo: repo,
	}
}

func (s *PerfilService) GetAll() []model.Perfil {
	return s.repo.FindAll()
}

func (s *PerfilService) Create(perfil model.Perfil) model.Perfil {
	return s.repo.Save(perfil)
}

func (s *PerfilService) Update(id int, perfil model.Perfil) model.Perfil {
	return s.repo.Updat(id, perfil)
}

func (s *PerfilService) Delete(id int) []model.Perfil {
	return s.repo.Delet(id)
}

type SessaoService struct {
	repo *repository.SessaoRepository
}

func NewSessaoService(repo *repository.SessaoRepository) *SessaoService {
	return &SessaoService{
		repo: repo,
	}
}

func (s *SessaoService) GetAll() []model.Sessao {
	return s.repo.FindAll()
}

func (s *SessaoService) Create(sessao model.Sessao) model.Sessao {
	return s.repo.Save(sessao)
}

func (s *SessaoService) Update(id int, sessao model.Sessao) model.Sessao {
	return s.repo.Updat(id, sessao)
}

func (s *SessaoService) Delete(id int) []model.Sessao {
	return s.repo.Delet(id)
}

type MatchService struct {
	repo *repository.MatchRepository
}

func NewMatchService(repo *repository.MatchRepository) *MatchService {
	return &MatchService{
		repo: repo,
	}
}

func (s *MatchService) GetAll() []model.Match {
	return s.repo.FindAll()
}

func (s *MatchService) Create(match model.Match) model.Match {
	return s.repo.Save(match)
}
