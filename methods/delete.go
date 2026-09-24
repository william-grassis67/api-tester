package methods

import (
	"fmt"
	"net/http"
)

func Delete(url string) {
	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	cliente := &http.Client{}

	resposta, err := cliente.Do(req)

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	defer resposta.Body.Close()

	fmt.Println("Status:", resposta.Status)
}
