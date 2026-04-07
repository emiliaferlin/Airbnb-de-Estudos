# Plataforma de Estudos Colaborativos(Match dos Estudos)

## 1. Problema
Pessoas desejam estudar em conjunto, mas enfrentam dificuldades para encontrar parceiros com interesses, níveis e horários compatíveis.

---

## 2. Solução
A plataforma permite conectar usuários com base em critérios como:

- Disciplina
- Nível de conhecimento
- Disponibilidade de horário
- Estilo de estudo (silencioso ou colaborativo)

---

## 3. Modelo de Domínio

### 3.1 PerfilEstudo
Representa as preferências e características acadêmicas do usuário.

### 3.2 SessaoEstudo
Define sessões organizadas onde usuários podem estudar em conjunto, com horário, duração e regras específicas.

### 3.3 Match
Representa a compatibilidade entre usuários com base em seus perfis.

---

## 4. Interface REST

### 4.1 Perfis

| Método | Rota            | Descrição                  |
|--------|-----------------|----------------------------|
| GET    | /perfis         | Lista todos os perfis      |
| POST   | /perfis         | Cria um novo perfil        |
| PUT    | /perfis/{id}    | Atualiza um perfil         |
| DELETE | /perfis/{id}    | Remove um perfil           |

---

### 4.2 Sessões

| Método | Rota             | Descrição                        |
|--------|------------------|----------------------------------|
| GET    | /sessoes         | Lista sessões disponíveis        |
| POST   | /sessoes         | Cria uma nova sessão             |
| PUT    | /sessoes/{id}    | Atualiza uma sessão              |
| DELETE | /sessoes/{id}    | Remove uma sessão                |

---

### 4.3 Matches

| Método | Rota      | Descrição                          |
|--------|----------|------------------------------------|
| GET    | /matches | Lista usuários compatíveis          |

---

## 5. Representação de Dados

### Exemplo de SessaoEstudo

```json
{
  "id": 1,
  "titulo": "Estrutura de Dados",
  "disciplina": "Algoritmos",
  "nivel": "intermediario",
  "estilo": "colaborativo",
  "dataHoraInicio": "2026-04-08T19:00:00",
  "duracaoMinutos": 120,
  "vagas": 5
}
