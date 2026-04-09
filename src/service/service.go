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
