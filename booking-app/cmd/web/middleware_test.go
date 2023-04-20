package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var m myHandler
	h := NoSurf(&m)

	// We can use a type switch to check that NoSurf returns a http.Handler
	switch v := h.(type) {
	case http.Handler:
		// If it is an http.Handler do nothing
	default:
		// If it is not of type http.Handler fail the test
		t.Errorf("type is not http.Handler, but is %T", v)
	}
}

func TestSessionLoad(t *testing.T) {
	var m myHandler
	h := SessionLoad(&m)

	switch v := h.(type) {
	case http.Handler:
		// If it is an http.Handler do nothing
	default:
		t.Errorf("type is not http.Handler, but is %T", v)
	}
}
