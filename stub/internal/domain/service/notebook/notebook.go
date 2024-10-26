package notebook

import (
	"context"
	"stub/internal/domain/model/notebook"
)

type NotebookService interface {
	ReadNotebooks(ctx context.Context, userId string) (notebook.Notebooks, error)
}
