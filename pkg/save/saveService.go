package save

type SaveService interface {
	Save() error
}

type Repository interface {
	Save() error
}

type service struct {
	r Repository
}

func NewService(r Repository) SaveService {
	return &service{r}
}

func (s *service) Save() error {
	return s.r.Save()
}