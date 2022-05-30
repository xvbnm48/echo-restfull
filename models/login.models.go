package models

import (
	"database/sql"
	"fmt"

	"github.com/xvbnm48/echo-restfull/db"
	"github.com/xvbnm48/echo-restfull/helpers"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM user WHERE username = ?"

	err := con.QueryRow(sqlStatement, username).Scan(&obj.Id, &obj.Username, &pwd)

	if err == sql.ErrNoRows {
		fmt.Println("username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("querry error")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("hash and password not match")
		return false, err
	}

	return true, nil
}
