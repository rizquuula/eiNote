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
	queryParams := r.URL.Query()
	notebookId := queryParams.Get("notebook")
	notesResult, err := n.noteService.ReadNotes(r.Context(), notebookId)
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

	if note.Content == "" {
		httpresponse.NewResponseError(w, err)
		return
	}

	note, err = n.noteService.WriteNote(r.Context(), note)
	if err != nil {
		httpresponse.NewResponseError(w, err)
		return
	}
	httpresponse.NewResponse(w, "Success Write Note", note)
}

func New(
	noteService noteservice.NoteService,
) noteController {
	return noteController{
		noteService: noteService,
	}
}
