package main

import (
	"bufio"
	"csvexplorer/Filter"
	"csvexplorer/GeneralFuncs"
	"csvexplorer/Helper"
	"os"
	"strconv"
	"strings"
)

var (
	Arquivo string = ""
)

const (
	Banner string = `
  ____ ______     __              _                     
 / ___/ ___\ \   / /____  ___ __ | | ___  _ __ ___ _ __ 
| |   \___ \\ \ / / _ \ \/ / '_ \| |/ _ \| '__/ _ \ '__|
| |___ ___) |\ V /  __/>  <| |_) | | (_) | | |  __/ |   
 \____|____/  \_/ \___/_/\_\ .__/|_|\___/|_|  \___|_|   
  by: guicapelleto         |_|
`
)

func defineParametros(param, sec string) {
	info := strings.ReplaceAll(sec, "\\/", " ")
	switch param {
	case "--arquivo", "-a":
		Arquivo = info
		Filter.Informacoes = append(Filter.Informacoes, "Arquivo: "+info)
	case "--limite", "-l":
		limite, err := strconv.Atoi(info)
		GeneralFuncs.ErroSaida(err)
		Filter.QuantidadeLinhas = limite
		Filter.Informacoes = append(Filter.Informacoes, "Limite de resultados: "+info)
	case "--ler-coluna", "-lc":
		var colunas []string
		for _, coluna := range strings.Split(info, ",") {
			if strings.Contains(coluna, "-") {
				i, err_i := strconv.Atoi(strings.Split(coluna, "-")[0])
				f, err_f := strconv.Atoi(strings.Split(coluna, "-")[1])
				GeneralFuncs.ErroSaida(err_i)
				GeneralFuncs.ErroSaida(err_f)
				if i < f {
					for n := range f + 1 {
						if n >= i {
							Filter.FiltroColuna = append(Filter.FiltroColuna, n)
						}
					}
				} else {
					for n := range i + 1 {
						if n >= f {
							Filter.FiltroColuna = append(Filter.FiltroColuna, n)
						}
					}
				}
			} else {
				n, err := strconv.Atoi(coluna)
				GeneralFuncs.ErroSaida(err)
				Filter.FiltroColuna = append(Filter.FiltroColuna, n)
			}
		}
		Filter.FiltrarColuna = true
		for _, val := range Filter.FiltroColuna {
			colunas = append(colunas, strconv.Itoa(val))
		}
		Filter.Informacoes = append(Filter.Informacoes, "Filtro de colunas: "+strings.Join(colunas, " "))
	case "--ler-linha", "-ll":
		var linhas []string
		for _, linha := range strings.Split(info, ",") {
			if strings.Contains(linha, "-") {
				i, err_i := strconv.Atoi(strings.Split(linha, "-")[0])
				f, err_f := strconv.Atoi(strings.Split(linha, "-")[1])
				GeneralFuncs.ErroSaida(err_i)
				GeneralFuncs.ErroSaida(err_f)
				if i < f {
					for n := range f + 1 {
						if n >= i {
							Filter.FiltroLinha = append(Filter.FiltroLinha, n)
						}
					}
				} else {
					for n := range i + 1 {
						if n >= f {
							Filter.FiltroLinha = append(Filter.FiltroLinha, n)
						}
					}
				}
			} else {
				n, err := strconv.Atoi(linha)
				GeneralFuncs.ErroSaida(err)
				Filter.FiltroLinha = append(Filter.FiltroLinha, n)
			}
		}
		Filter.FiltrarLinha = true
		for _, val := range Filter.FiltroLinha {
			linhas = append(linhas, strconv.Itoa(val))
		}
		Filter.Informacoes = append(Filter.Informacoes, "Filtro de linhas: "+strings.Join(linhas, ","))
	case "--filtrar", "-f":
		Filter.FiltroTexto = strings.Split(info, ",")
		Filter.FiltrarTexto = true
		Filter.Informacoes = append(Filter.Informacoes, "Filtro de texto: "+strings.Join(Filter.FiltroTexto, " "))
	case "--separador", "-s":
		Filter.Separador = info
		Filter.Informacoes = append(Filter.Informacoes, "Separador modificado: "+info)
	case "--break", "-b":
		Filter.BreakLine = true
		Filter.Informacoes = append(Filter.Informacoes, "Quebra de campos ativo")
	case "--contagem", "-C":
		Filter.Contagem = true
		Filter.Informacoes = append(Filter.Informacoes, "Contagem de linhas ativo")
	case "--verbose", "-v":
		Filter.Verbose = true
		Filter.Informacoes = append(Filter.Informacoes, "Verbose ativa")
	case "--unicos", "-u":
		Filter.Unicos = true
		Filter.Informacoes = append(Filter.Informacoes, "Resultados únicos ativo")
	case "--filtrarColuna", "-fc":
		Filter.FiltrarTextoColuna = true
		separador := strings.Split(info, ":")
		coluna, err := strconv.Atoi(separador[0])
		GeneralFuncs.ErroSaida(err)
		atributos := strings.Split(separador[1], ",")
		fcol := Filter.FiltroTextoColuna{Coluna: coluna}
		for _, item := range atributos {
			if !(item == "") {
				fcol.Filtros = append(fcol.Filtros, item)
			}
		}
		Filter.FiltrosColuna = append(Filter.FiltrosColuna, fcol)
		mensagem := "Filtrar textos (" + strings.Join(fcol.Filtros, ",") + ") na coluna " + separador[0]
		Filter.Informacoes = append(Filter.Informacoes, mensagem)
	case "--silenciar", "-S":
		Filter.Silenciado = true
		Filter.Informacoes = append(Filter.Informacoes, "Resultados silenciados ativo")
	case "--help", "-h":
		Helper.MostrarAjuda()
	}
}

