package account

import (
	"errors"
	Consts "library/consts"

	"github.com/google/uuid"
)

func CreateAcccountData(accountAttributes *AccountAttributes) AccountDTO {
	accountData := AccountData{
		ID:             uuid.NewString(),
		OrganisationID: uuid.NewString(),
		Type:           Consts.ACCOUNT_TYPE,
		Attributes:     accountAttributes,
	}

	return AccountDTO{
		Data: &accountData,
	}
}

func GetFirstAccount(accResponse AccountsResponse) (AccountData, error) {
	if len(accResponse.Data) > 0 {
		account := accResponse.Data[0]
		return account, nil
	} else {
		return AccountData{}, errors.New(Consts.NO_ACCOUNTS_ERROR)
	}
}
