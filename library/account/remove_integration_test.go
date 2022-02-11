package account

import (
	"encoding/json"
	"io/ioutil"
	Consts "library/consts"
	"net/http"
	"strconv"
	"testing"
)

// TestIntegrationDelete calls Delete checking
// for a valid response.
func TestIntegrationRemove(t *testing.T) {
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
		t.Fatalf("IntegrationDelete failed: " + firstAccountError.Error())
	}

	accountId := firstAccount.ID
	version := strconv.FormatInt(firstAccount.Version, 10)

	resp := acc.Delete(accountId, version)

	if resp.StatusCode != 204 {
		t.Fatalf("IntegrationDelete failed - Failed to delete account: " + accountId)
	}
}
