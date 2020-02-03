package main

import (
	"fmt"

	access_token "github.com/k-washi/golang-cookbook/cassandraDB/c1/c1/accessToken"

	"github.com/k-washi/golang-cookbook/cassandraDB/c1/c1/oauthdomain"
)

func main() {
	at := access_token.GetAccessToken()
	at.AccessToken = "doNaDonaEtc"
	at.UserID = 123

	if err := at.Validate(); err != nil {
		fmt.Println(err)
	}

	dbRepo := oauthdomain.NewRepo()
	err := dbRepo.Create(at)
	if err != nil {
		fmt.Println(err)
	}

	accessToken, err := dbRepo.GetID(at.AccessToken)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(accessToken) //&{doNaDonaEtc 123 1580498041}
	/*
			cqlsh:oauth> SELECT * from access_tokens where access_token='doNaDonaEtc';

		 access_token | expires    | user_id
		--------------+------------+---------
			doNaDonaEtc | 1580498041 |     123
	*/

	fmt.Println("Process finish")
}
