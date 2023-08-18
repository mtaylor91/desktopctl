package service

func (s *Service) Start() error {
	return s.server.ListenAndServe()
}
