package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func FetchSingleUser() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/tt?charset=utf8")
	//db, err := sql.Open("mysql", "root:Xixihaha_0504@/Estone?charset=utf8")
	// mysql：root 密码：Xixihaha_0504
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)

	for rows.Next() {
		var id int
		var user_id string
		var mobile string
		var name string
		var card_id string

		err = rows.Scan(&id, &user_id, &mobile, &name, &card_id,)
		checkErr(err)
		fmt.Println(user_id)
		fmt.Println(name)
		fmt.Println(mobile)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}