package confgin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Addr    string `env:""`
	Port    int    `env:""`
	appname string

	engine *gin.Engine
}

func (s *Server) SetDefaults() {
	if s.Port == 0 {
		s.Port = 80
	}
	if s.appname == "" {
		s.appname = "app"
	}
}

func (s *Server) Init() {
	if s.engine == nil {
		s.engine = gin.Default()
	}
}

func (s *Server) Run() error {
	addr := fmt.Sprintf("%s:%d", s.Addr, s.Port)
	return s.engine.Run(addr)
}

func (s *Server) RegisterRoutes(fn func(rg *gin.RouterGroup)) {
	base := s.engine.Group(s.appname)
	fn(base)
}
