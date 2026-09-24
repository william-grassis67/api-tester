package methods

import (
	"fmt"
	"net/http"
	"strings"
)

func Post(url string, body string) {
	resposta, err := http.Post(url, "application/json", strings.NewReader(body))

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	defer resposta.Body.Close()

	fmt.Println("Status:", resposta.Status)
}
