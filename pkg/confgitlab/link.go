package confgitlab

type AssertLink struct {
	Name string
	URL  string
}

// func (s *Server) CreateReleaseLink(tag string, link AssertLink) error {

// 	// s.gitlab.ReleaseLinks.CreateReleaseLink()
// 	pid := s.env("CI_PROJECT_ID")
// 	_url := filepath.Join(s.HostPrefix, link.URL)
// 	linkOpt := &gitlab.CreateReleaseLinkOptions{
// 		Name: &link.Name,
// 		URL:  &_url,
// 	}

// 	_, _, err := s.gitlab.ReleaseLinks.CreateReleaseLink(pid, tag, linkOpt)
// 	return err
// }
