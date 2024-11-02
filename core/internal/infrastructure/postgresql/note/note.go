package note

import (
	"context"
	"core/internal/domain/model/note"
	noterepository "core/internal/domain/repository/note"
	"core/pkg/customerror"
	"core/pkg/errorcode"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/lib/pq" // PostgreSQL driver
)

type NoteRepository struct {
	db           *sql.DB
	queryBuilder sq.StatementBuilderType
	table        string
}

// DeleteNotes implements note.NoteRepository.
func (n *NoteRepository) DeleteNote(ctx context.Context, noteId uuid.UUID) error {
	query := n.queryBuilder.Delete(n.table).Where(sq.Eq{"id": noteId.String()})
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
		err = customerror.NewBusinessError(fmt.Errorf("no row found with the specified ID: %s", noteId.String()), customerror.Opts{Code: errorcode.NotFoundError})
		return err
	}
	return nil
}

// ReadNotes implements note.NoteRepository.
func (n *NoteRepository) ReadNotes(ctx context.Context, notebookId uuid.UUID) (note.Notes, error) {

	query := n.queryBuilder.Select("id", "notebook_id", "content", "updated_at").
		From(n.table).
		Where(sq.Eq{"notebook_id": notebookId.String()})

	sql, args, err := query.ToSql()
	if err != nil {
		err = customerror.NewSystemError(fmt.Errorf("fail to build sql: %v", err), customerror.Opts{Code: errorcode.DatabaseError})
		return note.Notes{}, err
	}

	rows, err := n.db.Query(sql, args...)
	if err != nil {
		err = customerror.NewSystemError(fmt.Errorf("error executing sql: %v", err), customerror.Opts{Code: errorcode.DatabaseError})
		return note.Notes{}, err
	}
	defer rows.Close()

	notes := note.Notes{}
	for rows.Next() {
		var tmpID, tmpNotebookId, tmpContent, tmpUpdatedAt string
		if err := rows.Scan(&tmpID, &tmpNotebookId, &tmpContent, &tmpUpdatedAt); err != nil {
			err = customerror.NewSystemError(fmt.Errorf("error scanning row: %v", err), customerror.Opts{Code: errorcode.DatabaseError})
			return note.Notes{}, err
		}

		uuidID, err := uuid.Parse(tmpID)
		if err != nil {
			return note.Notes{}, err
		}

		uuidNotebookId, err := uuid.Parse(tmpNotebookId)
		if err != nil {
			return note.Notes{}, err
		}

		timeUpdatedAt, err := time.Parse(time.RFC3339, tmpUpdatedAt)
		if err != nil {
			return note.Notes{}, err
		}

		aNote := note.Note{
			ID:         uuidID,
			NotebookId: uuidNotebookId,
			Content:    tmpContent,
			UpdatedAt:  timeUpdatedAt,
		}
		notes.Notes = append(notes.Notes, aNote)
	}

	if err := rows.Err(); err != nil {
		err = customerror.NewSystemError(fmt.Errorf("error with rows iteration: %v", err), customerror.Opts{Code: errorcode.DatabaseError})
		return note.Notes{}, err
	}

	return notes, nil
}

// WriteNote implements note.NoteRepository.
func (n *NoteRepository) UpSertNote(ctx context.Context, aNote note.Note) (note.Note, error) {

	query := n.queryBuilder.Insert(n.table).
		Columns("id", "notebook_id", "content", "updated_at").
		Values(aNote.GetID().String(), aNote.NotebookId.String(), aNote.Content, aNote.GetUpdatedAt().Format(time.RFC3339)).
		Suffix("ON CONFLICT (id) DO UPDATE SET content = EXCLUDED.content, updated_at = EXCLUDED.updated_at")

	sql, args, err := query.ToSql()
	if err != nil {
		err = customerror.NewSystemError(fmt.Errorf("fail to build sql: %v", err), customerror.Opts{Code: errorcode.DatabaseError})
		return note.Note{}, err
	}

	_, err = n.db.Exec(sql, args...)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			err = customerror.NewBusinessError(fmt.Errorf("notebook_id not found"), customerror.Opts{Code: errorcode.NotFoundError})
			return note.Note{}, err
		}

		err = customerror.NewSystemError(fmt.Errorf("error executing sql: %v", err), customerror.Opts{Code: errorcode.DatabaseError})
		return note.Note{}, err
	}

	return aNote, nil
}

func New(
	db *sql.DB,
	table string,
) noterepository.NoteRepository {
	return &NoteRepository{
		db:           db,
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		table:        table,
	}
}
