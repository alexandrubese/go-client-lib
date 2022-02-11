package account

import (
	Consts "library/consts"
	Helpers "library/helpers"
	"net/http"
)

func (acc *Account) FetchAccounts() *http.Response {
	fetchRequest, fetchRequestError := http.NewRequest(
		http.MethodGet,
		Consts.API_URL+Consts.ACCOUNTS,
		nil,
	)

	if fetchRequestError != nil {
		return Helpers.HTTPInternalServerError(Consts.FETCH_ACCOUNTS_CREATION_ERROR + fetchRequestError.Error())
	}

	fetchExecResponse, fetchExecError := Client.Do(fetchRequest)
	if fetchExecError != nil {
		return Helpers.HTTPInternalServerError(Consts.FETCH_ACCOUNTS_ERROR + fetchExecError.Error())
	}

	return fetchExecResponse
}
