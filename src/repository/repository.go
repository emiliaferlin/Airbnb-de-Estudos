package repository

import (
	"match-dos-estudos/src/model"
)

type PerfilRepository struct {
	perfis []model.Perfil
}

func NewPerfilRepository() *PerfilRepository {
	return &PerfilRepository{
		perfis: []model.Perfil{},
	}
}

func (r *PerfilRepository) FindAll() []model.Perfil {
	return r.perfis
}

func (r *PerfilRepository) Save(perfil model.Perfil) model.Perfil {
	perfil.ID = len(r.perfis) + 1
	r.perfis = append(r.perfis, perfil)
	return perfil
}

func (r *PerfilRepository) Updat(id int, perfil model.Perfil) model.Perfil {
	for i, a := range r.perfis {
		if a.ID == id {
			r.perfis[i].Nome = perfil.Nome
			r.perfis[i].Idade = perfil.Idade
			r.perfis[i].Nivel = perfil.Nivel
			r.perfis[i].Disciplina = perfil.Disciplina
			r.perfis[i].Estilo = perfil.Estilo

			return r.perfis[i]
		}
	}
	return model.Perfil{}
}

func (r *PerfilRepository) Delet(id int) []model.Perfil {
	for i, a := range r.perfis {
		if a.ID == id {
			r.perfis = append(r.perfis[:i], r.perfis[i+1:]...)
			return r.perfis
		}
	}
	return []model.Perfil{}
}
