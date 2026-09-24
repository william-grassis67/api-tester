package methods

import (
	"fmt"
	"net/http"
)

func Put(url string) {
	req, err := http.NewRequest("PUT", url, nil)

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	client := &http.Client{}

	resposta, err := client.Do(req)

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	defer resposta.Body.Close()

	fmt.Println("Status:", resposta.Status)
}
