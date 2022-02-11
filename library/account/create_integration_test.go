package account

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

// IntegrationTestCreate calls Create checking
// for a valid response.
func TestIntegrationCreate(t *testing.T) {
	var acc Account
	Client = &http.Client{}

	accountAttributes := AccountAttributes{
		AccountClassification: "Personal",
		AccountMatchingOptOut: new(bool),
		AccountNumber:         "41426819",
		AlternativeNames:      []string{"Alex Bese - Integration"},
		BankID:                "400300",
		BankIDCode:            "GBDSC",
		BaseCurrency:          "GBP",
		Bic:                   "NWBKGB22",
		Country:               "GB",
		Iban:                  "GB11NWBK40030041426819",
		Name:                  []string{"Alexandru Bese - Integration"},
	}

	resp := acc.Create(&accountAttributes)

	if resp.StatusCode != 201 {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		t.Fatalf("IntegrationCreate failed: Status Code: " + strconv.FormatInt(int64(resp.StatusCode), 10) + " " + string(body))
	}
}
