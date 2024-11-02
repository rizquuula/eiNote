package notebook

import (
	"context"
	"core/internal/domain/model/notebook"

	"github.com/google/uuid"
)

type NotebookRepository interface {
	ReadNotebooks(ctx context.Context) (notebook.Notebooks, error)
	UpSertNotebook(ctx context.Context, aNotebook notebook.Notebook) (notebook.Notebook, error)
	DeleteNotebook(ctx context.Context, notebookId uuid.UUID) error
}
