package note

import (
	"encoding/json"
	"net/http"
	"stub/internal/domain/enum/errorcode"
	noteservice "stub/internal/domain/service/note"

	"stub/internal/domain/model/note"
	"stub/internal/interfaces/rest"
)

type noteController struct {
	noteService noteservice.NoteService
}

func (n *noteController) ReadNote(w http.ResponseWriter, r *http.Request) {
	userId := "test-user"
	queryParams := r.URL.Query()
	noteId := queryParams.Get("id")

	if noteId == "" {
		rest.NewBusinessError(w, errorcode.RequestParsingError, http.StatusUnprocessableEntity)
		return
	}

	noteResult, _ := n.noteService.ReadNote(r.Context(), noteId, userId)
	rest.NewResponse(w, noteResult)
}

func (n *noteController) ReadNotes(w http.ResponseWriter, r *http.Request) {
	userId := "test-user"
	notesResult, _ := n.noteService.ReadNotes(r.Context(), userId)
	rest.NewResponse(w, notesResult)
}

func (n *noteController) WriteNote(w http.ResponseWriter, r *http.Request) {
	var note note.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		rest.NewBusinessError(w, errorcode.RequestParsingError, http.StatusUnprocessableEntity)
		return
	}

	if note.Content == "" {
		rest.NewBusinessError(w, errorcode.SanitizeError, http.StatusUnprocessableEntity)
		return
	}

	_ = n.noteService.WriteNote(r.Context(), note)
	rest.NewResponse(w, nil)
}

func New(
	noteService noteservice.NoteService,
) noteController {
	return noteController{
		noteService: noteService,
	}
}
