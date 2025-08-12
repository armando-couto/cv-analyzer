# Extrator de Informações de Currículos em PDF com Go + Ollama

Este projeto em Go lê vários arquivos PDF de currículos de uma pasta local, extrai o texto de cada um e envia para um modelo local do Ollama para extrair informações específicas conforme um prompt fixo. A resposta da IA é exibida no terminal para cada currículo processado.

---

## Funcionalidades

- Lê todos os PDFs da pasta `pdfs` local.
- Extrai texto dos PDFs usando a biblioteca `github.com/ledongthuc/pdf`.
- Envia o texto extraído junto com um prompt fixo para o modelo local Ollama (`llama3`).
- Recebe e exibe as informações extraídas de cada currículo no terminal.
- Fácil de adaptar para outros prompts ou modelos.

---

## Requisitos

- Go 1.18+ instalado
- Ollama instalado e rodando localmente (https://ollama.com/download)
- Modelo `llama3` baixado no Ollama (exemplo usado no código)

---

## Instalação e Setup

1. **Instale o Ollama**

   Baixe e instale o Ollama para seu sistema operacional:  
   [https://ollama.com/download](https://ollama.com/download)

2. **Baixe o modelo `llama3`**

   No terminal, execute:
   ```bash
   ollama pull llama3
   ```

3. **Clone este repositório ou crie um diretório**

   Coloque o código `main.go` dentro do diretório.

4. **Crie a pasta `pdfs`**

   Coloque os arquivos PDF dos currículos dentro dessa pasta.

5. **Instale a dependência de PDF**

   No terminal, dentro do diretório do projeto:
   ```bash
   go get github.com/ledongthuc/pdf
   ```

---

## Uso

Execute o programa:
```bash
go run main.go
```

Para cada PDF na pasta `pdfs`, o programa:

- Extrairá o texto.
- Enviará o texto junto com o prompt fixo para o Ollama.
- Imprimirá a resposta formatada da IA no terminal.

---

## Prompt Utilizado

```txt
Extraia as seguintes informações do currículo enviado. Seja direto e utilize apenas os dados presentes no texto. Não faça suposições.

1. Cidade e estado (caso estejam informados)
2. Idade aproximada ou data de nascimento (caso esteja no currículo)
3. Está cursando faculdade? (Responda Sim ou Não + curso e instituição, se mencionados)
4. Linguagens de programação ou ferramentas que domina (com base no que está escrito)
5. Projetos relevantes (resuma brevemente até 2, se houver)
6. Perfil mais adequado: vaga de desenvolvimento ou teste de software

Formato de saída esperado:

Cidade/Estado: <cidade>, <estado>
Idade: <idade ou "não informado">
Faculdade: <Sim/Não> – <curso>, <instituição> (se houver)
Tipo da vaga: <Desenvolvimento/Teste>
Linguagens: <linguagens mencionadas>
Projetos relevantes:
- <projeto 1>
- <projeto 2>
...

Currículo:
```

---

## Considerações

- O Ollama deve estar rodando localmente (`ollama serve`) para aceitar requisições HTTP.
- O modelo `llama3` pode ser substituído por outro modelo disponível no Ollama, basta ajustar no código.
- O processamento é feito um PDF por vez. Pode ser adaptado para processamento paralelo.
- Caso tenha muitos PDFs, fique atento ao limite de tokens do modelo.

---

## Sugestões para melhorias

- Salvar resultados em arquivo CSV ou JSON.
- Criar interface web simples para upload de PDFs.
- Suporte a outras línguas e prompts customizados.
- Melhor tratamento de erros e logs detalhados.

---

## Licença

MIT License — Use livremente, adapte para seu projeto!

---

Se quiser, posso te ajudar a criar essas melhorias também! Quer seguir?
