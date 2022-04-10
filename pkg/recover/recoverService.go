package recover

type RecoverService interface {
	Recover() error
}

type Repository interface {
	Recover() error
}

type service struct {
	r Repository
}

func NewService(r Repository) RecoverService {
	return &service{r}
}

func (s *service) Recover() error {
	return s.r.Recover()
}
