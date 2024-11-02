package note

import (
	"context"
	"core/internal/domain/model/note"
)

type NoteService interface {
	ReadNotes(ctx context.Context, notebookId string) (note.Notes, error)
	WriteNote(ctx context.Context, note note.Note) (note.Note, error)
	DeleteNote(ctx context.Context, noteId string) error
}
