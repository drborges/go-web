package routes

import (
	"github.com/drborges/go-web/api/handlers"
	"github.com/go-martini/martini"
)

func RegisterAccountsApi(m *martini.ClassicMartini) {
	m.Get("/hello", handlers.HelloWorld)
}
