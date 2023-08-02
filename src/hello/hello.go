package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3

func main() {

	for {
		exibeIntroducao()
		exibeMenu()
		comando := leComando()

		nome, idade := devolveNomeEIdade()
		fmt.Println(nome, "e", idade)

		_, idades := devolveNomeEIdade()
		fmt.Println(nome, "e", idades)

		/* if comando == 1 {
			fmt.Println("Monitorando")
		} else if comando == 2 {
			fmt.Println("Exibindo logs")
		} else if comando == 0 {
			fmt.Println("Saindo do programa")
		} else {
			fmt.Println("Não conheço o comando")
		} */

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço o comando")
			os.Exit(-1)
		}
	}

}

func devolveNomeEIdade() (string, int) {
	nome := "Victor"
	idade := 24
	return nome, idade
}

func exibeIntroducao() {
	nome := "Victor"
	idade := 24
	versao := 1.1
	fmt.Println("Ola sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("O tipo da variavel versao é", reflect.TypeOf(versao))
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)

	fmt.Println("O endereço da minha variavel comando é", &comando)
	fmt.Println("o comando escolhido foi", comando)

	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando")
	sites := []string{"https://www.alura.com.br", "https://www.globo.com.br", "https://www.uol.com.br"}
	/* var sites [4]string
	sites[0] = "https://www.alura.com.br"
	sites[1] = "https://www.globo.com.br"
	sites[2] = "https://www.uol.com.br" */

	fmt.Println(sites)

	for i := 0; i < len(sites); i++ {
		fmt.Println(sites[i])
	}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Estou passando na posição", i+1, "do meu slice e essa posição tem o site", site)
			testaSite(site)
		}
		time.Sleep(monitoramentos * time.Second)
		fmt.Println("")
	}
	fmt.Println("")

	sites2 := lerSitesArquivo()

	fmt.Println(sites2)

	exibeNomes()

}

func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Erro ao acessar o site:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!!!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func lerSitesArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu o erro", err)
		return sites
	}
	defer arquivo.Close()

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		fmt.Println(linha)

		if err != nil && err != io.EOF {
			fmt.Println("Ocorreu o erro", err)
			return sites
		}

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + "-" + site + "- Online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

}

// Para saber mais: funções que retornam arrays
func devolveEstadosDoSudeste() [4]string {
	var estados [4]string
	estados[0] = "RJ"
	estados[1] = "SP"
	estados[2] = "MG"
	estados[3] = "ES"
	return estados
}

func exibeNomes() {
	nomes := []string{"Douglas", "Victor"}
	nomes = append(nomes, "Aparecida")
	fmt.Println(nomes)
	fmt.Println("O slice tem, ", len(nomes))
	fmt.Println("teste aquii!!! ", cap(nomes))
}
