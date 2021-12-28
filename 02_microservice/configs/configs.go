package configs

import (
	"fmt"
	"os"
)

const (
	HASHKEY = "c5dda7a77f7dc8e29cd2d949ccc201c02e1afdd5d4a44993d2a81509d53c6954"

	MYSQL_USER     = "user"
	MYSQL_PASSWORD = "2021Mysql!!"
	MYSQL_ADDRESS  = "172.17.0.1"
	MYSQL_PORT     = "3306"

	DB_NAME = "klikacc"
)

func GetMySqlDSN() (result string) {

	user := os.Getenv("MYSQL_USER")
	if user == "" {
		user = MYSQL_USER
	}
	password := os.Getenv("MYSQL_PASS")
	if password == "" {
		password = MYSQL_PASSWORD
	}
	address := os.Getenv("MYSQL_ADDRESS")
	if address == "" {
		address = MYSQL_ADDRESS
	}
	port := os.Getenv("MYSQL_PORT")
	if port == "" {
		port = MYSQL_PORT
	}
	user = MYSQL_USER
	password = MYSQL_PASSWORD
	address = MYSQL_ADDRESS
	port = MYSQL_PORT

	result = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, address, port, DB_NAME)
	return
}
