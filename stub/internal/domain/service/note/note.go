package note

import (
	"context"
	"stub/internal/domain/model/note"
)

type NoteService interface {
	ReadNote(ctx context.Context, noteId string, userId string) (note.Note, error)
	ReadNotes(ctx context.Context, userId string) (note.Notes, error)
	WriteNote(ctx context.Context, note note.Note) error
}
