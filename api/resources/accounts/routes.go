package accounts

import (
	"github.com/go-martini/martini"
	"github.com/drborges/go-web/api/resources/accounts/handlers"
)

func Register(router martini.Router) {
	router.Get("/accounts", handlers.FetchAllAccounts)
	router.Post("/accounts", handlers.CreateAccount)
}