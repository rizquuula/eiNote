package notebook

import (
	"context"
	"core/internal/domain/model/notebook"
	notebookrepository "core/internal/domain/repository/notebook"
	notebookservice "core/internal/domain/service/notebook"
	"core/pkg/customerror"
	"core/pkg/errorcode"
	"fmt"

	"github.com/google/uuid"
)

type notebookService struct {
	notebookRepository notebookrepository.NotebookRepository
}

// DeleteNotebook implements notebook.NotebookService.
func (n notebookService) DeleteNotebook(ctx context.Context, notebookId string) error {
	notebookUUID, err := uuid.Parse(notebookId)
	if err != nil {
		err = customerror.NewBusinessError(fmt.Errorf("error to parse notebook id: %v", err), customerror.Opts{Code: errorcode.RequestParsingError})
		return err
	}

	err = n.notebookRepository.DeleteNotebook(ctx, notebookUUID)
	return err
}

// WriteNotebook implements notebook.NotebookService.
func (n notebookService) WriteNotebook(ctx context.Context, aNotebook notebook.Notebook) (notebook.Notebook, error) {
	newNotebook, err := n.notebookRepository.UpSertNotebook(ctx, aNotebook)
	if err != nil {
		return notebook.Notebook{}, err
	}
	return newNotebook, nil
}

// ReadNotebooks implements notebook.NotebookService.
func (n notebookService) ReadNotebooks(ctx context.Context) (notebook.Notebooks, error) {
	notebooks, err := n.notebookRepository.ReadNotebooks(ctx)
	return notebooks, err
}

func New(
	notebookRepository notebookrepository.NotebookRepository,
) notebookservice.NotebookService {
	return notebookService{
		notebookRepository: notebookRepository,
	}
}
