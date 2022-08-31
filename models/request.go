package models

// Request send to NS/CNAME/txt  request body
type Request struct {
	Domain string `json:"domain"`
}
