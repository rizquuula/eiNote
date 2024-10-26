package notebook

import (
	"net/http"
	notebookservice "stub/internal/domain/service/notebook"
	"stub/internal/interfaces/rest"
)

type notebookController struct {
	notebookService notebookservice.NotebookService
}

func (n *notebookController) ReadNotebooks(w http.ResponseWriter, r *http.Request) {
	userId := "test-user"
	notesResult, _ := n.notebookService.ReadNotebooks(r.Context(), userId)
	rest.NewResponse(w, notesResult)
}

func New(
	notebookService notebookservice.NotebookService,
) notebookController {
	return notebookController{
		notebookService: notebookService,
	}
}
