package models

type ResponseNs struct {
	Status     string   `json:"status"`
	Error      string   `json:"error"`
	NameServer []string `json:"name_server"`
}

type ResponseErr struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func (e *ResponseErr) PresentError(err error) *ResponseErr {
	return &ResponseErr{
		Status: "false",
		Error:  err.Error(),
	}
}

// PresentNs is a presentation method to show response of name server response
func (ns *ResponseNs) PresentNs(resNs ResponseNs) *ResponseNs {
	return &ResponseNs{
		Status:     "true",
		Error:      "nil",
		NameServer: resNs.NameServer,
	}
}
