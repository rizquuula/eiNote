package note

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID        uuid.UUID
	Title     string
	Content   string
	UpdatedAt time.Time
}

func (n Note) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID        string `json:"id"`
		Title     string `json:"title"`
		Content   string `json:"content"`
		UpdatedAt string `json:"updated_at"`
	}{
		ID:        n.ID.String(),
		Title:     n.Title,
		Content:   n.Content,
		UpdatedAt: n.UpdatedAt.Format(time.RFC3339),
	})
}
