package middlewares

import (
	"github.com/go-martini/martini"
)

func Register(m *martini.ClassicMartini) {
	m.Use(Auth)
}