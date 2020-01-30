package app

import (
	"github.com/k-washi/golang-cookbook/oAuth/c1/domain/access_token"
	"github.com/k-washi/golang-cookbook/oAuth/c1/domain/access_token/db"
)

func StartApplication() {
	atService := access_token.NewService(db.NewRepository())
}
