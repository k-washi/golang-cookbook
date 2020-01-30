package db

import (
	"github.com/k-washi/golang-cookbook/oAuth/c1/domain/access_token"
)

func NewRepository() DbRepository {
	return &DbRepository{}
}

type DbRepository interface {
	GetByID(string) (*access_token.AccessToken, error)
}

type DbRepository struct{}

func (r *DbRepository) GetByID(string) (*access_token.AccessToken, error) {
	return nil, nil
}
