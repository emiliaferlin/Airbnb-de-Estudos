# Match dos Estudos

> API REST em Go para conectar estudantes com sessões de estudo compatíveis com seu perfil.

---

## 1. Descrição do Domínio

### Problema

Estudantes que desejam estudar em grupo frequentemente enfrentam dificuldades para encontrar sessões com disciplina, nível de conhecimento e estilo de aprendizagem compatíveis. Essa falta de conexão reduz a produtividade e desmotiva a aprendizagem colaborativa.

### Solução

O **Match dos Estudos** é uma API REST que permite cadastrar perfis de estudantes e sessões de estudo, e então calcular automaticamente a compatibilidade entre eles através de um algoritmo de pontuação (_score_). Um perfil recebe um match **aprovado** quando atinge pontuação suficiente com uma sessão.

### Entidades do Domínio

| Entidade  | Descrição                                                                                      |
| --------- | ---------------------------------------------------------------------------------------------- |
| `Perfil`  | Características acadêmicas do estudante: disciplina, nível e estilo de estudo                  |
| `Sessao`  | Sessão de estudo organizada, com título, disciplina, nível, estilo, data/hora, duração e vagas |
| `Match`   | Resultado da compatibilidade entre um perfil e uma sessão: score e status de aprovação         |
| `Usuario` | Usuário autenticado com email e senha para acesso às rotas protegidas                          |

### Algoritmo de Match

O score compara os campos de um `Perfil` com os de uma `Sessao`. Match **aprovado** quando score ≥ 60.

| Critério                    | Pontos  |
| --------------------------- | ------- |
| Mesma disciplina            | 40      |
| Mesmo nível de conhecimento | 30      |
| Mesmo estilo de estudo      | 30      |
| **Total máximo**            | **100** |

---

## 2. Instalação e Execução Local

### Pré-requisitos

