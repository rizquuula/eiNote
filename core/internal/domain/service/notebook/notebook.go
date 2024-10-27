package notebook

import (
	"context"
	"core/internal/domain/model/notebook"
)

type NotebookService interface {
	ReadNotebooks(ctx context.Context, userId string) (notebook.Notebooks, error)
}
