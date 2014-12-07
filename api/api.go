package api

import (
	"github.com/drborges/go-web/api/resources/accounts"
	"github.com/drborges/go-web/api/middlewares"
	"github.com/go-martini/martini"
)

func Register(m *martini.ClassicMartini) {
	accounts.Register(m)

	middlewares.Register(m.Martini)
}
