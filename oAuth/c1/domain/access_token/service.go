package access_token

type Repository interface {
	GetByID(string) (*AccessToken, error)
}

type Service interface {
	GetByID(string) (*AccessToken, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetByID(string) (*AccessToken, error) {
	return nil, nil
}
