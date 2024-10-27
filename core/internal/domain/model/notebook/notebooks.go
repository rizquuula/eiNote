package notebook

import "encoding/json"

type Notebooks struct {
	Notebooks []Notebook
}

func (n Notebooks) MarshalJSON() ([]byte, error) {
	if n.Notebooks == nil {
		n.Notebooks = []Notebook{}
	}

	return json.Marshal(&struct {
		Notebooks []Notebook `json:"notebooks"`
	}{
		Notebooks: n.Notebooks,
	})
}
