package main

import (
	"fmt"
	"os"
	"strings"

	"project/help"
	"project/methods"
)

const logo = `
███  ████  ███ █████ █████  ████ █████
█   █ █   █  █    █   █     █       █
█████ ████   █    █   ████   ███    █
█   █ █      █    █   █         █   █
█   █ █     ███   █   █████ ████    █
`

func main() {
	fmt.Println(logo)

	if len(os.Args) < 2 {
		help.Menu()
		return
	}

	comando := strings.ToLower(os.Args[1])

	if comando == "-h" || comando == "--help" {
		help.Menu()
		return
	}

	if len(os.Args) < 3 {
		fmt.Println("Erro: informe a URL.")
		fmt.Println("Use 'apitest --help' para ver os comandos.")
		return
	}

	url := os.Args[2]

	switch comando {
	case "get":
		methods.Get(url)
	case "post":
		if len(os.Args) >= 5 && os.Args[3] == "--body" {
			body := os.Args[4]
			methods.Post(url, body)
		} else {
			fmt.Println("Erro: POST precisa de --body")
		}
	case "put":
		methods.Put(url)

	case "delete":
		methods.Delete(url)

	default:
		fmt.Println("Erro: comando inválido:", comando)
		fmt.Println("Use 'apitest --help' para ver os comandos.")
	}
}
