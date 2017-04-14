package api

// Response is the response that is put together from the API
// It is to be extended by other structs based on the data that they expext to get back
type Response interface {
	SetStatus(int)
}

type responseBase struct {
	Success    bool
	StatusCode int
	Errors     []string `json:"errors"`
}

func (r *responseBase) SetStatus(status int) {
	r.StatusCode = status
	r.Success = status >= 200 && status < 300
}
