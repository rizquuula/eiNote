package note

import "encoding/json"

type Notes struct {
	Notes []Note
}

func (n Notes) MarshalJSON() ([]byte, error) {
	if n.Notes == nil {
		n.Notes = []Note{}
	}

	return json.Marshal(&struct {
		Notes []Note `json:"notes"`
	}{
		Notes: n.Notes,
	})
}
