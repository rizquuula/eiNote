package note

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID         uuid.UUID `json:"id"`
	NotebookId uuid.UUID `json:"notebook_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (n *Note) IdFromStr(id string) error {
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

func (n *Note) NotebookIdFromStr(id string) error {
	if id == "" {
		n.NotebookId = uuid.Nil
		return nil
	}

	idUuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	n.NotebookId = idUuid
	return nil
}

func (n *Note) GetID() uuid.UUID {
	if n.ID == uuid.Nil {
		n.ID = uuid.New()
	}
	return n.ID
}

func (n *Note) GetUpdatedAt() time.Time {
	if n.UpdatedAt.IsZero() {
		n.UpdatedAt = time.Now()
	}
	return n.UpdatedAt
}

func (n *Note) GetTitle() string {
	if n.Title == "" {
		titleCandidate := strings.Split(n.Content, "\n")[0]
		cleanTitle := strings.Replace(titleCandidate, "#", "", 1)
		trimmedTitle := strings.TrimSpace(cleanTitle)
		n.Title = trimmedTitle
	}
	return n.Title
}

func (n *Note) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID         string `json:"id"`
		NotebookId string `json:"notebook_id"`
		Title      string `json:"title"`
		Content    string `json:"content"`
		UpdatedAt  string `json:"updated_at"`
	}{
		ID:         n.ID.String(),
		NotebookId: n.NotebookId.String(),
		Title:      n.GetTitle(),
		Content:    n.Content,
		UpdatedAt:  n.UpdatedAt.Format(time.RFC3339),
	})
}
