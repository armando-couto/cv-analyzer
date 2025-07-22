package services

import (
	"github.com/ledongthuc/pdf"
	"io"
	"strings"
)

type Resultado struct {
	Arquivo string
	Nota    float64
	Resumo  string
}

func LerPDF(path string) (string, error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var texto strings.Builder
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	_, err = io.Copy(&texto, b)
	if err != nil {
		return "", err
	}
	return texto.String(), nil
}
