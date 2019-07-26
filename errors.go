package gson

import "encoding/json"

// APIerror ...
type APIerror struct {
	Status int             `json:"-"`
	Title  string          `json:"title"`
	Detail json.RawMessage `json:"detail,omitempty"`
	Tags   []string        `json:"tags,omitempty"`
	source error
}

// NewAPIerror ...
func NewAPIerror(title string, status int, src error, tags ...string) *APIerror {
	err := &APIerror{
		Status: status,
		Title:  title,
		Tags:   tags,
	}
	if src != nil {
		err.source = src
		if _, ok := src.(ValidationError); ok {
			err.Detail = json.RawMessage(src.Error())
		} else {
			err.Detail, _ = json.Marshal(src.Error())
		}
	}
	return err
}

func (err *APIerror) Error() string {
	b, _ := json.Marshal(err)
	return string(b)
}

// ValidationError ...
type ValidationError map[string]string

func (err ValidationError) Error() string {
	b, _ := json.Marshal(err)
	return string(b)
}

// Add ...
func (err ValidationError) Add(key, msg string) {
	err[key] = msg
}
