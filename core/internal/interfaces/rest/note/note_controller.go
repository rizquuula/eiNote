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
	notebookId := queryParams.Get("notebook_id")
	notesResult, err := n.noteService.ReadNotes(r.Context(), notebookId)
	if err != nil {
		httpresponse.NewResponseError(w, err)
		return
	}
	httpresponse.NewResponse(w, "Success Read Notes", notesResult)
}

func (n *noteController) WriteNote(w http.ResponseWriter, r *http.Request) {

	var requestBody struct {
		ID         string `json:"id"`
		NotebookId string `json:"notebook_id"`
		Content    string `json:"content"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		err = customerror.NewSystemError(err, customerror.Opts{
			Code:    errorcode.RequestParsingError,
			Message: err.Error(),
		})
		httpresponse.NewResponseError(w, err)
		return
	}

	if requestBody.Content == "" {
		httpresponse.NewResponseError(w, err)
		return
	}

	requestNote := note.Note{
		Content: requestBody.Content,
	}

	err = requestNote.IdFromStr(requestBody.ID)
	if err != nil {
		httpresponse.NewResponseError(w, err)
		return
	}

	err = requestNote.NotebookIdFromStr(requestBody.NotebookId)
	if err != nil {
		httpresponse.NewResponseError(w, err)
		return
	}

	newNote, err := n.noteService.WriteNote(r.Context(), requestNote)

	if err != nil {
		httpresponse.NewResponseError(w, err)
		return
	}
	httpresponse.NewResponse(w, "Success Write Note", newNote)
}

func (n *noteController) DeleteNote(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	noteId := queryParams.Get("note_id")
	err := n.noteService.DeleteNote(r.Context(), noteId)
	if err != nil {
		httpresponse.NewResponseError(w, err)
		return
	}
	httpresponse.NewResponse(w, "Success Delete Note", nil)
}

func New(
	noteService noteservice.NoteService,
) noteController {
	return noteController{
		noteService: noteService,
	}
}
