package helpers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
)

func HTTPInternalServerError(body string) *http.Response {
	result := &http.Response{
		Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
		StatusCode: 500,
	}
	return result
}

func ExtractBody(resp *http.Response) string {
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func GetHostname() string {
	API_HOSTNAME := os.Getenv("API_HOSTNAME")
	if len(API_HOSTNAME) == 0 {
		return "localhost"
	}

	return API_HOSTNAME
}
