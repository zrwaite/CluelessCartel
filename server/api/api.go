package api

type Response struct {
	Success  bool
	Errors   []string
	Message  string
	Response interface{}
}
