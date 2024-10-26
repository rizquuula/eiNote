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

// ReadNote implements note.NoteService.
func (n noteService) ReadNote(ctx context.Context, noteId string, userId string) (note.Note, error) {
	return mock.NoteMock, nil
}

// ReadNotes implements note.NoteService.
func (n noteService) ReadNotes(ctx context.Context, userId string) (note.Notes, error) {
	return mock.NotesMock, nil
}

func New() noteservice.NoteService {
	return noteService{}
}
