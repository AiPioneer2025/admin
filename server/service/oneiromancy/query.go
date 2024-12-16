package oneiromancy

type QueryService struct {
}

func (s *QueryService) Query(query string) string {
	return "回复:" + query
}
