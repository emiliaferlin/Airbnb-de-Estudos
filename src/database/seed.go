package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// Seed insere dados de exemplo nas coleções caso elas estejam vazias.
// Garante que toda execução do servidor tenha dados prontos para demonstração.
func Seed() {
	seedPerfis()
	seedSessoes()
	seedMatches()
}

func seedPerfis() {
	col := GetCollection("perfis")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, _ := col.CountDocuments(ctx, bson.D{})
	if count > 0 {
		return // já populado
	}

	perfis := []interface{}{
		bson.D{
			{Key: "_id", Value: 1},
			{Key: "nome", Value: "Ana Lima"},
			{Key: "idade", Value: 22},
			{Key: "disciplina", Value: "Algoritmos"},
			{Key: "nivel", Value: "conhecimento médio"},
			{Key: "estilo", Value: "gosta de argumentar"},
		},
		bson.D{
			{Key: "_id", Value: 2},
			{Key: "nome", Value: "Bruno Ferreira"},
			{Key: "idade", Value: 25},
			{Key: "disciplina", Value: "Banco de Dados"},
			{Key: "nivel", Value: "conhecimento alto"},
			{Key: "estilo", Value: "mais silencioso"},
		},
		bson.D{
			{Key: "_id", Value: 3},
			{Key: "nome", Value: "Carla Santos"},
			{Key: "idade", Value: 20},
			{Key: "disciplina", Value: "Algoritmos"},
			{Key: "nivel", Value: "conhecimento médio"},
			{Key: "estilo", Value: "gosta de argumentar"},
		},
		bson.D{
			{Key: "_id", Value: 4},
			{Key: "nome", Value: "Diego Ramos"},
			{Key: "idade", Value: 23},
			{Key: "disciplina", Value: "Redes"},
			{Key: "nivel", Value: "sem conhecimento"},
			{Key: "estilo", Value: "mais silencioso"},
		},
		bson.D{
			{Key: "_id", Value: 5},
			{Key: "nome", Value: "Emilia Ferlin"},
			{Key: "idade", Value: 21},
			{Key: "disciplina", Value: "Banco de Dados"},
			{Key: "nivel", Value: "conhecimento médio"},
			{Key: "estilo", Value: "gosta de argumentar"},
		},
	}

	_, err := col.InsertMany(ctx, perfis)
	if err != nil {
		log.Printf("Seed perfis: %v", err)
		return
	}
	log.Println("Seed: 5 perfis inseridos")
}

func seedSessoes() {
	col := GetCollection("sessoes")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, _ := col.CountDocuments(ctx, bson.D{})
	if count > 0 {
		return
	}

	sessoes := []interface{}{
		bson.D{
			{Key: "_id", Value: 1},
			{Key: "titulo", Value: "Revisão de Algoritmos"},
			{Key: "disciplina", Value: "Algoritmos"},
			{Key: "nivel", Value: "conhecimento médio"},
			{Key: "estilo", Value: "gosta de argumentar"},
			{Key: "dataHoraInicio", Value: "2026-04-15T19:00:00"},
			{Key: "duracaoMinutos", Value: 90},
			{Key: "vagas", Value: 4},
		},
		bson.D{
			{Key: "_id", Value: 2},
			{Key: "titulo", Value: "SQL Avançado"},
			{Key: "disciplina", Value: "Banco de Dados"},
			{Key: "nivel", Value: "conhecimento alto"},
			{Key: "estilo", Value: "mais silencioso"},
			{Key: "dataHoraInicio", Value: "2026-04-16T18:00:00"},
			{Key: "duracaoMinutos", Value: 120},
			{Key: "vagas", Value: 3},
		},
		bson.D{
			{Key: "_id", Value: 3},
			{Key: "titulo", Value: "Fundamentos de Redes"},
			{Key: "disciplina", Value: "Redes"},
			{Key: "nivel", Value: "sem conhecimento"},
			{Key: "estilo", Value: "mais silencioso"},
			{Key: "dataHoraInicio", Value: "2026-04-17T20:00:00"},
			{Key: "duracaoMinutos", Value: 60},
			{Key: "vagas", Value: 5},
		},
		bson.D{
			{Key: "_id", Value: 4},
			{Key: "titulo", Value: "Estrutura de Dados na Prática"},
			{Key: "disciplina", Value: "Algoritmos"},
			{Key: "nivel", Value: "conhecimento médio"},
			{Key: "estilo", Value: "gosta de argumentar"},
			{Key: "dataHoraInicio", Value: "2026-04-18T19:30:00"},
			{Key: "duracaoMinutos", Value: 90},
			{Key: "vagas", Value: 4},
		},
		bson.D{
			{Key: "_id", Value: 5},
			{Key: "titulo", Value: "NoSQL e MongoDB"},
			{Key: "disciplina", Value: "Banco de Dados"},
			{Key: "nivel", Value: "conhecimento médio"},
			{Key: "estilo", Value: "gosta de argumentar"},
			{Key: "dataHoraInicio", Value: "2026-04-19T14:00:00"},
			{Key: "duracaoMinutos", Value: 120},
			{Key: "vagas", Value: 6},
		},
	}

	_, err := col.InsertMany(ctx, sessoes)
	if err != nil {
		log.Printf("Seed sessoes: %v", err)
		return
	}
	log.Println("Seed: 5 sessões inseridas")
}

