package notebook

import (
	"context"
	"core/internal/domain/model/notebook"
)

type NotebookService interface {
	ReadNotebooks(ctx context.Context) (notebook.Notebooks, error)
	WriteNotebook(ctx context.Context, aNotebook notebook.Notebook) (notebook.Notebook, error)
	DeleteNotebook(ctx context.Context, notebookId string) error
}
