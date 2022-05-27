package api

import (
	"net/http"
)

type Response struct {
	Success  bool
	Errors   []string
	Status   int
	Response interface{}
}

func (res *Response) Init() *Response {
	res.Status = 400
	res.Response = nil
	res.Errors = []string{}
	return res
}

func AddJSONHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
