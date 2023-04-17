package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// A Form has a url.Values and errors
type Form struct {
	url.Values
	Errors errors
}

// There there are no errors on the form return true
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// Creates as new Form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// This variadic functions accepts any number of strings.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be empty")
		}
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

// Check that the field is at least of the minimum length
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters", length))
		return false
	}
	return true
}
