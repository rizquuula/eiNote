package rest

import (
	"encoding/json"
	"net/http"
	"stub/internal/domain/enum/errorcode"
	"stub/internal/domain/enum/responsestatus"
)

type Response struct {
	Status string      `json:"status"`
	Code   string      `json:"code"`
	Data   interface{} `json:"data"`
}

func (r Response) New(
	status responsestatus.ResponseStatus,
	code errorcode.ErrorCode,
	data interface{},
) Response {
	return Response{
		Status: status.String(),
		Code:   code.String(),
		Data:   data,
	}
}

func NewResponse(w http.ResponseWriter, data interface{}) {
	resp := Response{}.New(
		responsestatus.OK,
		errorcode.Null,
		data,
	)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewBusinessError(w http.ResponseWriter, code errorcode.ErrorCode, statusCode uint) {
	resp := Response{}.New(
		responsestatus.BusinessError,
		code,
		nil,
	)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(int(statusCode))

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewSystemError(w http.ResponseWriter) {
	resp := Response{}.New(
		responsestatus.BusinessError,
		errorcode.SystemError,
		nil,
	)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(int(http.StatusInternalServerError))

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
