package note

import (
	"context"
	"core/internal/domain/model/note"
	noterepository "core/internal/domain/repository/note"
	noteservice "core/internal/domain/service/note"
)

type NoteService struct {
	noteRepository noterepository.NoteRepository
}

// WriteNote implements note.NoteService.
func (n *NoteService) WriteNote(ctx context.Context, note note.Note) (note.Note, error) {
	note, err := n.noteRepository.UpSertNote(ctx, note)
	note.GetTitle()

	return note, err
}

// ReadNotes implements note.NoteService.
func (n *NoteService) ReadNotes(ctx context.Context, notebookId string) (note.Notes, error) {
	notes, err := n.noteRepository.ReadNotes(ctx, notebookId)
	return notes, err
}

func New(
	noteRepository noterepository.NoteRepository,
) noteservice.NoteService {
	return &NoteService{
		noteRepository: noteRepository,
	}
}
