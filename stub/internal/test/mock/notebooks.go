package mock

import (
	"stub/internal/domain/model/notebook"

	"github.com/google/uuid"
)

var NotebooksMock = notebook.Notebooks{
	Notebooks: []notebook.Notebook{
		{
			ID:   uuid.New(),
			Name: "Private note",
		},
		{
			ID:   uuid.New(),
			Name: "Company X",
		},
		{
			ID:   uuid.New(),
			Name: "Freelance",
		},
	},
}
