package Cep

type CepResult struct {
	Source     string `json:"source"`
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

type Searcher interface {
	Search(cep string, ch chan *CepResult)
}
