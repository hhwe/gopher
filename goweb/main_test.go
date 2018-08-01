//main_test.go

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// Here, we form a new HTTP request. This is the request that's going to be
	// passed to our handler.
	// The first argument is the method, the second argument is the route (which
	// we leave blank for now, and will get back to soon), and the third is the
	// request body, which we don't have in this case.
	req, err := http.NewRequest("GET", "", nil)

	// In case there is an error in forming the request, we fail and stop the test
	if err != nil {
		t.Fatal(err)
	}

	// We use Go's httptest library to create an http recorder. This recorder
	// will act as the target of our http request
	// (you can think of it as a mini-browser, which will accept the result of
	// the http request that we make)
	recorder := httptest.NewRecorder()

	// Create an HTTP handler from our handler function. "handler" is the handler
	// function defined in our main.go file that we want to test
	hf := http.HandlerFunc(handler)

	// Serve the HTTP request to our recorder. This is the line that actually
	// executes our the handler that we want to test
	hf.ServeHTTP(recorder, req)

	// Check the status code is what we expect.
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `Hello World!`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actual, expected)
	}
}
