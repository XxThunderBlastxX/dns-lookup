package models

// ResponseNs is response struct of name server
type ResponseNs struct {
	Status     string   `json:"status"`
	Error      string   `json:"error"`
	NameServer []string `json:"name_server"`
}

// ResponseCname is a response struct of cname
type ResponseCname struct {
	Status string `json:"status"`
	Error  string `json:"error"`
	Cname  string `json:"cname"`
}

// ResponseErr is a response thrown when caught an error
type ResponseErr struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

// PresentError is a method to present error
func (e *ResponseErr) PresentError() *ResponseErr {
	return &ResponseErr{
		Status: "false",
		Error:  e.Error,
	}
}

// PresentNs is a presentation method to show response of name server response
func (ns *ResponseNs) PresentNs() *ResponseNs {
	return &ResponseNs{
		Status:     "true",
		Error:      "nil",
		NameServer: ns.NameServer,
	}
}

// PresentCname is a presentation method to show response of cname response
func (c *ResponseCname) PresentCname() *ResponseCname {
	return &ResponseCname{
		Status: "true",
		Error:  "nil",
		Cname:  c.Cname,
	}
}
