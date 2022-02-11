package account

import (
	Consts "library/consts"
	Helpers "library/helpers"
	"library/mocks"
	"testing"
)

// TestUnitSuccessFetch calls Fetch, checking
// for a valid response.
func TestUnitSuccessFetch(t *testing.T) {
	var acc Account
	Client = &mocks.MockClient{}

	mocks.MockHTTPResponse(Consts.MOCK_ACCOUNT, 200)
	resp := acc.Fetch(Consts.TEST_ACCOUNT_ID)
	body := Helpers.ExtractBody(resp)

	if body != Consts.MOCK_ACCOUNT {
		t.Fatalf("TestFetch failed - response is not equal to the mock")
	}
}

// TestUnitFailFetch calls Fetch, checking
// for a valid response.
func TestUnitFailFetch(t *testing.T) {
	var acc Account
	Client = &mocks.MockClient{}
	errorMessage := `An error`

	mocks.MockHTTPError(errorMessage)
	resp := acc.Fetch(Consts.TEST_ACCOUNT_ID)
	body := Helpers.ExtractBody(resp)

	if body != Consts.FETCH_ERROR+errorMessage {
		t.Fatalf("TestFailFetch failed - response is not equal to the mock")
	}
}
