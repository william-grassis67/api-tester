package methods

import (
	"fmt"
	"io"
	"net/http"
)

func Get(url string) {
	resposta, err := http.Get(url)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer resposta.Body.Close()

	body, err := io.ReadAll(resposta.Body)

	if err != nil {
		fmt.Println("Erro ao ler a resposta: ", err)
		return
	}

	fmt.Println(string(body))
}
