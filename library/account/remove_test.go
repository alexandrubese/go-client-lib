package account

import (
	Consts "library/consts"
	Helpers "library/helpers"
	"library/mocks"
	"testing"
)

// TestUnitSuccessRemove calls Acc.Delete, checking
// for a valid response.
func TestUnitSuccessRemove(t *testing.T) {
	var acc Account
	Client = &mocks.MockClient{}

	mocks.MockHTTPResponse(Consts.MOCK_ACCOUNT, 200)
	resp := acc.Delete(Consts.TEST_ACCOUNT_ID, Consts.ACCOUNT_VERSION)
	body := Helpers.ExtractBody(resp)

	if body != Consts.MOCK_ACCOUNT {
		t.Fatalf("TestFetch failed - response is not equal to the mock")
	}
}

// TestUnitFailRemove calls Acc.Delete, checking
// for a valid response.
func TestUnitFailRemove(t *testing.T) {
	var acc Account
	Client = &mocks.MockClient{}
	errorMessage := `An error`

	mocks.MockHTTPError(errorMessage)
	resp := acc.Delete(Consts.TEST_ACCOUNT_ID, Consts.ACCOUNT_VERSION)
	body := Helpers.ExtractBody(resp)

	if body != Consts.DELETE_ERROR+errorMessage {
		t.Fatalf("TestFailFetch failed - response is not equal to the mock")
	}
}
