package repository

import (
	"context"
	"errors"
	"time"

	"match-dos-estudos/src/database"
	"match-dos-estudos/src/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// ctx rápido para operações de banco
func ctx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

// ---------- Helpers de ID auto-incremento ----------
// O MongoDB usa ObjectID por padrão, mas para manter compatibilidade
// com o restante do sistema (int ID), usamos um campo _id inteiro
// e buscamos o próximo valor pela maior chave existente na coleção.

func nextID(col *mongo.Collection) int {
	c, cancel := ctx()
	defer cancel()
	opts := options.FindOne().SetSort(bson.D{{Key: "_id", Value: -1}})
	var result bson.M
	if err := col.FindOne(c, bson.D{}, opts).Decode(&result); err != nil {
		return 1
	}
	id, _ := result["_id"].(int32)
	return int(id) + 1
}

// ================================================================
// PERFIL REPOSITORY
// ================================================================

type PerfilRepository struct {
	col *mongo.Collection
}

func NewPerfilRepository() *PerfilRepository {
	return &PerfilRepository{col: database.GetCollection("perfis")}
}

func (r *PerfilRepository) FindAll() []model.Perfil {
	c, cancel := ctx()
	defer cancel()

	cursor, err := r.col.Find(c, bson.D{})
	if err != nil {
		return []model.Perfil{}
	}
	defer cursor.Close(c)

	var perfis []model.Perfil
	for cursor.Next(c) {
		var p model.Perfil
		if err := cursor.Decode(&p); err == nil {
			perfis = append(perfis, p)
		}
	}
	return perfis
}

func (r *PerfilRepository) FindByID(id int) (model.Perfil, bool) {
	c, cancel := ctx()
	defer cancel()

	var p model.Perfil
	err := r.col.FindOne(c, bson.D{{Key: "_id", Value: id}}).Decode(&p)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return model.Perfil{}, false
	}
	return p, err == nil
}

func (r *PerfilRepository) Save(perfil model.Perfil) model.Perfil {
	c, cancel := ctx()
	defer cancel()

	perfil.ID = nextID(r.col)
	_, _ = r.col.InsertOne(c, perfil)
	return perfil
}

func (r *PerfilRepository) Updat(id int, perfil model.Perfil) model.Perfil {
	c, cancel := ctx()
	defer cancel()

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "nome", Value: perfil.Nome},
		{Key: "idade", Value: perfil.Idade},
		{Key: "disciplina", Value: perfil.Disciplina},
		{Key: "nivel", Value: perfil.Nivel},
		{Key: "estilo", Value: perfil.Estilo},
	}}}

	result := r.col.FindOneAndUpdate(c,
		bson.D{{Key: "_id", Value: id}},
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)

	var updated model.Perfil
	if err := result.Decode(&updated); err != nil {
		return model.Perfil{}
	}
	return updated
}

func (r *PerfilRepository) Delet(id int) []model.Perfil {
	c, cancel := ctx()
	defer cancel()

	res, err := r.col.DeleteOne(c, bson.D{{Key: "_id", Value: id}})
	if err != nil || res.DeletedCount == 0 {
		return []model.Perfil{}
	}
	return r.FindAll()
}

// ================================================================
// SESSAO REPOSITORY
// ================================================================

type SessaoRepository struct {
	col *mongo.Collection
}

func NewSessaoRepository() *SessaoRepository {
	return &SessaoRepository{col: database.GetCollection("sessoes")}
}

func (r *SessaoRepository) FindAll() []model.Sessao {
	c, cancel := ctx()
	defer cancel()

	cursor, err := r.col.Find(c, bson.D{})
	if err != nil {
		return []model.Sessao{}
	}
	defer cursor.Close(c)

	var sessoes []model.Sessao
	for cursor.Next(c) {
		var s model.Sessao
		if err := cursor.Decode(&s); err == nil {
			sessoes = append(sessoes, s)
		}
	}
	return sessoes
}

