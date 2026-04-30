package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Endereco struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
	Erro       bool   `json:"erro"`
}

func BuscarCep(cep string) (*Endereco, error) {
	// 1. Configurar cliente com Timeout para evitar goroutines penduradas
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("falha na requisição: %v", err)
	}
	defer resp.Body.Close() // Importante para evitar vazamento de memória

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code inesperado: %d", resp.StatusCode)
	}

	// 3. Decodificar o JSON diretamente do stream do corpo (mais performático que ler tudo para string)
	var addr Endereco
	if err := json.NewDecoder(resp.Body).Decode(&addr); err != nil {
		return nil, fmt.Errorf("erro ao decodificar json: %v", err)
	}

	if addr.Erro {
		return nil, fmt.Errorf("CEP %s não encontrado", cep)
	}

	return &addr, nil
}

func main() {
	cep := "01001000" //Exemplo: Praça da Sé, SP
	fmt.Printf("Buscando CEP: %s...\n", cep)

	addr, err := BuscarCep(cep)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	fmt.Printf("Endereço: %s, %s - %s/%s\n", addr.Logradouro, addr.Bairro, addr.Localidade, addr.Uf)
}