func pegarParametros() {
	var param, info string

	for _, arg := range os.Args[0:] {
		if strings.Contains(arg, "=") {
			param, info = strings.Split(arg, "=")[0], strings.Split(arg, "=")[1]
		} else {
			param = arg
			info = ""
		}
		defineParametros(param, info)

	}
}

func checarArquivo() {
	if Arquivo == "" {
		GeneralFuncs.GerarErro("Parâmetro de arquivo de leitura não foi encontrado.")
	}
	_, err := os.Stat(Arquivo)
	GeneralFuncs.ErroSaida(err)
}

func gerarLeitura() {
	file, err := os.Open(Arquivo)
	GeneralFuncs.ErroSaida(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	counter := 0
	for scanner.Scan() {
		counter++
		line := scanner.Text()
		Filter.Filtrar(line, counter)

	}
}

func init() {
	Helper.ModoUso = os.Args[0] + " [opções]"
	Helper.CadastrarHelper("--arquivo", "-a", "=C:\\pasta\\/01\\arquivo.csv", "Atribui um arquivo para leitura. Obs: usar \\/ como espaço.")
	Helper.CadastrarHelper("--limite", "-l", "=4", "Limita o número dos resultados encontrados.")
	Helper.CadastrarHelper("--ler-coluna", "-lc", "=1,6-8", "Limita a exibição dos resultados encontrados para as colunas mencionadas.")
	Helper.CadastrarHelper("--ler-linha", "-ll", "=1,6-8", "Limita a exibição dos resultados encontrados para as linhas mencionadas.")
	Helper.CadastrarHelper("--filtrar", "-f", "=palavra1,palavra2", "Realiza um filtro de texto, onde somente será exibido caso o texto mencionado for encontrado na linha.")
	Helper.CadastrarHelper("--filtrarColuna", "-fc", "=7:palavra1,palavra2", "Realiza um filtro de texto, onde somente será exibido caso o texto mencionado for encontrado na linha.")
	Helper.CadastrarHelper("--separador", "-s", "=;", "Altera o padrão de vírgula como separação de campos por outro, no exemplo ponto e vírgula.")
	Helper.CadastrarHelper("--break", "-b", "", "Quebra a exibição das colunas em linhas, ideal para ver separadamente os elementos de uma linha específica.")
	Helper.CadastrarHelper("--verbose", "-v", "", "Ativa informações adicionais.")
	Helper.CadastrarHelper("--contagem", "-c", "", "Exibe o total de linhas obtidas.")
	Helper.CadastrarHelper("--unicos", "-u", "", "Não exibe linhas repetidas.")
	Helper.CadastrarHelper("--silenciar", "-S", "", "Não exibe o resultado das linhas.")
	Helper.CadastrarHelper("--help", "-h", "", "Exibe esse menu de ajuda.")
}

func main() {
	GeneralFuncs.MostrarBanner(Banner, "1.5")
	pegarParametros()
	checarArquivo()
	Filter.MostrarColunas()
	gerarLeitura()
	if Filter.Contagem {
		Filter.MostrarTotal()
	}

}
