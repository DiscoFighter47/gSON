package gson

import (
	"fmt"
)

// APIerror ...
type APIerror struct {
	Status int    `json:"-"`
	Title  string `json:"title"`
	Detail string `json:"detail,omitempty"`
	source error
}

// NewAPIerror ...
func NewAPIerror(title string, status int, src error) *APIerror {
	err := &APIerror{
		Status: status,
		Title:  title,
	}
	if src != nil {
		err.source = src
		err.Detail = src.Error()
	}
	return err
}

func (err *APIerror) Error() string {
	return fmt.Sprintf("%s: %s", err.Title, err.Detail)
}
