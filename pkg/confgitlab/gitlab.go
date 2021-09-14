package confgitlab

import (
	"errors"
	"log"
	"os"

	"github.com/xanzy/go-gitlab"
)

type Server struct {
	HostPrefix string `env:""`
	envMap     map[string]string
	gitlab     *gitlab.Client
}

func (s *Server) SetDefaults() {
	// s.loadEnv()
}

func (s *Server) Init() {
	s.SetDefaults()

	// https://docs.gitlab.com/ee/ci/jobs/ci_job_token.html

	git, err := s.gitlabClient()
	if err != nil {
		log.Fatalf("gitlab initial client failed: %v", err)
	}

	s.gitlab = git
}

func (s *Server) gitlabClient() (*gitlab.Client, error) {
	if s.env("CI_JOB_TOKEN") != "" {
		return gitlab.NewJobClient(
			s.env("CI_JOB_TOKEN"),
			gitlab.WithBaseURL(s.env("CI_API_V4_URL")),
		)
	}

	if s.env("PRIVATE_TOKEN") != "" {
		return gitlab.NewClient(
			s.env("PRIVATE_TOKEN"),
			gitlab.WithBaseURL(s.env("CI_API_V4_URL")),
		)
	}
	return nil, errors.New("not support")
}

func (s *Server) env(key string) string {
	return os.Getenv(key)
}

// func (s *Server) loadEnv() {
// 	if s.envMap == nil {
// 		s.envMap = make(map[string]string)
// 	}
// 	for _, env := range os.Environ() {
// 		kv := strings.Split(env, "=")
// 		k, v := kv[0], kv[1:]
// 		s.envMap[k] = strings.Join(v, "=")
// 	}
// }
