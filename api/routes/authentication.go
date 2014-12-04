package routes

import (
	"github.com/drborges/go-web/api/handlers"
	"github.com/go-martini/martini"
)

func RegisterAuthenticationApi(m *martini.ClassicMartini) {
	m.Get("/authenticate", handlers.Authentication)
}
