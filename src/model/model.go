package model

// As tags bson mapeiam os campos Go para os campos do MongoDB.
// O campo ID usa "_id" que é a chave primária padrão do MongoDB.

type Perfil struct {
	ID         int    `json:"id"         bson:"_id"`
	Nome       string `json:"nome"       bson:"nome"`
	Idade      int    `json:"idade"      bson:"idade"`
	Disciplina string `json:"disciplina" bson:"disciplina"`
	Nivel      string `json:"nivel"      bson:"nivel"`  // sem conhecimento | conhecimento médio | conhecimento alto
	Estilo     string `json:"estilo"     bson:"estilo"` // mais silencioso | gosta de argumentar
}

type Sessao struct {
	ID             int    `json:"id"             bson:"_id"`
	Titulo         string `json:"titulo"         bson:"titulo"`
	Disciplina     string `json:"disciplina"     bson:"disciplina"`
	Nivel          string `json:"nivel"          bson:"nivel"`
	Estilo         string `json:"estilo"         bson:"estilo"`
	DataHoraInicio string `json:"dataHoraInicio" bson:"dataHoraInicio"`
	DuracaoMinutos int    `json:"duracaoMinutos" bson:"duracaoMinutos"`
	Vagas          int    `json:"vagas"          bson:"vagas"`
}

type Match struct {
	ID       int  `json:"id"       bson:"_id"`
	PerfilID int  `json:"perfilId" bson:"perfilId"`
	SessaoID int  `json:"sessaoId" bson:"sessaoId"`
	Score    int  `json:"score"    bson:"score"`
	Aprovado bool `json:"aprovado" bson:"aprovado"`
}
