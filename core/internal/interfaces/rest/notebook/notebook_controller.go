package notebook

import (
	notebookservice "core/internal/domain/service/notebook"
	"core/pkg/httpresponse"
	"net/http"
)

type notebookController struct {
	notebookService notebookservice.NotebookService
}

func (n *notebookController) ReadNotebooks(w http.ResponseWriter, r *http.Request) {
	userId := "test-user"
	notesResult, err := n.notebookService.ReadNotebooks(r.Context(), userId)
	if err != nil {
		httpresponse.NewResponseError(w, err)
		return
	}
	httpresponse.NewResponse(w, "Success Read Notebooks", notesResult)
}

func New(
	notebookService notebookservice.NotebookService,
) notebookController {
	return notebookController{
		notebookService: notebookService,
	}
}
