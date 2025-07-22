package services

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"strings"
)

func AvaliarCurriculo(client *openai.Client, texto string) (float64, string, error) {
	prompt := fmt.Sprintf(`
Extraia as seguintes informações do currículo abaixo. Seja objetivo e baseado apenas no que estiver escrito.

1. Cidade e estado (se houver)
2. Idade aproximada (ou data de nascimento, se presente)
3. Está fazendo faculdade? (Sim/Não + curso e instituição se houver)
4. Quais linguagens de programação sabe (baseado no que menciona)
5. Projetos relevantes (descreva brevemente, se possível)

Formato de saída esperado:

Cidade/Estado: <cidade>, <estado>
Idade: <idade>
Faculdade: <Sim/Não + detalhes>
Linguagens: <lista de linguagens>
Projetos relevantes:
- <projeto 1>
- <projeto 2>
...

Currículo:
%s
`, texto)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "user",
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		return 0, "", err
	}

	content := resp.Choices[0].Message.Content

	// Pegar nota (esperado: "Nota: 8.5")
	var nota float64
	var resumo string
	fmt.Sscanf(content, "Nota: %f", &nota)

	// Tentar extrair o resumo também
	parts := strings.SplitN(content, "Resumo:", 2)
	if len(parts) == 2 {
		resumo = strings.TrimSpace(parts[1])
	}

	return nota, resumo, nil
}
