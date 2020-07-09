package server

import (
	"github.com/TicketsBot/translationgenerator/config"
	"github.com/TicketsBot/translationgenerator/server/endpoints"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config config.Config
	router *gin.Engine
}

func NewServer(config config.Config) *Server {
	return &Server{
		config: config,
		router: gin.Default(),
	}
}

func (s *Server) ConfigureEngine() {
	s.router.LoadHTMLFiles("./public/templates/index.tmpl")

	s.router.GET("/", endpoints.IndexHandler)
	s.router.POST("/json", endpoints.JsonHandler)
}

func (s *Server) Start() {
	if err := s.router.Run(s.config.Server.Host); err != nil {
		panic(err)
	}
}
