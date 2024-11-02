package GeneralFuncs

//versão 1.3
import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

func Println(textorecv, cor string) {
	texto := textorecv + "\n"
	switch cor {
	case "blue":
		color.Blue(texto)
	case "red":
		color.Red(texto)
	case "green":
		color.Green(texto)
	case "yellow":
		color.Yellow(texto)
	case "magenta":
		color.Magenta(texto)
	}
	color.Unset()
}

func Print(texto, cor string) {
	switch cor {
	case "blue":
		color.Blue(texto)
	case "red":
		color.Red(texto)
	case "green":
		color.Green(texto)
	case "yellow":
		color.Yellow(texto)
	case "magenta":
		color.Magenta(texto)
	}
	color.Unset()
}

func GetInput(msg string) (retorno string) {
	buf := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	ret, _ := buf.ReadString('\n')
	retorno = strings.Trim(ret, "\r\n")
	return
}

func MostrarErro(err error) bool {
	if err != nil {
		Println(err.Error(), "red")
		return true
	}
	return false
}

func ErroSaida(err error) {
	if err != nil {
		Println(err.Error(), "red")
		os.Exit(1)
	}
}

func CheckCSV(csvfile string) bool {
	if !strings.Contains(strings.ToLower(csvfile), ".csv") {
		return false
	}
	_, err := os.Stat(csvfile)
	if err != nil {
		return false
	}
	return true
}

func SliceStrContains(yourslice []string, elemento string) bool {
	for _, value := range yourslice {
		if value == elemento {
			return true
		}
	}
	return false
}

func SliceIntContains(yourslice []int, elemento int) bool {
	for _, value := range yourslice {
		if value == elemento {
			return true
		}
	}
	return false
}

func GerarErro(erro string) {
	Println(erro, "red")
	os.Exit(1)
}

func MostrarBanner(banner, versao string) {
	limpatela := strings.Repeat("\n", 5)
	Println(limpatela+banner, "blue")
	Println("Versão: "+versao+"\n", "red")
}
