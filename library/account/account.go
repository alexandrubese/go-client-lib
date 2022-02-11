package account

type Methods interface {
	Create()
	Fetch(id string)
	Delete(id string)
}

type Account struct {
}
