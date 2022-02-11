package consts

import (
	Helpers "library/helpers"
)

var (
	API_HOSTNAME = Helpers.GetHostname()
	API_URL      = "http://" + API_HOSTNAME + ":8080/v1/"
)

const ACCOUNT_TYPE = "accounts"
const ACCOUNTS = "organisation/accounts/"
const ACCOUNT_VERSION = "0"

const TEST_ACCOUNT_ID = "553c7766-04fa-43e0-a12d-c096b2cc86c2"
const MOCK_ACCOUNT = `{"data":{"attributes":{"account_classification":"Personal","account_matching_opt_out":false,"account_number":"41426819","alternative_names":["Alex Bese1"],"bank_id":"400300","bank_id_code":"GBDSC","base_currency":"GBP","bic":"NWBKGB22","country":"GB","iban":"GB11NWBK40030041426819","name":["Alexandru Bese1"]},"created_on":"2022-01-25T08:08:25.917Z","id":"553c7766-04fa-43e0-a12d-c096b2cc86c5","modified_on":"2022-01-25T08:08:25.917Z","organisation_id":"fd1c7e3d-4371-4356-895e-5ab3a44e24d2","type":"accounts","version":0},"links":{"self":"/v1/organisation/accounts/68763849-5ede-4b4c-8256-ced2e0419e68"}}`

const FETCH_CREATION_ERROR = "An error occured with the account retrieval request: "
const FETCH_ERROR = "An error occured with the account retrieval request: "
const CREATE_CREATION_ERROR = "An error occured with the account create request creation: "
const CREATE_ERROR = "An error occured with the account create request: "
const DELETE_CREATION_ERROR = "An error occured with the account deletion request creation: "
const DELETE_ERROR = "An error occured with the account deletion request: "
const MARSHAL_ERROR = "An error occured while trying to marshal: "
const FETCH_ACCOUNTS_CREATION_ERROR = "GetAccounts failed - Failed to create retrieve request"
const FETCH_ACCOUNTS_ERROR = "GetAccounts failed - Failed to retrieve accounts"
const NO_ACCOUNTS_ERROR = "There are no accounts"
