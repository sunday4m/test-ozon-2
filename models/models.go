package models

type Document struct {
	Header    string   `json:"header"`
	LineItems []string `json:"line_items"`
}
