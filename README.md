# 📂 Currículo Analyzer com IA (Go + OpenAI)

Este projeto tem como objetivo analisar centenas ou milhares de currículos em PDF utilizando inteligência artificial (OpenAI GPT-4), extraindo informações relevantes de cada candidato e classificando os melhores com base em critérios definidos.

---

## 🚀 Funcionalidades

- 📥 Leitura automatizada de currículos em PDF
- 🤖 Uso da OpenAI para extrair:
    - Cidade e estado
    - Idade
    - Se está fazendo faculdade
    - Linguagens de programação mencionadas
    - Projetos relevantes
- 🏆 Classificação dos currículos por nota
- 📝 Geração de relatório com os 20 melhores candidatos

---

## 🛠️ Tecnologias utilizadas

- [Go (Golang)](https://golang.org)
- [OpenAI API (GPT-4)](https://platform.openai.com)
- [`github.com/ledongthuc/pdf`](https://pkg.go.dev/github.com/ledongthuc/pdf) – para leitura de PDFs
- [`github.com/sashabaranov/go-openai`](https://pkg.go.dev/github.com/sashabaranov/go-openai) – cliente OpenAI
- [`github.com/joho/godotenv`](https://pkg.go.dev/github.com/joho/godotenv) – para carregar variáveis do `.env`

---

## 📦 Instalação

```bash
git clone https://github.com/seu-usuario/curriculo-analyzer.git
cd curriculo-analyzer
go mod tidy
