package main

import (
	"github.com/drborges/go-web/api"
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	m.Use(martini.Logger())
	m.Use(martini.Static("/public"))

	api.Register(m)

	m.Run()
}
