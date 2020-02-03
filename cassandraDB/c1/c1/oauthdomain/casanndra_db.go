package oauthdomain

import (
	"errors"

	access_token "github.com/k-washi/golang-cookbook/cassandraDB/c1/c1/accessToken"

	"github.com/gocql/gocql"

	"github.com/k-washi/golang-cookbook/cassandraDB/c1/c1/cassandra"
)

const (
	queryGetAcessToken     = "SELECT access_token, user_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, expires) VALUES (?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

func NewRepo() DbRepo {
	return &dbRepo{}
}

type DbRepo interface {
	GetID(string) (*access_token.AccessToken, error)
	Create(access_token.AccessToken) error
	UpdateExpirationTime(access_token.AccessToken) error
}

type dbRepo struct{}

func (r *dbRepo) GetID(at string) (*access_token.AccessToken, error) {
	session := cassandra.GetSession()

	var result access_token.AccessToken
	if err := session.Query(queryGetAcessToken, at).Scan(
		&result.AccessToken,
		&result.UserID,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.New("no access token found with given id")
		}
		return nil, err
	}

	return &result, nil

}

func (r *dbRepo) Create(at access_token.AccessToken) error {
	session := cassandra.GetSession()

	if err := session.Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserID,
		at.Expires,
	).Exec(); err != nil {
		return err
	}
	return nil
}

func (r *dbRepo) UpdateExpirationTime(at access_token.AccessToken) error {
	session := cassandra.GetSession()

	if err := session.Query(queryUpdateExpires,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		return err
	}
	return nil
}
