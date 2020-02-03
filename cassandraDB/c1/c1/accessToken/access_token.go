package access_token

import (
	"errors"
	"strings"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	Expires     int64  `json:"expires"`
}

//Web frontend - Client-Id : 123
//Andorid App - client-id : 234

func (at *AccessToken) Validate() error {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.New("invalid access token id")
	}

	if at.UserID <= 0 {
		return errors.New("invalid user id")
	}

	if at.Expires <= 0 {
		return errors.New("invalid expiration time")
	}
	return nil
}

func GetAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
