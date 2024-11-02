package note

import (
	"context"
	"core/internal/domain/model/note"
	noterepository "core/internal/domain/repository/note"
	noteservice "core/internal/domain/service/note"
	"core/pkg/customerror"
	"core/pkg/errorcode"
	"fmt"

	"github.com/google/uuid"
)

type NoteService struct {
	noteRepository noterepository.NoteRepository
}

// DeleteNote implements note.NoteService.
func (n *NoteService) DeleteNote(ctx context.Context, noteId string) error {
	noteUUID, err := uuid.Parse(noteId)
	if err != nil {
		err = customerror.NewBusinessError(fmt.Errorf("error to parse notebook id: %v", err), customerror.Opts{Code: errorcode.RequestParsingError})
		return err
	}
	err = n.noteRepository.DeleteNote(ctx, noteUUID)
	return err
}

// WriteNote implements note.NoteService.
func (n *NoteService) WriteNote(ctx context.Context, note note.Note) (note.Note, error) {
	note, err := n.noteRepository.UpSertNote(ctx, note)
	note.GetTitle()

	return note, err
}

// ReadNotes implements note.NoteService.
func (n *NoteService) ReadNotes(ctx context.Context, notebookId string) (note.Notes, error) {
	if notebookId == "" {
		return note.Notes{}, nil
	}

	notebookUUID, err := uuid.Parse(notebookId)
	if err != nil {
		err = customerror.NewBusinessError(fmt.Errorf("error to parse notebook id: %v", err), customerror.Opts{Code: errorcode.RequestParsingError})
		return note.Notes{}, err
	}

	notes, err := n.noteRepository.ReadNotes(ctx, notebookUUID)
	return notes, err
}

func New(
	noteRepository noterepository.NoteRepository,
) noteservice.NoteService {
	return &NoteService{
		noteRepository: noteRepository,
	}
}