func (r *SessaoRepository) FindByID(id int) (model.Sessao, bool) {
	c, cancel := ctx()
	defer cancel()

	var s model.Sessao
	err := r.col.FindOne(c, bson.D{{Key: "_id", Value: id}}).Decode(&s)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return model.Sessao{}, false
	}
	return s, err == nil
}

func (r *SessaoRepository) Save(sessao model.Sessao) model.Sessao {
	c, cancel := ctx()
	defer cancel()

	sessao.ID = nextID(r.col)
	_, _ = r.col.InsertOne(c, sessao)
	return sessao
}

func (r *SessaoRepository) Updat(id int, sessao model.Sessao) model.Sessao {
	c, cancel := ctx()
	defer cancel()

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "titulo", Value: sessao.Titulo},
		{Key: "disciplina", Value: sessao.Disciplina},
		{Key: "nivel", Value: sessao.Nivel},
		{Key: "estilo", Value: sessao.Estilo},
		{Key: "dataHoraInicio", Value: sessao.DataHoraInicio},
		{Key: "duracaoMinutos", Value: sessao.DuracaoMinutos},
		{Key: "vagas", Value: sessao.Vagas},
	}}}

	result := r.col.FindOneAndUpdate(c,
		bson.D{{Key: "_id", Value: id}},
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)

	var updated model.Sessao
	if err := result.Decode(&updated); err != nil {
		return model.Sessao{}
	}
	return updated
}

func (r *SessaoRepository) Delet(id int) []model.Sessao {
	c, cancel := ctx()
	defer cancel()

	res, err := r.col.DeleteOne(c, bson.D{{Key: "_id", Value: id}})
	if err != nil || res.DeletedCount == 0 {
		return []model.Sessao{}
	}
	return r.FindAll()
}

// ================================================================
// MATCH REPOSITORY
// ================================================================

type MatchRepository struct {
	col *mongo.Collection
}

func NewMatchRepository() *MatchRepository {
	return &MatchRepository{col: database.GetCollection("matches")}
}

func (r *MatchRepository) FindAll() []model.Match {
	c, cancel := ctx()
	defer cancel()

	cursor, err := r.col.Find(c, bson.D{})
	if err != nil {
		return []model.Match{}
	}
	defer cursor.Close(c)

	var matches []model.Match
	for cursor.Next(c) {
		var m model.Match
		if err := cursor.Decode(&m); err == nil {
			matches = append(matches, m)
		}
	}
	return matches
}

// FindByPerfilID retorna apenas os matches aprovados de um perfil específico.
func (r *MatchRepository) FindByPerfilID(perfilID int) []model.Match {
	c, cancel := ctx()
	defer cancel()

	filter := bson.D{
		{Key: "perfilId", Value: perfilID},
		{Key: "aprovado", Value: true},
	}

	cursor, err := r.col.Find(c, filter)
	if err != nil {
		return []model.Match{}
	}
	defer cursor.Close(c)

	var matches []model.Match
	for cursor.Next(c) {
		var m model.Match
		if err := cursor.Decode(&m); err == nil {
			matches = append(matches, m)
		}
	}
	return matches
}

func (r *MatchRepository) Save(match model.Match) model.Match {
	c, cancel := ctx()
	defer cancel()

	match.ID = nextID(r.col)
	_, _ = r.col.InsertOne(c, match)
	return match
}

type UsuarioRepository struct {
	col *mongo.Collection
}

func NewUsuarioRepository() *UsuarioRepository {
	return &UsuarioRepository{col: database.GetCollection("usuarios")}
}

func (r *UsuarioRepository) FindByEmail(email string) (model.Usuario, bool) {
	c, cancel := ctx()
	defer cancel()

	var u model.Usuario
	err := r.col.FindOne(c, bson.D{{Key: "email", Value: email}}).Decode(&u)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return model.Usuario{}, false
	}
	return u, err == nil
}

func (r *UsuarioRepository) Save(usuario model.Usuario) model.Usuario {
	c, cancel := ctx()
	defer cancel()

	usuario.ID = nextID(r.col)
	_, _ = r.col.InsertOne(c, usuario)
	return usuario
}
