package models

// Password - JSON struct for request
type Password struct {
	New string `json:"new"`
	Now string `json:"now"`
}
