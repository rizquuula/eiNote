package note

import (
	"context"
	"core/internal/domain/model/note"

	"github.com/google/uuid"
)

type NoteRepository interface {
	ReadNotes(ctx context.Context, notebookId uuid.UUID) (note.Notes, error)
	UpSertNote(ctx context.Context, aNote note.Note) (note.Note, error)
	DeleteNote(ctx context.Context, noteId uuid.UUID) error
}
