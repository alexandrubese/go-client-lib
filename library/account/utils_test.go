package account

import (
	"encoding/json"
	Consts "library/consts"
	"testing"
)

// TestUnitCreateAcccountData calls CreateAcccountData with some AccountAttributes, checking
// for a valid return value.
func TestUnitCreateAcccountData(t *testing.T) {
	accountAttributes := AccountAttributes{
		AccountClassification: "Personal",
		AccountMatchingOptOut: new(bool),
		AccountNumber:         "41426819",
		AlternativeNames:      []string{"Alex Bese"},
		BankID:                "400300",
		BankIDCode:            "GBDSC",
		BaseCurrency:          "GBP",
		Bic:                   "NWBKGB22",
		Country:               "GB",
		Iban:                  "GB11NWBK40030041426819",
		Name:                  []string{"Alexandru Bese"},
	}
	accountDataDTO := CreateAcccountData(&accountAttributes)

	jsonAccountAttrData, marshalError := json.Marshal(accountDataDTO.Data.Attributes)
	if marshalError != nil {
		t.Fatalf("Marshal Error: " + marshalError.Error())
	}

	expectedAccountData := `{"account_classification":"Personal","account_matching_opt_out":false,"account_number":"41426819","alternative_names":["Alex Bese"],"bank_id":"400300","bank_id_code":"GBDSC","base_currency":"GBP","bic":"NWBKGB22","country":"GB","iban":"GB11NWBK40030041426819","name":["Alexandru Bese"]}`

	if expectedAccountData != string(jsonAccountAttrData) {
		t.Fatalf("Test failed ")
	}
}

// TestUnitGetFirstAccountWithAccounts calls GetFirstAccount, checking
// for a valid response.
func TestUnitGetFirstAccountWithAccounts(t *testing.T) {
	accountAttributes := AccountAttributes{
		AccountClassification: "Personal",
		AccountMatchingOptOut: new(bool),
		AccountNumber:         "41426819",
		AlternativeNames:      []string{"Alex Bese"},
		BankID:                "400300",
		BankIDCode:            "GBDSC",
		BaseCurrency:          "GBP",
		Bic:                   "NWBKGB22",
		Country:               "GB",
		Iban:                  "GB11NWBK40030041426819",
		Name:                  []string{"Alexandru Bese"},
	}

	accResponse := AccountsResponse{
		Data: []AccountData{
			{
				ID:         "553c7766-04fa-43e0-a12d-c096b2cc86c5",
				Attributes: &accountAttributes,
			},
		},
	}

	_, err := GetFirstAccount(accResponse)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

// TestUnitGetFirstAccountWithoutAccounts calls GetFirstAccount, checking
// for a valid response.
func TestUnitGetFirstAccountWithoutAccounts(t *testing.T) {
	accResponse := AccountsResponse{
		Data: []AccountData{},
	}

	_, err := GetFirstAccount(accResponse)

	if err.Error() != Consts.NO_ACCOUNTS_ERROR {
		t.Fatalf("TestUnitGetFirstAccountWithoutAccounts failed")
	}
}
