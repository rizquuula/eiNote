package notebook

import (
	"context"
	"core/internal/domain/model/notebook"
	notebookrepository "core/internal/domain/repository/notebook"
	notebookservice "core/internal/domain/service/notebook"
)

type notebookService struct {
	notebookRepository notebookrepository.NoteRepository
}

// ReadNotebooks implements notebook.NotebookService.
func (n notebookService) ReadNotebooks(ctx context.Context, userId string) (notebook.Notebooks, error) {
	return notebook.Notebooks{}, nil
}

func New(
	notebookRepository notebookrepository.NoteRepository,
) notebookservice.NotebookService {
	return notebookService{
		notebookRepository: notebookRepository,
	}
}
