# Match dos Estudos

> API REST em Go para conectar estudantes com sessões de estudo compatíveis com seu perfil.

---

## 1. Descrição do Domínio

### Problema

Estudantes que desejam estudar em grupo frequentemente enfrentam dificuldades para encontrar sessões com disciplina, nível de conhecimento e estilo de aprendizagem compatíveis. Essa falta de conexão reduz a produtividade e desmotiva a aprendizagem colaborativa.

### Solução

O **Match dos Estudos** é uma API REST que permite cadastrar perfis de estudantes e sessões de estudo, e então calcular automaticamente a compatibilidade entre eles através de um algoritmo de pontuação (_score_). Um perfil recebe um match **aprovado** quando atinge pontuação suficiente com uma sessão.

### Entidades do Domínio

| Entidade | Descrição                                                                                                              |
| -------- | ---------------------------------------------------------------------------------------------------------------------- |
| `Perfil` | Representa as preferências e características acadêmicas de um estudante: disciplina, nível e estilo de estudo          |
| `Sessao` | Define uma sessão de estudo organizada, com título, disciplina, nível, estilo, data/hora, duração e vagas              |
| `Match`  | Representa o resultado da compatibilidade calculada entre um perfil e uma sessão, contendo score e status de aprovação |

### Algoritmo de Match

O score é calculado comparando os campos de um `Perfil` com os de uma `Sessao`. O match é **aprovado** quando o score atinge **60 pontos ou mais**.

| Critério                    | Pontos  |
| --------------------------- | ------- |
| Mesma disciplina            | 40      |
| Mesmo nível de conhecimento | 30      |
| Mesmo estilo de estudo      | 30      |
| **Total máximo**            | **100** |

---

## 2. Instalação e Execução Local

### Pré-requisitos

