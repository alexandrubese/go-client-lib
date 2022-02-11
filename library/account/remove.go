package account

import (
	Consts "library/consts"
	Helpers "library/helpers"
	"net/http"
)

func (acc *Account) Delete(accountId string, version string) *http.Response {
	deleteRequest, deleteRequestError := http.NewRequest(
		http.MethodDelete,
		Consts.API_URL+Consts.ACCOUNTS+accountId+"?version="+version,
		nil,
	)

	if deleteRequestError != nil {
		return Helpers.HTTPInternalServerError(Consts.DELETE_CREATION_ERROR + deleteRequestError.Error())
	}

	deleteExecResponse, deleteExecError := Client.Do(deleteRequest)
	if deleteExecError != nil {
		return Helpers.HTTPInternalServerError(Consts.DELETE_ERROR + deleteExecError.Error())
	}

	return deleteExecResponse
}
