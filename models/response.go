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

// ResponseMx is a response struct of mx record
type ResponseMx struct {
	Status string   `json:"status"`
	Error  string   `json:"error"`
	Mx     []string `json:"mx"`
}

// ResponseTxt is a response struct of txt record
type ResponseTxt struct {
	Status string   `json:"status"`
	Error  string   `json:"error"`
	Txt    []string `json:"txt_record"`
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

// PresentTxt is a presentation method to show response of txt response
func (t *ResponseTxt) PresentTxt() *ResponseTxt {
	return &ResponseTxt{
		Status: "true",
		Error:  "nil",
		Txt:    t.Txt,
	}
}

// PresentMx is a presentation method to show response of mx records response
func (m *ResponseMx) PresentMx() *ResponseMx {
	return &ResponseMx{
		Status: "true",
		Error:  "nil",
		Mx:     m.Mx,
	}
}