- [Go 1.21+](https://go.dev/dl/) instalado
- Git instalado

### Passos

```bash
# 1. Clone o repositório
git clone https://github.com/emiliaferlin/Match-dos-Estudos.git
cd Match-dos-Estudos

# 2. Instale as dependências
go mod tidy

# 3. Execute o servidor
go run ./src/main.go
```

O servidor iniciará em `http://localhost:8080`.

---

## 3. Tabela de Rotas da API

### 3.1 Perfis

| Método   | Rota          | Descrição                         | Status de sucesso |
| -------- | ------------- | --------------------------------- | ----------------- |
| `GET`    | `/perfis`     | Lista todos os perfis cadastrados | 200               |
| `POST`   | `/perfis`     | Cria um novo perfil               | 201               |
| `PUT`    | `/perfis/:id` | Atualiza um perfil existente      | 200               |
| `DELETE` | `/perfis/:id` | Remove um perfil pelo ID          | 200               |

**Exemplo de body — POST/PUT `/perfis`:**

```json
{
  "nome": "Ana Lima",
  "idade": 22,
  "disciplina": "Algoritmos",
  "nivel": "conhecimento médio",
  "estilo": "gosta de argumentar"
}
```

---

### 3.2 Sessões

| Método   | Rota           | Descrição                          | Status de sucesso |
| -------- | -------------- | ---------------------------------- | ----------------- |
| `GET`    | `/sessoes`     | Lista todas as sessões disponíveis | 200               |
| `POST`   | `/sessoes`     | Cria uma nova sessão de estudo     | 201               |
| `PUT`    | `/sessoes/:id` | Atualiza uma sessão existente      | 200               |
| `DELETE` | `/sessoes/:id` | Remove uma sessão pelo ID          | 200               |

**Exemplo de body — POST/PUT `/sessoes`:**

```json
{
  "titulo": "Revisão de Algoritmos",
  "disciplina": "Algoritmos",
  "nivel": "conhecimento médio",
  "estilo": "gosta de argumentar",
  "dataHoraInicio": "2026-04-08T19:00:00",
  "duracaoMinutos": 120,
  "vagas": 5
}
```

---

### 3.3 Matches

| Método | Rota                  | Descrição                                                           | Status de sucesso |
| ------ | --------------------- | ------------------------------------------------------------------- | ----------------- |
| `POST` | `/matches`            | Calcula o score entre um perfil e uma sessão e retorna o resultado  | 201               |
| `GET`  | `/perfis/:id/matches` | Retorna todos os matches **aprovados** do perfil com o ID informado | 200               |

**Exemplo de body — POST `/matches`:**

```json
{
  "perfilId": 1,
  "sessaoId": 2
}
```

**Exemplo de resposta — POST `/matches` (aprovado):**

```json
{
  "id": 1,
  "perfilId": 1,
  "sessaoId": 2,
  "score": 70,
  "aprovado": true
}
```

**Exemplo de resposta — POST `/matches` (reprovado):**

```json
{
  "id": 2,
  "perfilId": 1,
  "sessaoId": 3,
  "score": 40,
  "aprovado": false
}
```

**Exemplo de resposta — GET `/perfis/1/matches`:**

```json
[
  {
    "id": 1,
    "perfilId": 1,
    "sessaoId": 2,
    "score": 70,
    "aprovado": true
  }
]
```

---

## 4. Estrutura do Projeto

```
Match-dos-Estudos/
├── src/
│   ├── main.go              # Ponto de entrada — inicia o servidor Gin
│   ├── router/
│   │   └── router.go        # Registro de todas as rotas e injeção de dependências
│   ├── controller/
│   │   └── controller.go    # Handlers HTTP (Perfil, Sessao, Match)
│   ├── service/
│   │   └── service.go       # Regras de negócio e algoritmo de score do match
│   ├── repository/
│   │   └── repository.go    # Persistência em memória (Perfil, Sessao, Match)
│   └── model/
│       └── model.go         # Structs: Perfil, Sessao, Match
├── go.mod
├── go.sum
└── README.md
```

---

## 5. Exemplos de Uso (fluxo completo)

```bash
# 1. Criar um perfil
curl -X POST http://localhost:8080/perfis \
  -H "Content-Type: application/json" \
  -d '{"nome":"Ana","idade":22,"disciplina":"Algoritmos","nivel":"conhecimento médio","estilo":"gosta de argumentar"}'

# 2. Criar uma sessão
curl -X POST http://localhost:8080/sessoes \
  -H "Content-Type: application/json" \
  -d '{"titulo":"Revisão","disciplina":"Algoritmos","nivel":"conhecimento médio","estilo":"gosta de argumentar","dataHoraInicio":"2026-04-10T19:00:00","duracaoMinutos":90,"vagas":4}'

# 3. Calcular o match (perfilId=1, sessaoId=1)
curl -X POST http://localhost:8080/matches \
  -H "Content-Type: application/json" \
  -d '{"perfilId":1,"sessaoId":1}'

# 4. Ver matches aprovados do perfil 1
curl http://localhost:8080/perfis/1/matches
```

---

## 6. Pesquisa e Contextualização

### 6.1 Contexto e Motivação

A aprendizagem colaborativa é amplamente reconhecida na literatura educacional como uma estratégia eficaz para o desenvolvimento de competências técnicas e interpessoais. Quando estudantes trabalham juntos em problemas desafiadores, tendem a consolidar o conhecimento com mais profundidade, identificar lacunas no próprio raciocínio e desenvolver habilidades de comunicação e argumentação.

Contudo, a formação espontânea de grupos de estudo esbarra em um problema prático: **a dificuldade de encontrar parceiros com perfis compatíveis**. Incompatibilidades de nível de conhecimento (um estudante avançado com um iniciante) ou de estilo (alguém que prefere silêncio com alguém que aprende debatendo) comprometem a experiência de ambos.

### 6.2 Matchmaking Aplicado à Educação

Sistemas de _matchmaking_ — algoritmos que calculam compatibilidade entre perfis — são amplamente utilizados em plataformas de recrutamento (LinkedIn, Gupy), relacionamentos (Tinder) e jogos online. O princípio central é o mesmo: atribuir pesos a critérios relevantes e calcular uma pontuação de compatibilidade.

Neste projeto, o algoritmo foi adaptado para o domínio educacional com três critérios principais: disciplina de interesse, nível de conhecimento e estilo de aprendizagem. O threshold de aprovação (60 pontos) foi definido para exigir compatibilidade em pelo menos dois dos três critérios, garantindo que matches superficiais sejam filtrados.

### 6.3 Escolha Tecnológica — Go + Gin

A linguagem **Go (Golang)** foi escolhida pelas seguintes características:

- **Performance**: compilação para código nativo com alto throughput em servidores HTTP, adequado para APIs com múltiplas requisições simultâneas.
- **Tipagem estática**: reduz erros em tempo de execução e facilita a manutenção.
- **Concorrência nativa**: o modelo de goroutines é eficiente para processar requisições paralelas sem overhead de threads.
- **Ecossistema maduro para APIs**: o framework **Gin** oferece roteamento eficiente, binding de JSON e middleware com sintaxe enxuta.

A arquitetura em camadas (Router → Controller → Service → Repository) foi adotada para separar responsabilidades, facilitando testes unitários de cada camada e a substituição futura da persistência em memória por um banco de dados real (como PostgreSQL ou MongoDB, já presente no `go.mod`).

### 6.4 Possíveis Evoluções

- Endpoint de sugestão automática: dado um perfilId, retornar as sessões com maior score
- Notificação ao criador da sessão quando um match aprovado ocorrer
