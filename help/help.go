package help

import "fmt"

func Menu() {
	fmt.Println(`
API Tester CLI

Usage:
  apitest [command] [options]

Commands:
  get       Envia uma requisição GET
  post      Envia uma requisição POST
  put       Envia uma requisição PUT
  delete    Envia uma requisição DELETE

Options:
  -h, --help       Mostra esta ajuda
  -v, --version    Mostra a versão
		`)

	return
}
