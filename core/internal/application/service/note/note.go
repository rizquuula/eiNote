package note

import (
	"context"
	"core/internal/domain/model/note"
	noteservice "core/internal/domain/service/note"
)

type noteService struct {
}

// WriteNote implements note.NoteService.
func (n noteService) WriteNote(ctx context.Context, note note.Note) error {
	return nil
}

// ReadNotes implements note.NoteService.
func (n noteService) ReadNotes(ctx context.Context, notebookId string, userId string) (note.Notes, error) {
	return note.Notes{}, nil
}

func New() noteservice.NoteService {
	return noteService{}
}
