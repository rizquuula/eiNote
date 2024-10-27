package note

import (
	"context"
	"core/internal/domain/model/note"
	noterepository "core/internal/domain/repository/note"
	"database/sql"
)

type NoteRepository struct {
	db *sql.DB
}

// ReadNotes implements note.NoteRepository.
func (n *NoteRepository) ReadNotes(ctx context.Context, notebookId string) (note.Notes, error) {
	panic("unimplemented")
}

// WriteNote implements note.NoteRepository.
func (n *NoteRepository) UpSertNote(ctx context.Context, note note.Note) (note.Note, error) {
	// query := sq.Insert("note").
	//
	//	Columns("username", "email").
	//	Values(username, email).
	//	Suffix("RETURNING id")
	panic("unimplemented")
}

func New(
	db *sql.DB,
) noterepository.NoteRepository {
	return &NoteRepository{
		db: db,
	}
}
