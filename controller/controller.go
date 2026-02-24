package controller

import (
	"fmt"
	"net/http"
)

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type Controller struct {
	Base   string
	Routes []Route
}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) SetBase(path string) {
	c.Base = path
}

func (c *Controller) RegisterRoute(route Route, mux *http.ServeMux) {
	mux.HandleFunc(c.Base+route.Path, route.Handler)

	c.Routes = append(c.Routes, route)

	fmt.Println("Registered: " + c.Base + route.Path)
}
