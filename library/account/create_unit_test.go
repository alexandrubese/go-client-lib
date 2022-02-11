package account

import (
	Consts "library/consts"
	Helpers "library/helpers"
	"library/mocks"
	"testing"
)

// TestUnitSuccessCreate calls Create checking
// for a valid response.
func TestUnitSuccessCreate(t *testing.T) {
	var acc Account
	Client = &mocks.MockClient{}

	accountAttributes := AccountAttributes{
		AccountClassification: "Personal",
		AccountMatchingOptOut: new(bool),
		AccountNumber:         "41426819",
		AlternativeNames:      []string{"Alex Bese1"},
		BankID:                "400300",
		BankIDCode:            "GBDSC",
		BaseCurrency:          "GBP",
		Bic:                   "NWBKGB22",
		Country:               "GB",
		Iban:                  "GB11NWBK40030041426819",
		Name:                  []string{"Alexandru Bese1"},
	}

	mocks.MockHTTPResponse(Consts.MOCK_ACCOUNT, 200)
	resp := acc.Create(&accountAttributes)
	body := Helpers.ExtractBody(resp)

	if body != Consts.MOCK_ACCOUNT {
		t.Fatalf("TestCreate failed - response is not equal to the mock")
	}
}

// TestUnitFailCreate calls Create checking
// for a valid response.
func TestUnitFailCreate(t *testing.T) {
	var acc Account
	Client = &mocks.MockClient{}
	errorMessage := `An error`

	accountAttributes := AccountAttributes{
		AccountClassification: "Personal",
		AccountMatchingOptOut: new(bool),
		AccountNumber:         "41426819",
		AlternativeNames:      []string{"Alex Bese1"},
		BankID:                "400300",
		BankIDCode:            "GBDSC",
		BaseCurrency:          "GBP",
		Bic:                   "NWBKGB22",
		Country:               "GB",
		Iban:                  "GB11NWBK40030041426819",
		Name:                  []string{"Alexandru Bese1"},
	}

	mocks.MockHTTPError(errorMessage)
	resp := acc.Create(&accountAttributes)
	body := Helpers.ExtractBody(resp)

	if body != Consts.CREATE_ERROR+errorMessage {
		t.Fatalf("TestCreate failed - response is not equal to the mock")
	}
}
