package main

import (
	"fmt"
	"log"

	"github.com/k-washi/golang-cookbook/database/user_db/method"
)

func main() {

	//ユーザデータの挿入
	user := &method.User{ID: 1, FirstName: "Trou", LastName: "Tanaka", Email: "123s456@test.com", Status: 1}
	if err := user.Save(); err != nil {
		log.Println(err)
		//2020/01/23 03:15:08 email already exist
	}
	log.Println(user)
	//2020/01/23 04:06:32 &{1 Trou Tanaka 123s456@test.com <nil>}

	//ユーザーデータの取得(IDミス)
	user = &method.User{ID: 3}
	if err := user.Get(); err != nil {
		log.Println(err)
		//2020/01/23 03:26:24 user not found
	}
	log.Println(user)

	//ユーザデータ挿入
	user = &method.User{ID: 2, FirstName: "k", LastName: "washi", Email: "654321@test.com", Status: 1}
	if err := user.Save(); err != nil {
		log.Println(err)
	}

	//ユーザデータのStatus:1をFind
	res, err := user.FindByStatus(1)
	if err != nil {
		log.Println(err)
	}
	for i, u := range res {
		fmt.Println("user :", i, u)
		//user : 0 {37 Trou Tanaka 123s456@test.com 2020-01-23 05:07:35 +0000 UTC 1}

		//ユーザデータの取得(IDを使用するため、ループ内で処理)
		user = &method.User{ID: u.ID}
		if err := user.Get(); err != nil {
			log.Println(err)
		}
		log.Println(user)
		//2020/01/23 03:03:52 &{1 Trou Tanaka 123s456@test.com 2020-01-23 02:29:56 +0000 UTC}

		//ユーザデータ更新
		user.FirstName = "Taro2"
		user.LastName = "Tanaka2"
		user.Email = "123s456@test.com"
		user.Status = 1
		if err := user.Update(); err != nil {
			log.Println(err)
		}
		log.Println(user)
		//2020/01/23 03:54:51 &{1 Taro2 Tanaka2 123s456@test.com 2020-01-23 02:29:56 +0000 UTC}

		//ユーザデータの削除
		user = &method.User{ID: u.ID}
		if err := user.Delete(); err != nil {
			log.Println(err)
		}
		log.Println("Success Delete")

	}

}
