package models

// Request send to name server/CNAME request body
type Request struct {
	Domain string `json:"domain"`
}
