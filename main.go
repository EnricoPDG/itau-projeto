package main

import (
	"bufio" 	// O pacote bufio ajuda com I/O em buffer. Através de vários exemplos: Reader, Writer e Scanner…
	"fmt"		// Output na linha de comando
	"log"		// Log dos erros
	"os"		// Trás funcionalidades de sistema operacional
	"regexp"	// Expressões regulares
	"strings"	// Manipulação de strings
)

func main() {
	nomes := leituraText("Massa de Dados.txt")
	logins := tratamentoLogin(nomes)
	apresentar(nomes, logins)
}

func leituraText(pathFile string) []string {
	// Abre o arquivo
	file, err := os.Open(pathFile)

	// Caso tenha encontrado algum erro ao tentar abrir o arquivo retorne o erro encontrado
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo, %s", err)
	}

	// Garante que o arquivo sera fechado após o uso
	defer file.Close()

	// Cria um scanner que le cada linha do arquivo
	var linhas []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linhas = append(linhas, scanner.Text())
	}

	// Retorna as linhas lidas
	return linhas
}

func tratamentoLogin(nomes []string) []string {
	var logins []string 					// Criando variavel para armazenar o slice de logins
	for _, nomeCompleto := range nomes {	// Laço for range para percorrer todos os elementos
		// Verificação se a string apenas tem caracteres de A-Z
		nomeValidado, err := verificaLetra(nomeCompleto)
		if err != nil {
			log.Fatalf("Erro na verificação de letras A-Z, %s", err)
		} 

		// Nome sendo transformado em caracteres maiúsculos
		nomeUpper := strings.ToUpper(nomeValidado)

		// Fatiando nome
		nomeSplit := strings.SplitAfter(nomeUpper, " ")

		// Removendo elementos de ligacao
		nomeSplit = removeElementosLigacao(nomeSplit)

		// Pegando primeiro e ultimo nome
		primeiroNome := nomeSplit[0]
		ultimoNome := nomeSplit[len(nomeSplit) - 1]

		// Juntando as primeiras 4 letras com as ultimas 3
		primeiroNomeQuatroLetras := primeiroNome[0:4]
		ultimoNomeTresLetras := ultimoNome[0:3]
		login := primeiroNomeQuatroLetras + ultimoNomeTresLetras

		// Comparando logins
		for _, v := range logins {
			if login == v {
				primeiroNomeTresLetras := primeiroNome[0:3]
				ultimoNomeQuatroLetras := ultimoNome[0:4]
				login = ultimoNomeQuatroLetras + primeiroNomeTresLetras
			}
		}

		// Adiciona o login no slice logins
		logins = append(logins, login)

	}

	return logins
}

func apresentar(nomes []string, logins []string) {
	for i, nome := range nomes{
		fmt.Printf("%d Colaborador\n", i+1)
		fmt.Printf("Nome: %s \nLogin: %s \n", nome, logins[i])
		fmt.Println("------------------------------------")
	}
}

func verificaLetra(nomeCompleto string) (string, error) {
	// Uso de Expressão regular para vetificar se o nome tem apenas caracteres de A-Z
	reg, err := regexp.Compile("^[A-Za-zÀ-ÖØ-öø-ÿ]+$") 
	if err != nil {
		log.Fatal(err)
	}
	nomeValidado := reg.ReplaceAllString(nomeCompleto, "") // Caso não seja letra ele retira neste momento
	return nomeValidado, err
}

func removeElementosLigacao(nomeSplit []string) []string{
	tamanhoDoSliceNome := len(nomeSplit) // Tamanho do nome por exemplo {"ENRICO, PAPSCH, DI, GIACOMO"} = 4
	for i, pedacoNome := range nomeSplit {
		if pedacoNome == "DE" || pedacoNome == "DA" {
			nomeSplit[i] = nomeSplit[tamanhoDoSliceNome - 1] // Elemento fica na última posição
			nomeSplit[tamanhoDoSliceNome - 1] = ""			 // Remove o elemento da última posição	
			nomeSplit = nomeSplit[:len(nomeSplit) - 1]   	 //	Corta o slice	
		}
	}
	return nomeSplit
}