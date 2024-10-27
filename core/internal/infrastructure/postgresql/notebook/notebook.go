package notebook

import (
	"context"
	"core/internal/domain/model/notebook"
	notebookrepository "core/internal/domain/repository/notebook"
	"database/sql"
)

type NotebookRepository struct {
	db *sql.DB
}

// ReadNotebooks implements notebook.NoteRepository.
func (n *NotebookRepository) ReadNotebooks(ctx context.Context) (notebook.Notebooks, error) {
	panic("unimplemented")
}

func New(
	db *sql.DB,
) notebookrepository.NoteRepository {
	return &NotebookRepository{
		db: db,
	}
}
