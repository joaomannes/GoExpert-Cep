package ViaCep

import (
	"encoding/json"
	"github.com/joaomannes/GoExpert-Cep/internal/Cep"
	"net/http"
)

type ViaCepResult struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
}

func (vcr *ViaCepResult) ToCepResult() *Cep.CepResult {
	return &Cep.CepResult{
		Source:     "ViaCep",
		Cep:        vcr.Cep,
		Uf:         vcr.Uf,
		Localidade: vcr.Localidade,
		Bairro:     vcr.Bairro,
		Logradouro: vcr.Logradouro,
	}
}

type ViaCepSearcher struct{}

func (vcs *ViaCepSearcher) Search(cep string, ch chan *Cep.CepResult) {
	//time.Sleep(time.Second * 2)
	req, err := http.NewRequest("GET", "https://viacep.com.br/ws/"+cep+"/json/", nil)
	if err != nil {
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return
	}

	var result ViaCepResult
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return
	}

	ch <- result.ToCepResult()
}
