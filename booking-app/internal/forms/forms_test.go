package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/testurl", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/testurl", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/testurl", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

// Just like the above we should write two cases for each Validator
func TestForm_Has(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	if form.Has("a") {
		t.Error("form shows valid when has field is missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	if !form.Has("a") {
		t.Error("form show value is missing when it is present")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)
	if form.MinLength("a", 10) {
		t.Error("the field passes min length validation despite being empty")
	}

	isError := form.Errors.Get("a")
	if isError == "" {
		t.Error("there should be an error for the given value but it is not present")
	}

	postedValues = url.Values{}
	postedValues.Add("b", "ab")
	form = New(postedValues)
	if form.MinLength("b", 5) {
		t.Error("the field passes min length validation despite being less than 5 characters")
	}

	postedValues = url.Values{}
	postedValues.Add("c", "abcdefghijklmnop")
	form = New(postedValues)
	if !form.MinLength("c", 5) {
		t.Error("the field fails validation despite having more than 5 characters")
	}

	isError = form.Errors.Get("c")
	if isError != "" {
		t.Error("there should not be an error for the given value but one is present")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)
	form.IsEmail("a")
	if form.Valid() {
		t.Error("form shows valid email for field that is not present")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "abc123")
	form = New(postedValues)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("form shows email is valid when it is not")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "abc@email.com")
	form = New(postedValues)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("form shows email is not valid when it is a valid email")
	}
}
