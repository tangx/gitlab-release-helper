package confgitlab

import (
	"errors"
	"log"
	"os"

	"github.com/xanzy/go-gitlab"
)

type Server struct {
	gitlab *gitlab.Client
}

func (s *Server) Init() {

	git, err := s.gitlabClient()
	if err != nil {
		log.Fatalf("gitlab initial client failed: %v", err)
	}

	s.gitlab = git
}

func (s *Server) gitlabClient() (*gitlab.Client, error) {
	// https://docs.gitlab.com/ee/ci/jobs/ci_job_token.html
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
