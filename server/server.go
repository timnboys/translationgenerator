package server

import (
	"github.com/TicketsBot/translationgenerator/server/endpoints"
	"github.com/gin-gonic/gin"
	"os"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	return &Server{
		router: gin.Default(),
	}
}

func (s *Server) ConfigureEngine() {
	s.router.LoadHTMLFiles("./public/templates/index.tmpl")

	s.router.GET("/", endpoints.IndexHandler)
	s.router.POST("/json", endpoints.JsonHandler)
}

func (s *Server) Start() {
	if err := s.router.Run(os.Getenv("SERVER_ADDR")); err != nil {
		panic(err)
	}
}
