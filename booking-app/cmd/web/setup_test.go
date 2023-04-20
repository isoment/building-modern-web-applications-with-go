package main

import (
	"net/http"
	"os"
	"testing"
)

/*
Everything in this setup_test file will run when the tests for this package run.
os.Exit(m.Run()) will exit but before it does run the tests.
*/
func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

/*
We can create this struct that serves as a stub for the http.Handler(), it must have a
ServeHTTP method to satisfy the interface so we can create it and attach it below.
*/
type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
