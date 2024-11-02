package notebook

import (
	"context"
	"core/internal/domain/model/notebook"
	notebookrepository "core/internal/domain/repository/notebook"
	"core/pkg/customerror"
	"core/pkg/errorcode"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

var notebookCols = []string{"id", "name", "updated_at"}

type NotebookRepository struct {
	db           *sql.DB
	queryBuilder sq.StatementBuilderType
	table        string
}

// DeleteNotebook implements notebook.NotebookRepository.
func (n *NotebookRepository) DeleteNotebook(ctx context.Context, notebookId uuid.UUID) error {
	query := n.queryBuilder.Delete(n.table).Where(sq.Eq{"id": notebookId.String()})
	sql, args, err := query.ToSql()
	if err != nil {
		err = customerror.NewSystemError(fmt.Errorf("fail to build sql: %v", err), customerror.Opts{Code: errorcode.DatabaseError})
		return err
	}

	result, err := n.db.Exec(sql, args...)
	if err != nil {
		err = customerror.NewSystemError(fmt.Errorf("error executing sql: %v", err), customerror.Opts{Code: errorcode.DatabaseError})
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		err = customerror.NewSystemError(fmt.Errorf("error fetching rows affected: %v", err), customerror.Opts{Code: errorcode.DatabaseError})
		return err
	}

	if rowsAffected == 0 {
		err = customerror.NewBusinessError(fmt.Errorf("no row found with the specified ID: %s", notebookId.String()), customerror.Opts{Code: errorcode.NotFoundError})
		return err
	}
	return nil
}

// upSertNotebook implements notebook.NotebookRepository.
func (n *NotebookRepository) UpSertNotebook(ctx context.Context, aNotebook notebook.Notebook) (notebook.Notebook, error) {
	if aNotebook.ID == uuid.Nil {
		aNotebook.GenerateID()
	}

	query := n.queryBuilder.Insert(n.table).
		Columns(notebookCols...).
		Values(aNotebook.ID.String(), aNotebook.Title, aNotebook.GetUpdatedAt().Format(time.RFC3339)).
		Suffix("ON CONFLICT (id) DO UPDATE SET name = EXCLUDED.name, updated_at = EXCLUDED.updated_at")

	sql, args, err := query.ToSql()
	if err != nil {
		err = customerror.NewSystemError(fmt.Errorf("fail to build sql: %v", err), customerror.Opts{Code: errorcode.DatabaseError})
		return notebook.Notebook{}, err
	}

	_, err = n.db.Exec(sql, args...)
	if err != nil {
		err = customerror.NewSystemError(fmt.Errorf("error executing sql: %v", err), customerror.Opts{Code: errorcode.DatabaseError})
		return notebook.Notebook{}, err
	}

	return aNotebook, nil

}

// ReadNotebooks implements notebook.NoteRepository.
func (n *NotebookRepository) ReadNotebooks(ctx context.Context) (notebook.Notebooks, error) {

	query := n.queryBuilder.Select(notebookCols...).From(n.table)

	sql, args, err := query.ToSql()
	if err != nil {
		err = customerror.NewSystemError(fmt.Errorf("fail to build sql: %v", err), customerror.Opts{Code: errorcode.DatabaseError})
		return notebook.Notebooks{}, err
	}

	rows, err := n.db.Query(sql, args...)
	if err != nil {
		err = customerror.NewSystemError(fmt.Errorf("error executing sql: %v", err), customerror.Opts{Code: errorcode.DatabaseError})
		return notebook.Notebooks{}, err
	}
	defer rows.Close()

	notebooks := notebook.Notebooks{}
	for rows.Next() {
		var tmpNotebookId, tmpTitle, tmpUpdatedAt string
		if err := rows.Scan(&tmpNotebookId, &tmpTitle, &tmpUpdatedAt); err != nil {
			err = customerror.NewSystemError(fmt.Errorf("error scanning row: %v", err), customerror.Opts{Code: errorcode.DatabaseError})
			return notebook.Notebooks{}, err
		}

		uuidNotebookId, err := uuid.Parse(tmpNotebookId)
		if err != nil {
			return notebook.Notebooks{}, err
		}

		timeUpdatedAt, err := time.Parse(time.RFC3339, tmpUpdatedAt)
		if err != nil {
			return notebook.Notebooks{}, err
		}

		aNotebook := notebook.Notebook{
			ID:        uuidNotebookId,
			Title:     tmpTitle,
			UpdatedAt: timeUpdatedAt,
		}
		notebooks.Notebooks = append(notebooks.Notebooks, aNotebook)
	}
	return notebooks, nil
}

func New(
	db *sql.DB,
	table string,
) notebookrepository.NotebookRepository {
	return &NotebookRepository{
		db:           db,
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		table:        table,
	}
}
