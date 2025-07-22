package main

import (
	"cv-analyzer/services"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	// Carrega variáveis do .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("Defina a variável de ambiente OPENAI_API_KEY com sua chave da OpenAI.")
	}

	client := openai.NewClient(apiKey)

	caminho := os.Getenv("FILES_PATH")
	if caminho == "" {
		log.Fatal("Defina a variável de ambiente FILES_PATH.")
	}

	var resultados []services.Resultado

	err = filepath.Walk(caminho, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".pdf") {
			return nil
		}
		fmt.Println("Lendo:", info.Name())
		texto, err := services.LerPDF(path)
		if err != nil {
			log.Println("Erro lendo PDF:", err)
			return nil
		}

		nota, resumo, err := services.AvaliarCurriculo(client, texto)
		if err != nil {
			log.Println("Erro na avaliação:", err)
			return nil
		}

		resultados = append(resultados, services.Resultado{
			Arquivo: info.Name(),
			Nota:    nota,
			Resumo:  resumo,
		})
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// Ordenar por nota (desc)
	sort.Slice(resultados, func(i, j int) bool {
		return resultados[i].Nota > resultados[j].Nota
	})

	// Criar arquivo de relatório
	relatorio, err := os.Create("relatorio_top20.txt")
	if err != nil {
		log.Fatal("Erro criando o relatório:", err)
	}
	defer relatorio.Close()

	// Mostrar os top 20 e salvar no arquivo
	fmt.Println("\nTop 20 Currículos:\n")
	limite := 20
	if len(resultados) < 20 {
		limite = len(resultados)
	}
	for i := 0; i < limite; i++ {
		saida := fmt.Sprintf("%02d. %s — Nota: %.1f\nResumo: %s\n\n",
			i+1,
			resultados[i].Arquivo,
			resultados[i].Nota,
			resultados[i].Resumo,
		)

		fmt.Print(saida) // Imprime no terminal
		_, err := relatorio.WriteString(saida)
		if err != nil {
			log.Println("Erro ao escrever no relatório:", err)
		}
	}

	fmt.Println("Relatório salvo em: relatorio_top20.txt")
}
