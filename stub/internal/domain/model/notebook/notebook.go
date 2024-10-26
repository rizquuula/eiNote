package notebook

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Notebook struct {
	ID   uuid.UUID
	Name string
}

func (n Notebook) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}{
		ID:   n.ID.String(),
		Name: n.Name,
	})
}
