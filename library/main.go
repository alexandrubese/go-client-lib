package main

import (
	Acc "library/account"
)

type Library struct {
	Account Acc.Account
	// Add new resources here
}

func main() {
	// Usage example
	/* var lib Library

	/*accountAttributes := Acc.AccountAttributes{
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

	resp := lib.Account.Create(&accountAttributes)
	*/

	/*
		resp := lib.Account.Fetch("553c7766-04fa-43e0-a12d-c096b2cc86c5")
		//resp := lib.Account.Delete("239e0182-3e60-4d1e-93a9-b67a5e5ce7dd", Consts.ACCOUNT_VERSION)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Println(string(body))
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Status)
	*/
}
