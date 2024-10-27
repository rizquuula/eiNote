package note

import (
	"context"
	"stub/internal/domain/model/note"
	noteservice "stub/internal/domain/service/note"
	"stub/internal/test/mock"
)

type noteService struct {
}

// WriteNote implements note.NoteService.
func (n noteService) WriteNote(ctx context.Context, note note.Note) error {
	return nil
}

// ReadNotes implements note.NoteService.
func (n noteService) ReadNotes(ctx context.Context, notebookId string, userId string) (note.Notes, error) {
	return mock.NotesMock, nil
}

func New() noteservice.NoteService {
	return noteService{}
}
