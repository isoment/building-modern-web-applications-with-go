package forms

import (
	"net/http"
	"net/url"
)

// A Form has a url.Values and errors
type Form struct {
	url.Values
	Errors errors
}

// Creates as new Form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Check if the field is in the Form and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}
