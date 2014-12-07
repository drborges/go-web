package middlewares

import (
	"github.com/go-martini/martini"
)

func Register(m *martini.Martini) {
	m.Use(Auth)
}