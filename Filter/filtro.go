package Filter

import (
	"csvexplorer/GeneralFuncs"
	"os"
	"strconv"
	"strings"
)

type FiltroTextoColuna struct {
	Coluna  int
	Filtros []string
}

var (
	QuantidadeAtual int = 0
	Resultados      []string
)

var (
	QuantidadeLinhas   int  = 0
	FiltrarColuna      bool = false
	FiltroColuna       []int
	FiltrarLinha       bool = false
	FiltroLinha        []int
	FiltroTexto        []string
	FiltrarTexto       bool   = false
	Separador          string = ","
	BreakLine          bool   = false
	Informacoes        []string
	Verbose            bool = false
	Contagem           bool = false
	ValorContagem      int  = 0
	Unicos             bool = false
	FiltrarTextoColuna bool = false
	FiltrosColuna      []FiltroTextoColuna
	Silenciado         bool = false
)

func MostrarColunas() {
	buf := ""
	apn := "Col:  "
	GeneralFuncs.Println("Opções atribuídas: \n"+strings.Join(Informacoes, "\n")+"\n", "green")
	if FiltrarColuna {
		for _, col := range FiltroColuna {
			//apn = buf
			buf = strconv.Itoa(col)
			apn += buf + " | "
		}
		GeneralFuncs.Println(apn, "blue")
	}
}

func MostrarTotal() {
	GeneralFuncs.Println("Total de resultados encontrados: "+strconv.Itoa(ValorContagem), "green")
}

func Filtrar(linha string, contador int) {

	var resultado string = linha
	if FiltrarTexto {
		mark := false
		for _, texto := range FiltroTexto {
			if strings.Contains(linha, texto) {
				mark = true
				break
			}
		}
		if !mark {
			return
		}
	}
	if FiltrarTextoColuna {
		mark := 0
		elementos := strings.Split(linha, Separador)
		for _, listafiltros := range FiltrosColuna {
			coluna := listafiltros.Coluna
			filtros := listafiltros.Filtros
			if GeneralFuncs.SliceStrContains(filtros, elementos[coluna]) {
				mark++
			}
		}
		if mark != len(FiltrosColuna) {
			return
		}
	}
	if FiltrarLinha {

		if !GeneralFuncs.SliceIntContains(FiltroLinha, contador) {
			return
		}
	}
	if FiltrarColuna {
		var buf []string
		elementos := strings.Split(linha, Separador)
		for n, coluna := range elementos {
			if GeneralFuncs.SliceIntContains(FiltroColuna, n) {
				buf = append(buf, coluna)
			}
		}
		resultado = strings.Join(buf, Separador+" ")

	}
	if BreakLine {
		elementos := strings.Split(resultado, Separador)
		resultado = ""
		for n, elemento := range elementos {
			if Verbose {
				resultado += strconv.Itoa(n) + " " + elemento + "\n"
			} else {
				resultado += elemento + "\n"
			}
		}
	}
	if Unicos {
		if GeneralFuncs.SliceStrContains(Resultados, resultado) {
			return
		} else {
			Resultados = append(Resultados, resultado)
		}

	}
	if Verbose {
		GeneralFuncs.Println("Linha: "+strconv.Itoa(contador), "blue")
	}
	if !(Silenciado) {
		GeneralFuncs.Println(resultado, "yellow")
	}
	if Contagem {
		ValorContagem++
	}
	if QuantidadeLinhas != 0 {
		if QuantidadeAtual >= QuantidadeLinhas-1 {
			if Contagem {
				MostrarTotal()
			}
			os.Exit(0)
		}
		QuantidadeAtual++
	}
}
