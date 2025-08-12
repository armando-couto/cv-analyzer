package services

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const BASEP_ROMPT = `Extraia as informações abaixo do currículo exatamente conforme estão no texto.  
Se não encontrar a informação, escreva "não informado".  

1. Cidade e estado (exemplo: São Gonçalo, Rio de Janeiro)  
2. Idade aproximada ou data de nascimento  
3. Está cursando faculdade? (Sim ou Não + curso e instituição, se mencionados)  
4. Linguagens de programação ou ferramentas que domina  
5. Projetos relevantes (resuma até 2 projetos)  
6. Perfil mais adequado: vaga de desenvolvimento ou teste de software  

Formato da resposta:  
Cidade/Estado: <cidade>, <estado>  
Idade: <idade ou "não informado">  
Faculdade: <Sim/Não> – <curso>, <instituição> (se houver)  
Tipo da vaga: <Desenvolvimento/Teste>  
Linguagens: <linguagens mencionadas>  
Projetos relevantes:  
- <projeto 1>  
- <projeto 2>  

Exemplo:  
Cidade/Estado: São Gonçalo, Rio de Janeiro  
Idade: não informado  
Faculdade: Sim – Análise e Desenvolvimento de Sistemas, EBAC  
Tipo da vaga: Desenvolvimento  
Linguagens: Javascript, HTML, CSS, C#, React.js, Node.js  
Projetos relevantes:  
- Site de Game Shop: site interativo para venda de jogos digitais com formulários e animações em JavaScript  
- Blog Gastronômico: plataforma para recomendações de restaurantes, com cards interativos  

Currículo:  
`

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type OllamaChunk struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

func Ollama(text string, err error) strings.Builder {
	finalPrompt := BASEP_ROMPT + text

	reqBody, err := json.Marshal(OllamaRequest{
		Model:  "llama3",
		Prompt: finalPrompt,
	})
	if err != nil {
		fmt.Println("Erro ao montar JSON:", err)
		return strings.Builder{}
	}

	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewReader(reqBody))
	if err != nil {
		fmt.Println("Erro enviando para Ollama:", err)
		return strings.Builder{}
	}
	defer resp.Body.Close()

	var fullResponse strings.Builder
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		var chunk OllamaChunk
		err := json.Unmarshal([]byte(line), &chunk)
		if err != nil {
			fmt.Println("Erro decodificando chunk JSON:", err)
			continue
		}
		fullResponse.WriteString(chunk.Response)
		if chunk.Done {
			break
		}
	}
	return fullResponse
}
