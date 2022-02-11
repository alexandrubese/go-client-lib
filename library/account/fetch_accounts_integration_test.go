package account

import (
	"net/http"
	"testing"
)

// TestIntegrationFetch calls Fetch checking
// for a valid response.
func TestIntegrationFetchAcccounts(t *testing.T) {
	var acc Account
	Client = &http.Client{}

	resp := acc.FetchAccounts()

	if resp.StatusCode != 200 {
		t.Fatalf("IntegrationFetchAcccounts failed - Failed to retrieve account: ")
	}
}
