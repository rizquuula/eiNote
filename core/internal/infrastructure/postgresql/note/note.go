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
)

type NoteRepository struct {
	db           *sql.DB
	queryBuilder sq.StatementBuilderType
}

// ReadNotes implements note.NoteRepository.
func (n *NoteRepository) ReadNotes(ctx context.Context, notebookId string) (note.Notes, error) {

	query := n.queryBuilder.Select("id", "notebook_id", "content", "updated_at").
		From("note").
		Where(sq.Eq{"notebook_id": notebookId})

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

	query := n.queryBuilder.Insert("note").
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
		err = customerror.NewSystemError(fmt.Errorf("error executing sql: %v", err), customerror.Opts{Code: errorcode.DatabaseError})
		return note.Note{}, err
	}

	return aNote, nil
}

func New(
	db *sql.DB,
) noterepository.NoteRepository {
	return &NoteRepository{
		db:           db,
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}
