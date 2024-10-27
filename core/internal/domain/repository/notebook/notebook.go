package notebook

import (
	"context"
	"core/internal/domain/model/notebook"
)

type NoteRepository interface {
	ReadNotebooks(ctx context.Context) (notebook.Notebooks, error)
}
