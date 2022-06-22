package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	nomes := leituraText("Massa de Dados.txt")
	tratamento := tratamentoLogin(nomes)
	apresentar(nomes, tratamento)
}

func leituraText(pathFile string) []string {
	// Abre o file
	file, err := os.Open(pathFile)
	// Caso tenha encontrado algum erro ao tentar abrir o file retorne o erro encontrado
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo, %s", err)
	}
	// Garante que o file sera fechado apos o uso
	defer file.Close()

	// Cria um scanner que le cada linha do file
	var linhas []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linhas = append(linhas, scanner.Text())
	}

	// Retorna as linhas lidas
	return linhas
}

func tratamentoLogin(nomes []string) []string {
	var logins []string
	for _, nomeCompleto := range nomes {
		nomeValidado, err := verificaLetra(nomeCompleto)
		if err != nil {
			log.Fatalf("Erro na verificação de letras A-Z, %s", err)
		}


		// nome sendo transformado em caracteres maiúsculos
		nomeUpper := strings.ToUpper(nomeValidado)

		// Fatiando nome
		nomeSplit := strings.SplitAfter(nomeUpper, " ")
		tamanhoDoSliceNome := len(nomeSplit)

		// removendo elementos de ligacao
		for i, pedacoNome := range nomeSplit {
			if pedacoNome == "DE" || pedacoNome == "DA" {
				nomeSplit[i] = nomeSplit[tamanhoDoSliceNome-1]
				nomeSplit[tamanhoDoSliceNome-1] = ""
				nomeSplit = nomeSplit[:tamanhoDoSliceNome-1]
			}
		}

		// pegando primeiro e ultimo nome
		primeiroNome := nomeSplit[0]
		ultimoNome := nomeSplit[len(nomeSplit)-1]

		// juntando as primeiras 4 letras com as ultimas 3
		primeiroNomeQuatroLetras := primeiroNome[0:4]
		ultimoNomeTresLetras := ultimoNome[0:3]
		login := primeiroNomeQuatroLetras + ultimoNomeTresLetras

		// comparando logins
		for _, v := range logins {
			if login == v {
				primeiroNomeTresLetras := primeiroNome[0:3]
				ultimoNomeQuatroLetras := ultimoNome[0:4]
				login = ultimoNomeQuatroLetras + primeiroNomeTresLetras
			}
		}

		logins = append(logins, login)

	}
	return logins
}

func apresentar(nomes []string, login []string) {
	for i, nome := range nomes {
		fmt.Println("NOME: ", nome+" " + "LOGIN: ", login[i])
	}
}

func verificaLetra(nomeCompleto string) (string, error) {
	// verifica A-Z
	reg, err := regexp.Compile("^[A-Za-zÀ-ÖØ-öø-ÿ]+$")
	if err != nil {
		log.Fatal(err)
	}
	nomeValidado := reg.ReplaceAllString(nomeCompleto, "")
	return nomeValidado, err
}