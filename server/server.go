package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	c "github.com/srutherhub/web-app/controller"
	m "github.com/srutherhub/web-app/middleware"
)

type ServerConfig struct {
	Port string
}

type Server struct {
	Controllers []c.Controller
}

func New() *Server {
	return &Server{}
}

func (s *Server) Start(config ServerConfig) {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mux := http.NewServeMux()
	registerControllers(mux, s.Controllers)

	serveStaticFiles(mux)

	err = http.ListenAndServe(":"+config.Port, mux)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s *Server) RegisterController(c c.Controller) {
	s.Controllers = append(s.Controllers, c)
}

func serveStaticFiles(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/public/", http.StripPrefix("/public/", m.SetStaticCacheHeader(fs)))
}

func registerControllers(mux *http.ServeMux, controllers []c.Controller) {
	for _, controller := range controllers {
		for _, route := range controller.Routes {
			mux.HandleFunc(controller.Base+route.Path, route.Handler)
		}
	}
}
