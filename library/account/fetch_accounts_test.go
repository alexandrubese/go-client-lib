package account

import (
	Consts "library/consts"
	Helpers "library/helpers"
	"library/mocks"
	"testing"
)

// TestUnitSuccessFetchAccounts calls FetchAcccounts, checking
// for a valid response.
func TestUnitSuccessFetchAccounts(t *testing.T) {
	var acc Account
	Client = &mocks.MockClient{}

	mocks.MockHTTPResponse(Consts.MOCK_ACCOUNT, 200)
	resp := acc.FetchAccounts()
	body := Helpers.ExtractBody(resp)

	if body != Consts.MOCK_ACCOUNT {
		t.Fatalf("TestFetchAccounts failed - response is not equal to the mock")
	}
}

// TestUnitFailFetchAccounts calls FetchAccounts, checking
// for a valid response.
func TestUnitFailFetchAccounts(t *testing.T) {
	var acc Account
	Client = &mocks.MockClient{}
	errorMessage := `An error`

	mocks.MockHTTPError(errorMessage)
	resp := acc.FetchAccounts()
	body := Helpers.ExtractBody(resp)

	if body != Consts.FETCH_ACCOUNTS_ERROR+errorMessage {
		t.Fatalf("TestFailFetchAccounts failed - response is not equal to the mock")
	}
}
