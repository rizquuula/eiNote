package httpresponse

import (
	"core/pkg/customerror"
	"core/pkg/errorcode"
	"core/pkg/responsestatus"
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  responsestatus.ResponseStatus
	Code    errorcode.ErrorCode
	Message string
	Data    interface{}
}

func (r Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Status  string      `json:"status"`
		Code    string      `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Status:  r.Status.String(),
		Code:    r.Code.String(),
		Message: r.Message,
		Data:    r.Data,
	})
}

func NewResponse(w http.ResponseWriter, message string, data interface{}) {
	resp := Response{
		Status:  responsestatus.OK,
		Code:    errorcode.Null,
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewResponseError(w http.ResponseWriter, err error) {
	if customErr, ok := err.(*customerror.CustomError); ok {
		w.Header().Set("Content-Type", "application/json")

		var message string
		if customErr.Message != "" {
			message = customErr.Message
		} else {
			message = customErr.Error()
		}

		resp := Response{
			Status:  responsestatus.SystemError,
			Code:    customErr.Code,
			Message: message,
		}

		if customErr.IsBusinessErr {
			resp.Status = responsestatus.BusinessError

			if resp.Code == errorcode.NotFoundError {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
