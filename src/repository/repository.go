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

func (r *PerfilRepository) FindByID(id int) (model.Perfil, bool) {
	for _, p := range r.perfis {
		if p.ID == id {
			return p, true
		}
	}
	return model.Perfil{}, false
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

// ---------- Sessao ----------

type SessaoRepository struct {
	sessoes []model.Sessao
}

func NewSessaoRepository() *SessaoRepository {
	return &SessaoRepository{
		sessoes: []model.Sessao{},
	}
}

func (r *SessaoRepository) FindAll() []model.Sessao {
	return r.sessoes
}

func (r *SessaoRepository) FindByID(id int) (model.Sessao, bool) {
	for _, s := range r.sessoes {
		if s.ID == id {
			return s, true
		}
	}
	return model.Sessao{}, false
}

func (r *SessaoRepository) Save(sessao model.Sessao) model.Sessao {
	sessao.ID = len(r.sessoes) + 1
	r.sessoes = append(r.sessoes, sessao)
	return sessao
}

func (r *SessaoRepository) Updat(id int, sessao model.Sessao) model.Sessao {
	for i, a := range r.sessoes {
		if a.ID == id {
			r.sessoes[i].Titulo = sessao.Titulo
			r.sessoes[i].DataHoraInicio = sessao.DataHoraInicio
			r.sessoes[i].DuracaoMinutos = sessao.DuracaoMinutos
			r.sessoes[i].Nivel = sessao.Nivel
			r.sessoes[i].Disciplina = sessao.Disciplina
			r.sessoes[i].Estilo = sessao.Estilo
			r.sessoes[i].Vagas = sessao.Vagas
			return r.sessoes[i]
		}
	}
	return model.Sessao{}
}

func (r *SessaoRepository) Delet(id int) []model.Sessao {
	for i, a := range r.sessoes {
		if a.ID == id {
			r.sessoes = append(r.sessoes[:i], r.sessoes[i+1:]...)
			return r.sessoes
		}
	}
	return []model.Sessao{}
}

// ---------- Match ----------

type MatchRepository struct {
	matchs []model.Match
}

func NewMatchRepository() *MatchRepository {
	return &MatchRepository{
		matchs: []model.Match{},
	}
}

func (r *MatchRepository) FindAll() []model.Match {
	return r.matchs
}

// FindByPerfilID retorna apenas os matches aprovados de um perfil específico
func (r *MatchRepository) FindByPerfilID(perfilID int) []model.Match {
	var resultado []model.Match
	for _, m := range r.matchs {
		if m.PerfilID == perfilID && m.Aprovado {
			resultado = append(resultado, m)
		}
	}
	return resultado
}

func (r *MatchRepository) Save(match model.Match) model.Match {
	match.ID = len(r.matchs) + 1
	r.matchs = append(r.matchs, match)
	return match
}
