package confgitlab

import (
	"log"
	"os"
	"strings"

	"github.com/xanzy/go-gitlab"
)

type Server struct {
	ReleasePrefix string `env:""`
	envMap        map[string]string
	gitlab        *gitlab.Client
}

func (s *Server) SetDefaults() {
	s.loadEnv()
}

func (s *Server) Init() {
	s.SetDefaults()

	git, err := gitlab.NewClient(
		s.env("CI_REGISTRY_PASSWORD"),
		gitlab.WithBaseURL(s.env("CI_API_V4_URL")),
	)
	if err != nil {
		log.Fatalf("gitlab initial client failed: %v", err)
	}

	s.gitlab = git
}

func (s *Server) env(key string) string {
	return s.envMap[key]
}

func (s *Server) loadEnv() {
	if s.envMap == nil {
		s.envMap = make(map[string]string)
	}
	for _, env := range os.Environ() {
		kv := strings.Split(env, "=")
		k, v := kv[0], kv[1:]
		s.envMap[k] = strings.Join(v, "=")
	}
}
