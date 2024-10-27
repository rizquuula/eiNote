package note

import (
	noteservice "core/internal/domain/service/note"
	"core/pkg/customerror"
	"core/pkg/errorcode"
	"core/pkg/httpresponse"
	"encoding/json"
	"net/http"

	"core/internal/domain/model/note"
)

type noteController struct {
	noteService noteservice.NoteService
}

func (n *noteController) ReadNotes(w http.ResponseWriter, r *http.Request) {
	userId := "test-user"
	queryParams := r.URL.Query()
	notebookId := queryParams.Get("notebook")
	notesResult, err := n.noteService.ReadNotes(r.Context(), notebookId, userId)
	if err != nil {
		httpresponse.NewResponseError(w, err)
		return
	}
	httpresponse.NewResponse(w, "Success Read Notes", notesResult)
}

func (n *noteController) WriteNote(w http.ResponseWriter, r *http.Request) {
	var note note.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		err = customerror.NewSystemError(err, customerror.Opts{
			Code:    errorcode.RequestParsingError,
			Message: err.Error(),
		})
		httpresponse.NewResponseError(w, err)
		return
	}

	if note.ID.String() == "" || note.Content == "" {
		httpresponse.NewResponseError(w, err)
		return
	}

	_ = n.noteService.WriteNote(r.Context(), note)
	httpresponse.NewResponse(w, "Success Write Note", nil)
}

func New(
	noteService noteservice.NoteService,
) noteController {
	return noteController{
		noteService: noteService,
	}
}
