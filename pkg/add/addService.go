package add

type AddService interface {
	Set(key string, value string)
}

type Repository interface {
	Set(key string, value string)
}

type service struct {
	r Repository
}

func NewService(r Repository) AddService {
	return &service{r}
}

func (s *service) Set(key string, value string) {
	s.r.Set(key, value)
}
