# 📂 Currículo Analyzer com IA (Go + OpenAI)

Este projeto analisa automaticamente milhares de currículos em PDF usando inteligência artificial (OpenAI GPT-4), extraindo informações relevantes dos candidatos e classificando os melhores com base em critérios definidos.

---

## 📚 Índice

- [Funcionalidades](#-funcionalidades)
- [Tecnologias utilizadas](#-tecnologias-utilizadas)
- [Configuração](#configuração)
- [Executando](#executando)
- [Exemplo de saída no relatório](#exemplo-de-saída-no-relatório)
- [Possíveis melhorias](#-possíveis-melhorias)
- [Licença](#-licença)

---

## 🚀 Funcionalidades

- 📥 Leitura automatizada de currículos em PDF
- 🤖 Extração via OpenAI das seguintes informações:
  - Cidade e estado
  - Idade aproximada
  - Se está cursando faculdade (e qual)
  - Linguagens de programação que domina
  - Projetos relevantes
- 🏆 Geração de nota/classificação
- 📝 Criação de relatório `.txt` com os 20 melhores currículos

---

## 🛠️ Tecnologias utilizadas

- [Go (Golang)](https://golang.org)
- [OpenAI GPT-4 API](https://platform.openai.com)
- [`github.com/ledongthuc/pdf`](https://pkg.go.dev/github.com/ledongthuc/pdf)
- [`github.com/sashabaranov/go-openai`](https://pkg.go.dev/github.com/sashabaranov/go-openai)
- [`github.com/joho/godotenv`](https://pkg.go.dev/github.com/joho/godotenv)

---

## ⚙️ Configuração

1. Clone este repositório:

```bash
git clone https://github.com/seu-usuario/curriculo-analyzer.git
cd curriculo-analyzer
go mod tidy
```

2. Crie um arquivo `.env` com sua chave da OpenAI:

```env
OPENAI_API_KEY=sk-sua-chave-aqui
FILES_PATH="./curriculos"
```

## ▶️ Executando

Execute o projeto com:

```bash
go run main.go
```

O sistema irá:

- Ler todos os currículos PDF da pasta `./curriculos`
- Enviar o conteúdo para o GPT-4 com um prompt personalizado
- Gerar um relatório chamado `relatorio_top20.txt` contendo os candidatos mais bem avaliados

---

## 📝 Exemplo de saída no relatório

```
01. joao_silva.pdf — Nota: 9.2
Resumo:
Cidade/Estado: São Paulo, SP
Idade: 22
Faculdade: Sim — Análise e Desenvolvimento de Sistemas na FIAP
Linguagens: Python, JavaScript, SQL
Projetos relevantes:
- Sistema de controle de estoque com Django
- API RESTful com Go + PostgreSQL

02. maria_oliveira.pdf — Nota: 8.8
Resumo:
Cidade/Estado: Belo Horizonte, MG
Idade: 20
Faculdade: Sim — Ciência da Computação na UFMG
Linguagens: Java, Kotlin
Projetos relevantes:
- Aplicativo de monitoramento de estudos
```

---

## ✅ Possíveis melhorias

- Exportar os dados em formato CSV ou JSON
- Rodar em paralelo para melhor performance
- Filtrar candidatos por localização, linguagem ou idade
- Criar uma interface web para upload e visualização dos currículos

---