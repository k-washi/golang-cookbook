package user_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	//Client user_id database
	Client *sql.DB
	err    error
)

/*
CREATE TABLE `user_db`.`users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `first_name` VARCHAR(45) NULL,
  `last_name` VARCHAR(45) NULL,
  `email` VARCHAR(45) NOT NULL,
  `data_created` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
	UNIQUE INDEX `email_UNIQUE` (`email` ASC));
*/

/*
//実際は、環境変数を読み込んで、DBの設定を行う。
//export mysql_user_name="root"
const (
	mysql_user_naem: "mysql_user_name"
)

var userName := os.Getenv(mysql_user_naem)
*/

func init() {
	//user:password@tcp(host:port)/dbname
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		"root", "", "127.0.0.1:3306", "user_db",
	)

	log.Println(fmt.Sprintf("about to connect to %s", dataSourceName))
	//open db by mysql driver and data source name.
	//https://github.com/go-sql-driver/mysql/issues/150 (Issue)
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
