package api

import (
	"github.com/drborges/go-web/api/middlewares"
	"github.com/drborges/go-web/api/routes"
	"github.com/go-martini/martini"
)

func Register(m *martini.ClassicMartini) {
	routes.RegisterAccountsApi(m)
	routes.RegisterAuthenticationApi(m)

	middlewares.Register(m)
}
