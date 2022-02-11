package helpers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

var (
	body = "Test body"
)

// TestHTTPInternalServerError calls HTTPInternalServerError with a body, checking
// for a valid return value.
func TestHTTPInternalServerError(t *testing.T) {
	msg := HTTPInternalServerError(body)

	responseBody, _ := ioutil.ReadAll(msg.Body)
	if msg.StatusCode != 500 || string(responseBody) != body {
		t.Fatalf("Test failed ! Expected a different StatusCode or Body")
	}
}

// TestExtractBody calls ExtractBody with a request, checking
// for a valid return value.
func TestExtractBody(t *testing.T) {
	response := &http.Response{
		Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
		StatusCode: 500,
	}

	msg := ExtractBody(response)
	if msg != body {
		t.Fatalf("Test failed ! " + "Got: " + msg + "Expected: " + body)
	}
}

func TestGetHostname(t *testing.T) {
	msg := GetHostname()

	if msg != "localhost" && msg != "interviewapp" {
		t.Fatalf("Test failed ! " + "Got: " + msg + "Expected: localhost || interviewapp")
	}
}
