package account

import (
	"encoding/json"
	"io/ioutil"
	Consts "library/consts"
	"net/http"
	"testing"
)

// TestIntegrationFetch calls Fetch checking
// for a valid response.
func TestIntegrationFetch(t *testing.T) {
	var acc Account
	var accResponse AccountsResponse
	Client = &http.Client{}

	fetchAccountsResponse := acc.FetchAccounts()

	defer fetchAccountsResponse.Body.Close()
	body, _ := ioutil.ReadAll(fetchAccountsResponse.Body)

	if fetchAccountsResponse.StatusCode != 200 {
		t.Fatalf("IntegrationFetch failed: " + string(body))
	}

	marshalError := json.Unmarshal(body, &accResponse)
	if marshalError != nil {
		t.Fatalf(Consts.MARSHAL_ERROR + marshalError.Error())
	}

	firstAccount, firstAccountError := GetFirstAccount(accResponse)
	if firstAccountError != nil {
		t.Fatalf("IntegrationFetch failed: " + firstAccountError.Error())
	}

	accountId := firstAccount.ID

	resp := acc.Fetch(accountId)

	if resp.StatusCode != 200 {
		t.Fatalf("IntegrationFetch failed - Failed to retrieve account: " + accountId)
	}
}
