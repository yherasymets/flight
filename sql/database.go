package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

func main() {
	db, err := sql.Open("mysql", "yara24:password@tcp(127.0.0.1:3306)/golang")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	// insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES ('Zvir', 27)")

	// if err != nil {
	// 	panic(err)
	// }

	// defer insert.Close()

	res, err := db.Query("SELECT `name`, `age` FROM `users`")

	if err != nil {
		panic(err)
	}

	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Age)

		fmt.Println(fmt.Sprintf("User: %s with age %d", user.Name, user.Age))
	}

}
