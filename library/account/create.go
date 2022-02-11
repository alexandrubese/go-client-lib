package account

import (
	"bytes"
	"encoding/json"
	Consts "library/consts"
	Helpers "library/helpers"
	"net/http"
)

func (acc *Account) Create(accountAttributes *AccountAttributes) *http.Response {
	accountData := CreateAcccountData(accountAttributes)

	jsonAccountData, marshalError := json.Marshal(accountData)
	if marshalError != nil {
		return Helpers.HTTPInternalServerError(Consts.MARSHAL_ERROR + marshalError.Error())
	}

	createRequest, createRequestError := http.NewRequest(
		http.MethodPost,
		Consts.API_URL+Consts.ACCOUNTS,
		bytes.NewBuffer(jsonAccountData),
	)

	if createRequestError != nil {
		return Helpers.HTTPInternalServerError(Consts.CREATE_CREATION_ERROR + createRequestError.Error())
	}

	createResponse, createResponseError := Client.Do(createRequest)
	if createResponseError != nil {
		return Helpers.HTTPInternalServerError(Consts.CREATE_ERROR + createResponseError.Error())
	}

	return createResponse
}
