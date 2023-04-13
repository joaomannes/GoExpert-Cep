package ApiCep

import (
	"encoding/json"
	"github.com/joaomannes/GoExpert-Cep/internal/Cep"
	"net/http"
)

type ApiCepResult struct {
	Code     string `json:"code"`
	State    string `json:"state"`
	City     string `json:"city"`
	District string `json:"district"`
	Address  string `json:"address"`
}

func (acr *ApiCepResult) ToCepResult() *Cep.CepResult {
	return &Cep.CepResult{
		Source:     "ApiCep",
		Cep:        acr.Code,
		Uf:         acr.State,
		Localidade: acr.City,
		Bairro:     acr.District,
		Logradouro: acr.Address,
	}
}

type ApiCepSearcher struct{}

func (acs *ApiCepSearcher) Search(cep string, ch chan *Cep.CepResult) {
	//time.Sleep(time.Second * 2)
	req, err := http.NewRequest("GET", "https://cdn.apicep.com/file/apicep/"+cep+".json", nil)
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

	var result ApiCepResult
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return
	}

	ch <- result.ToCepResult()
}
