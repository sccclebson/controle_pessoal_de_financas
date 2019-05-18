package helper

import (
	"fmt"
	"testing"
	"time"
)

func TestGetSenhaSha256(t *testing.T) {
	senha := "123456"
	senhaSHA256 := "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92"

	if GetSenhaSha256(senha) != senhaSHA256 {
		t.Error("HASH sha256 gerado inválido", GetSenhaSha256(senha))
	}
}

func TestDataFormatada(t *testing.T) {
	data := time.Date(2012, 12, 27, 23, 57, 33, 0, new(time.Location))
	dataString := "27/12/2012 23:57:33"

	if DataFormatada(data) != dataString {
		t.Error("Retorno de DataFormatada inválida")
	}
}

func TestFormatarPorta(t *testing.T) {
	porta := "8080"
	portaFormatada := FormatarPorta(porta)

	if portaFormatada != fmt.Sprintf(":%s", porta) {
		t.Error("Retorno de função For")
	}
}

func TestGetDiretorioAbs(t *testing.T) {
	diretorio, err := GetDiretorioAbs()
	if err != nil {
		t.Error(err, diretorio)
	}
}
