package account

import (
	Consts "library/consts"
	Helpers "library/helpers"
	"net/http"
)

func (acc *Account) Fetch(accountId string) *http.Response {
	fetchRequest, fetchRequestError := http.NewRequest(
		http.MethodGet,
		Consts.API_URL+Consts.ACCOUNTS+accountId,
		nil,
	)

	if fetchRequestError != nil {
		return Helpers.HTTPInternalServerError(Consts.FETCH_CREATION_ERROR + fetchRequestError.Error())
	}

	fetchExecResponse, fetchExecError := Client.Do(fetchRequest)
	if fetchExecError != nil {
		return Helpers.HTTPInternalServerError(Consts.FETCH_ERROR + fetchExecError.Error())
	}

	return fetchExecResponse
}
