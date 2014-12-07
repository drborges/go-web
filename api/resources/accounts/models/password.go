package models

import "errors"

type PasswordConstraint func(string) (bool, error)

type Password struct {
	Value string
	Constraints []PasswordConstraint
}

// Constraints a password to be at least 8 characters long
func MinCharactersConstraint(password string) (bool, error) {
	if len(password) > 8 {
		return true, ""
	}

	return false, errors.New("Password must be at least 8 characters long")
}

func (p Password) Validate() (bool, []error) {
	return true, _ // TODO implement proper logic
}

//
func NewWeakPassword(value string) (Password, []error) {
	password := Password{value, []PasswordConstraint{MinCharactersConstraint}}

	if valid, errors := password.Validate(); !valid {
		return _, errors
	}

	return password, _
}