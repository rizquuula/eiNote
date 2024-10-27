package notebook

import (
	"context"
	"core/internal/domain/model/notebook"
	notebookservice "core/internal/domain/service/notebook"
)

type notebookService struct{}

// ReadNotebooks implements notebook.NotebookService.
func (n notebookService) ReadNotebooks(ctx context.Context, userId string) (notebook.Notebooks, error) {
	return notebook.Notebooks{}, nil
}

func New() notebookservice.NotebookService {
	return notebookService{}
}
