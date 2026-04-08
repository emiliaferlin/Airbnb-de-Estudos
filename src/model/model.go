package model

type Perfil struct {
	ID         int    `json:"id"`
	Nome       string `json:"nome"`
	Disciplina string `json:"disciplina"`
	Nivel      string `json:"nivel"`
	Estilo     string `json:"estilo"`
}
