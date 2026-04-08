package repository

import "match-dos-estudos/src/model"

type PerfilRepository struct{}

func NewPerfilRepository() *PerfilRepository {
	return &PerfilRepository{}
}

func (r *PerfilRepository) FindAll() []model.Perfil {
	return []model.Perfil{
		{
			ID:         1,
			Nome:       "Emilia",
			Disciplina: "Algoritmos",
			Nivel:      "Intermediario",
			Estilo:     "Colaborativo",
		},
	}
}