func seedMatches() {
	col := GetCollection("matches")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, _ := col.CountDocuments(ctx, bson.D{})
	if count > 0 {
		return
	}

	// Matches pré-calculados para demonstração:
	// Ana (perfil 1) x Revisão de Algoritmos (sessão 1): disciplina✓ nível✓ estilo✓ = 100 → aprovado
	// Ana (perfil 1) x SQL Avançado (sessão 2):          disciplina✗ nível✗ estilo✗ = 0   → reprovado
	// Bruno (perfil 2) x SQL Avançado (sessão 2):        disciplina✓ nível✓ estilo✓ = 100 → aprovado
	// Carla (perfil 3) x Revisão de Algoritmos (sessão 1): disciplina✓ nível✓ estilo✓ = 100 → aprovado
	// Diego (perfil 4) x Fundamentos de Redes (sessão 3): disciplina✓ nível✓ estilo✓ = 100 → aprovado
	// Emilia (perfil 5) x NoSQL e MongoDB (sessão 5):    disciplina✓ nível✓ estilo✓ = 100 → aprovado
	matches := []interface{}{
		bson.D{
			{Key: "_id", Value: 1},
			{Key: "perfilId", Value: 1},
			{Key: "sessaoId", Value: 1},
			{Key: "score", Value: 100},
			{Key: "aprovado", Value: true},
		},
		bson.D{
			{Key: "_id", Value: 2},
			{Key: "perfilId", Value: 1},
			{Key: "sessaoId", Value: 2},
			{Key: "score", Value: 0},
			{Key: "aprovado", Value: false},
		},
		bson.D{
			{Key: "_id", Value: 3},
			{Key: "perfilId", Value: 2},
			{Key: "sessaoId", Value: 2},
			{Key: "score", Value: 100},
			{Key: "aprovado", Value: true},
		},
		bson.D{
			{Key: "_id", Value: 4},
			{Key: "perfilId", Value: 3},
			{Key: "sessaoId", Value: 1},
			{Key: "score", Value: 100},
			{Key: "aprovado", Value: true},
		},
		bson.D{
			{Key: "_id", Value: 5},
			{Key: "perfilId", Value: 4},
			{Key: "sessaoId", Value: 3},
			{Key: "score", Value: 100},
			{Key: "aprovado", Value: true},
		},
		bson.D{
			{Key: "_id", Value: 6},
			{Key: "perfilId", Value: 5},
			{Key: "sessaoId", Value: 5},
			{Key: "score", Value: 100},
			{Key: "aprovado", Value: true},
		},
	}

	_, err := col.InsertMany(ctx, matches)
	if err != nil {
		log.Printf("Seed matches: %v", err)
		return
	}
	log.Println("Seed: 6 matches inseridos")
}
