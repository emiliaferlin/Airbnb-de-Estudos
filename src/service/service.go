package service

import (
	"errors"
	"strings"
	"time"

	"match-dos-estudos/src/model"
	"match-dos-estudos/src/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// ---------- Perfil ----------

type PerfilService struct {
	repo *repository.PerfilRepository
}

func NewPerfilService(repo *repository.PerfilRepository) *PerfilService {
	return &PerfilService{repo: repo}
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

// ---------- Sessao ----------

type SessaoService struct {
	repo *repository.SessaoRepository
}

func NewSessaoService(repo *repository.SessaoRepository) *SessaoService {
	return &SessaoService{repo: repo}
}

func (s *SessaoService) GetAll() []model.Sessao {
	return s.repo.FindAll()
}

func (s *SessaoService) Create(sessao model.Sessao) model.Sessao {
	return s.repo.Save(sessao)
}

func (s *SessaoService) Update(id int, sessao model.Sessao) model.Sessao {
	return s.repo.Updat(id, sessao)
}

func (s *SessaoService) Delete(id int) []model.Sessao {
	return s.repo.Delet(id)
}

// ---------- Match ----------

type MatchService struct {
	repo       *repository.MatchRepository
	repoPerfil *repository.PerfilRepository
	repoSessao *repository.SessaoRepository
}

func NewMatchService(
	repo *repository.MatchRepository,
	repoPerfil *repository.PerfilRepository,
	repoSessao *repository.SessaoRepository,
) *MatchService {
	return &MatchService{
		repo:       repo,
		repoPerfil: repoPerfil,
		repoSessao: repoSessao,
	}
}

// calcularScore compara os campos de um Perfil e uma Sessao.
// Critérios e pontuações:
//   - Mesma disciplina  → 40 pts
//   - Mesmo nível       → 30 pts
//   - Mesmo estilo      → 30 pts
//     Total máximo: 100 pts  |  Aprovado se score >= 60
func calcularScore(p model.Perfil, s model.Sessao) int {
	score := 0
	if strings.EqualFold(p.Disciplina, s.Disciplina) {
		score += 40
	}
	if strings.EqualFold(p.Nivel, s.Nivel) {
		score += 30
	}
	if strings.EqualFold(p.Estilo, s.Estilo) {
		score += 30
	}
	return score
}

// Create recebe um Match com PerfilID e SessaoID, calcula o score,
// define Aprovado e persiste o resultado.
func (s *MatchService) Create(match model.Match) (model.Match, error) {
	perfil, okP := s.repoPerfil.FindByID(match.PerfilID)
	if !okP {
		return model.Match{}, &NotFoundError{"perfil", match.PerfilID}
	}

	sessao, okS := s.repoSessao.FindByID(match.SessaoID)
	if !okS {
		return model.Match{}, &NotFoundError{"sessao", match.SessaoID}
	}

	match.Score = calcularScore(perfil, sessao)
	match.Aprovado = match.Score >= 60

	return s.repo.Save(match), nil
}

// GetByPerfilID retorna todos os matches aprovados de um perfil.
func (s *MatchService) GetByPerfilID(perfilID int) ([]model.Match, error) {
	_, ok := s.repoPerfil.FindByID(perfilID)
	if !ok {
		return nil, &NotFoundError{"perfil", perfilID}
	}
	return s.repo.FindByPerfilID(perfilID), nil
}

// ---------- Erro customizado ----------

type NotFoundError struct {
	Entidade string
	ID       int
}

func (e *NotFoundError) Error() string {
	return e.Entidade + " não encontrado(a)"
}

type UsuarioService struct {
	repo *repository.UsuarioRepository
}

func NewUsuarioService(repo *repository.UsuarioRepository) *UsuarioService {
	return &UsuarioService{repo: repo}
}

func (s *UsuarioService) Register(email, senha string) (model.Usuario, error) {
	// Verifica se email já existe
	_, existe := s.repo.FindByEmail(email)
	if existe {
		return model.Usuario{}, errors.New("email já cadastrado")
	}

	// Gera o hash da senha
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		return model.Usuario{}, errors.New("erro ao processar senha")
	}

	usuario := model.Usuario{
		Email: email,
		Senha: string(hash),
	}
	return s.repo.Save(usuario), nil
}

func (s *UsuarioService) Login(email, senha string) (string, error) {
	// Busca o usuário
	usuario, ok := s.repo.FindByEmail(email)
	if !ok {
		return "", errors.New("credenciais inválidas")
	}

	// Compara a senha com o hash
	err := bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(senha))
	if err != nil {
		return "", errors.New("credenciais inválidas")
	}

	// Gera o token JWT
	claims := jwt.MapClaims{
		"sub":   usuario.ID,
		"email": usuario.Email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("minha_chave_secreta"))
}