- [Go 1.21+](https://go.dev/dl/)
- [Docker](https://www.docker.com/) e Docker Compose

### Passos

```bash
# 1. Clone o repositório
git clone https://github.com/emiliaferlin/Match-dos-Estudos.git
cd Match-dos-Estudos

# 2. Suba o banco de dados MongoDB
docker-compose up -d

# 3. Instale as dependências Go
go mod tidy

# 4. Execute o servidor
go run ./src/main.go
```

O servidor iniciará em `http://localhost:8080`.

> **Ao iniciar**, o servidor conecta ao MongoDB e executa o **seed automático**, inserindo 2 usuários, 5 perfis, 5 sessões e 6 matches de exemplo — caso as coleções ainda estejam vazias.

### Credenciais de teste (seed)

| Email              | Senha      |
| ------------------ | ---------- |
| `admin@match.com`  | `senha123` |
| `emilia@match.com` | `senha123` |

### Parar o banco

```bash
docker-compose down
```

---

## 3. Autenticação JWT

As rotas de criação, atualização e exclusão são protegidas por **JWT (JSON Web Token)**. Para acessá-las:

### 1. Registrar um usuário

**POST** `/register`

```json
{ "email": "seu@email.com", "senha": "suasenha" }
```

### 2. Fazer login e obter o token

**POST** `/login`

```json
{ "email": "admin@match.com", "senha": "senha123" }
```

Resposta:

```json
{ "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." }
```

### 3. Usar o token nas rotas protegidas

Adicione o header em todas as requisições marcadas com 🔒:

```
Authorization: Bearer <token>
```

---

## 4. Tabela de Rotas da API

### 4.1 Autenticação

| Método | Rota        | Descrição                       | Status |
| ------ | ----------- | ------------------------------- | ------ |
| `POST` | `/register` | Cria um novo usuário            | 201    |
| `POST` | `/login`    | Autentica e retorna o token JWT | 200    |

### 4.2 Perfis

| Método   | Rota          | Descrição             | Status | Auth   |
| -------- | ------------- | --------------------- | ------ | ------ |
| `GET`    | `/perfis`     | Lista todos os perfis | 200    | —      |
| `POST`   | `/perfis`     | Cria um novo perfil   | 201    | 🔒 JWT |
| `PUT`    | `/perfis/:id` | Atualiza um perfil    | 200    | 🔒 JWT |
| `DELETE` | `/perfis/:id` | Remove um perfil      | 200    | 🔒 JWT |

**Body — POST/PUT `/perfis`:**

```json
{
  "nome": "Ana Lima",
  "idade": 22,
  "disciplina": "Algoritmos",
  "nivel": "conhecimento médio",
  "estilo": "gosta de argumentar"
}
```

### 4.3 Sessões

| Método   | Rota           | Descrição              | Status | Auth   |
| -------- | -------------- | ---------------------- | ------ | ------ |
| `GET`    | `/sessoes`     | Lista todas as sessões | 200    | —      |
| `POST`   | `/sessoes`     | Cria uma nova sessão   | 201    | 🔒 JWT |
| `PUT`    | `/sessoes/:id` | Atualiza uma sessão    | 200    | 🔒 JWT |
| `DELETE` | `/sessoes/:id` | Remove uma sessão      | 200    | 🔒 JWT |

**Body — POST/PUT `/sessoes`:**

```json
{
  "titulo": "Revisão de Algoritmos",
  "disciplina": "Algoritmos",
  "nivel": "conhecimento médio",
  "estilo": "gosta de argumentar",
  "dataHoraInicio": "2026-04-15T19:00:00",
  "duracaoMinutos": 90,
  "vagas": 4
}
```

### 4.4 Matches

| Método | Rota                  | Descrição                             | Status | Auth   |
| ------ | --------------------- | ------------------------------------- | ------ | ------ |
| `POST` | `/matches`            | Calcula o score entre perfil e sessão | 201    | 🔒 JWT |
| `GET`  | `/perfis/:id/matches` | Lista matches aprovados do perfil     | 200    | —      |

**Body — POST `/matches`:**

```json
{
  "perfilId": 1,
  "sessaoId": 2
}
```

**Resposta — POST `/matches` (aprovado):**

```json
{ "id": 7, "perfilId": 1, "sessaoId": 1, "score": 100, "aprovado": true }
```

**Resposta — GET `/perfis/1/matches`:**

```json
[{ "id": 1, "perfilId": 1, "sessaoId": 1, "score": 100, "aprovado": true }]
```

---

## 5. Documentação OpenAPI (Swagger)

A documentação interativa da API está disponível com o servidor rodando em:

```
http://localhost:8080/api-docs/index.html
```

Gerada automaticamente com [swaggo/swag](https://github.com/swaggo/swag) a partir das anotações nos handlers. Para regenerar após alterações:

```bash
swag init -g src/main.go
```

---

## 6. Banco de Dados — MongoDB

### Por que MongoDB?

O projeto usa o **MongoDB** como banco de dados NoSQL, com o driver oficial **`mongo-driver/v2`** como ODM (Object-Document Mapper). O MongoDB se adequa bem ao domínio por:

- Documentos JSON/BSON espelham diretamente as structs Go, sem necessidade de mapeamento relacional complexo
- Flexibilidade de schema permite evoluir os campos de Perfil e Sessao sem migrações
- Consultas por campos arbitrários (como `perfilId + aprovado`) são simples e eficientes

### Mapeamento (ODM) com `mongo-driver/v2`

As structs do pacote `model` usam **tags `bson`** para mapear campos Go ↔ MongoDB:

```go
type Perfil struct {
    ID         int    `json:"id"         bson:"_id"`
    Nome       string `json:"nome"       bson:"nome"`
    Disciplina string `json:"disciplina" bson:"disciplina"`
    Nivel      string `json:"nivel"      bson:"nivel"`
    Estilo     string `json:"estilo"     bson:"estilo"`
}
```

O campo `_id` é a chave primária do MongoDB. O projeto usa IDs inteiros auto-incrementais para manter compatibilidade com as rotas REST (ex: `/perfis/1`).

### Coleções

| Coleção    | Descrição                                  |
| ---------- | ------------------------------------------ |
| `usuarios` | Documentos de `Usuario` com senha hasheada |
| `perfis`   | Documentos de `Perfil`                     |
| `sessoes`  | Documentos de `Sessao`                     |
| `matches`  | Documentos de `Match` com score e aprovado |

### Seed de dados

O arquivo `src/database/seed.go` popula automaticamente as coleções na primeira execução:

| Coleção    | Registros de exemplo                                                           |
| ---------- | ------------------------------------------------------------------------------ |
| `usuarios` | 2 usuários de teste (admin@match.com e emilia@match.com, senha: senha123)      |
| `perfis`   | 5 perfis (Ana, Bruno, Carla, Diego, Emilia) com disciplinas e estilos variados |
| `sessoes`  | 5 sessões (Algoritmos, Banco de Dados, Redes) em diferentes níveis             |
| `matches`  | 6 matches pré-calculados demonstrando resultados aprovados e reprovados        |

---

## 7. Estrutura do Projeto

```
Match-dos-Estudos/
├── docker-compose.yml           # Sobe o MongoDB na porta 27017
├── docs/                        # Gerado pelo swaggo (swagger.json, docs.go)
├── src/
│   ├── main.go                  # Inicializa banco, seed e servidor
│   ├── database/
│   │   ├── connection.go        # Conexão com MongoDB
│   │   └── seed.go              # Dados de exemplo (inserção idempotente)
│   ├── router/
│   │   └── router.go            # Registro de rotas e injeção de dependências
│   ├── controller/
│   │   └── controller.go        # Handlers HTTP com anotações Swagger
│   ├── service/
│   │   └── service.go           # Regras de negócio e algoritmo de score
│   ├── repository/
│   │   └── repository.go        # Acesso ao MongoDB (ODM com bson tags)
│   ├── middleware/
│   │   ├── auth.go              # Middleware de validação JWT
│   │   └── error.go             # Middleware de tratamento centralizado de erros
│   └── model/
│       └── model.go             # Structs com tags json e bson
├── testes/
│   ├── postman_collection.json  # Collection exportada do Postman
│   └── screenshots/             # Capturas de tela dos testes
├── go.mod
├── go.sum
└── README.md
```

---

## 8. Exemplos de Uso

```bash
# 1. Login para obter token
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@match.com","senha":"senha123"}'

# 2. Listar perfis (aberto)
curl http://localhost:8080/perfis

# 3. Criar perfil (requer token)
curl -X POST http://localhost:8080/perfis \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"nome":"João","idade":24,"disciplina":"Redes","nivel":"conhecimento médio","estilo":"mais silencioso"}'

# 4. Calcular match (requer token)
curl -X POST http://localhost:8080/matches \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"perfilId":1,"sessaoId":1}'

# 5. Ver matches aprovados do perfil 1 (aberto)
curl http://localhost:8080/perfis/1/matches
```

---

## 9. Pesquisa e Contextualização

### 9.1 Contexto e Motivação

A aprendizagem colaborativa é amplamente reconhecida na literatura educacional como uma estratégia eficaz para o desenvolvimento de competências técnicas e interpessoais. Quando estudantes trabalham juntos, tendem a consolidar o conhecimento com mais profundidade, identificar lacunas no próprio raciocínio e desenvolver habilidades de comunicação.

Contudo, a formação espontânea de grupos esbarra em um problema prático: a dificuldade de encontrar parceiros com perfis compatíveis. Incompatibilidades de nível (avançado com iniciante) ou de estilo (silencioso com debatedor) comprometem a experiência de todos.

### 9.2 Matchmaking Aplicado à Educação

Sistemas de _matchmaking_ — algoritmos de compatibilidade — são amplamente usados em recrutamento (LinkedIn), relacionamentos (Tinder) e jogos online. O princípio é atribuir pesos a critérios relevantes e calcular uma pontuação. Aqui, o algoritmo foi adaptado para o domínio educacional com três critérios: disciplina, nível e estilo. O threshold de 60 pontos exige compatibilidade em pelo menos dois dos três critérios, filtrando matches superficiais.

### 9.3 Escolha Tecnológica

**Go + Gin**: compilação nativa com alto throughput, tipagem estática e framework HTTP eficiente, adequado para APIs REST com múltiplas requisições simultâneas.

**MongoDB + mongo-driver/v2**: banco NoSQL orientado a documentos. As structs Go mapeiam diretamente para documentos BSON sem camadas de ORM relacionais. O driver oficial `mongo-driver/v2` atua como ODM via tags `bson`, oferecendo operações como `FindOne`, `Find`, `InsertOne`, `FindOneAndUpdate` e `DeleteOne` com tipagem forte.

**JWT (JSON Web Token)**: padrão stateless de autenticação. O servidor não armazena sessões — cada requisição carrega o próprio token assinado, que é validado pelo middleware sem consulta ao banco. Isso mantém a API alinhada com os princípios REST de ausência de estado (_stateless_).

A arquitetura em camadas (Router → Controller → Service → Repository) isola o acesso ao banco na camada Repository, facilitando a manutenção e possibilitando trocar o banco sem afetar as demais camadas.

### 9.4 Possíveis Evoluções

- Endpoint de sugestão automática: dado um `perfilId`, retornar as sessões com maior score
- Critério adicional de compatibilidade de horário
- Notificação ao criador da sessão quando um match aprovado ocorrer
- Refresh token para renovação automática da sessão
