package note

import (
	"context"
	"core/internal/domain/model/note"
)

type NoteRepository interface {
	ReadNotes(ctx context.Context, notebookId string) (note.Notes, error)
	UpSertNote(ctx context.Context, note note.Note) (note.Note, error)
}
