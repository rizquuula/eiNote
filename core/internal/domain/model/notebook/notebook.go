package notebook

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Notebook struct {
	ID        uuid.UUID
	Name      string
	UpdatedAt time.Time
}

func (n *Notebook) GenerateID() {
	n.ID = uuid.New()
}

func (n *Notebook) GetUpdatedAt() time.Time {
	if n.UpdatedAt.IsZero() {
		n.UpdatedAt = time.Now()
	}
	return n.UpdatedAt
}

func (n *Notebook) NotebookIdFromStr(id string) error {
	if id == "" {
		n.ID = uuid.Nil
		return nil
	}

	idUuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	n.ID = idUuid
	return nil
}

func (n *Notebook) SetName(aName string) {
	n.Name = aName
}

func (n Notebook) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		UpdatedAt string `json:"updated_at"`
	}{
		ID:        n.ID.String(),
		Name:      n.Name,
		UpdatedAt: n.UpdatedAt.Format(time.RFC3339),
	})
}
