package notebook

import (
	"core/internal/domain/model/notebook"
	notebookservice "core/internal/domain/service/notebook"
	"core/pkg/customerror"
	"core/pkg/errorcode"
	"core/pkg/httpresponse"
	"encoding/json"
	"net/http"
)

type notebookController struct {
	notebookService notebookservice.NotebookService
}

func (n *notebookController) ReadNotebooks(w http.ResponseWriter, r *http.Request) {
	notesResult, err := n.notebookService.ReadNotebooks(r.Context())
	if err != nil {
		httpresponse.NewResponseError(w, err)
		return
	}
	httpresponse.NewResponse(w, "Success Read Notebooks", notesResult)
}

func (n *notebookController) WriteNotebook(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		NotebookId string `json:"id"`
		Name       string `json:"name"`
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

	if requestBody.Name == "" {
		httpresponse.NewResponseError(w, err)
		return
	}

	requestNotebook := notebook.Notebook{}
	requestNotebook.SetName(requestBody.Name)
	err = requestNotebook.NotebookIdFromStr(requestBody.NotebookId)
	if err != nil {
		httpresponse.NewResponseError(w, err)
		return
	}

	newNote, err := n.notebookService.WriteNotebook(r.Context(), requestNotebook)

	if err != nil {
		httpresponse.NewResponseError(w, err)
		return
	}
	httpresponse.NewResponse(w, "Success Write Notebook", newNote)
}

func (n *notebookController) DeleteNotebook(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	notebookId := queryParams.Get("notebook_id")

	err := n.notebookService.DeleteNotebook(r.Context(), notebookId)
	if err != nil {
		httpresponse.NewResponseError(w, err)
		return
	}
	httpresponse.NewResponse(w, "Success Delete Notebook", nil)
}

func New(
	notebookService notebookservice.NotebookService,
) notebookController {
	return notebookController{
		notebookService: notebookService,
	}
}
