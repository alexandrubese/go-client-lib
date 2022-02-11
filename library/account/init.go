package account

import (
	C "library/httpclient"
	"net/http"
)

var (
	// Need to define an interface with a Do function that we can change for Mocking
	Client C.HTTPClient
)

func init() {
	Client = &http.Client{}
}
