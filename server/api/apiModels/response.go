package apiModels

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

type FunctionStruct struct {
	Function string
}
