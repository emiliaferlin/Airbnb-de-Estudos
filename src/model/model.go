package model

type Perfil struct {
	ID         int    `json:"id"`
	Nome       string `json:"nome"`
	Idade      int    `json:"idade"`
	Disciplina string `json:"disciplina"`
	Nivel      string `json:"nivel"`  //sem conhecimento - conhecimento médio - conhecimento alto
	Estilo     string `json:"estilo"` // mais silencioso - gosta de argumentar
}

type Sessao struct {
	ID             int    `json:"id"`
	Titulo         string `json:"titulo"`
	Disciplina     string `json:"disciplina"`
	Nivel          string `json:"nivel"`  //sem conhecimento - conhecimento médio - conhecimento alto
	Estilo         string `json:"estilo"` // mais silencioso - gosta de argumentar
	DataHoraInicio string `json:"dataHoraInicio"`
	DuracaoMinutos int    `json:"duracaoMinutos"`
	Vagas          int    `json:"vagas"`
}

type Match struct {
	ID       int  `json:"id"`
	PerfilID int  `json:"perfilId"`
	SessaoID int  `json:"sessaoId"`
	Score    int  `json:"score"`
	Aceito   bool `json:"aceito"` // usuário aceitou ou não
}
