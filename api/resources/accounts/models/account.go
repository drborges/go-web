package models

type AccountFlags struct {
	IsActive  bool
	IsBlocked bool
	IsAdmin   bool
}

type Account struct {
	Name string
	Credentials
	AccountFlags
}

func NewAccount(name string, email string, password string) (Account, error) {
	validPassword, err := NewWeakPassword(password)

	if err != nil {
		return _, err
	}

	return Account{
		name,
		Credentials{email, validPassword},
		AccountFlags{IsActive: true, IsBlocked: false, IsAdmin: false},
	}, nil
}