package note

import (
	"encoding/json"
	"time"
)

type Note struct {
	Content   string
	UpdatedAt time.Time
}

func (n Note) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Content   string `json:"content"`
		UpdatedAt string `json:"updated_at"`
	}{
		Content:   n.Content,
		UpdatedAt: n.UpdatedAt.Format("2006-01-02 15:04:05"),
	})
}
