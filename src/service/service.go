package service

import (
	"match-dos-estudos/src/model"
	"match-dos-estudos/src/repository"
)

type PerfilService struct {
	repo *repository.PerfilRepository
}

func NewPerfilService() *PerfilService {
	return &PerfilService{
		repo: repository.NewPerfilRepository(),
	}
}

func (s *PerfilService) GetAll() []model.Perfil {
	return s.repo.FindAll()
}
