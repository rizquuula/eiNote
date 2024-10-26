package notebook

import (
	"context"
	"stub/internal/domain/model/notebook"
	notebookservice "stub/internal/domain/service/notebook"
	"stub/internal/test/mock"
)

type notebookService struct{}

// ReadNotebooks implements notebook.NotebookService.
func (n notebookService) ReadNotebooks(ctx context.Context, userId string) (notebook.Notebooks, error) {
	return mock.NotebooksMock, nil
}

func New() notebookservice.NotebookService {
	return notebookService{}
}
