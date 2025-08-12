package main

import (
	"cv-analyzer/services"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Carrega variáveis do .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	pasta := os.Getenv("FILES_PATH")
	if pasta == "" {
		log.Fatal("Defina a variável de ambiente FILES_PATH.")
	}

	resultFile := "resultados.txt"

	err = os.WriteFile(resultFile, []byte(""), 0644)
	if err != nil {
		panic(err)
	}

	err = filepath.Walk(pasta, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".pdf") {
			fmt.Println("\n📄 Processando:", path)

			text, err := services.ExtractTextFromPDF(path)
			if err != nil {
				fmt.Println("Erro lendo PDF:", err)
				return nil
			}

			fullResponse := services.Ollama(text, err)
			resposta := fullResponse.String()

			// Salva no arquivo
			output := fmt.Sprintf("Arquivo: %s\n\n%s\n\n------------------------------------------------------------------------------------------\n\n", info.Name(), resposta)
			f, err := os.OpenFile(resultFile, os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("Erro abrindo arquivo de resultados:", err)
				return nil
			}
			defer f.Close()
			if _, err := f.WriteString(output); err != nil {
				fmt.Println("Erro escrevendo no arquivo de resultados:", err)
			}

			fmt.Println("\n=== Resposta da IA ===\n", resposta)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("✅ Processamento finalizado. Resultados salvos em %s\n", resultFile)
}
