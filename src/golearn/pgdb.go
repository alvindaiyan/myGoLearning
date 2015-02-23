package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // _操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数
)

func main() {
	// start connection
	db, err := sql.Open("postgres", "user=postgres password=Password1 dbname=godb sslmode=disable")
	checkErr(err)

	// insert
	fmt.Println("insert data")
	stmt, err := db.Prepare("INSERT INTO userinfo(username, password, created) VALUES($1, $2, $3) RETURNING uid")
	checkErr(err)

	res, err := stmt.Exec("test", "Fronde2", "2014-1-6")
	checkErr(err)
	// id, err := res.LastInsertId()
	id, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("row affected ", id)

	// stmt, err := db.Prepare("select uid, username, password from userinfo where uid=$1")
	// checkErr(err)

	// rows, err := stmt.Query(1)

	// for rows.Next() {
	// 	var uid int
	// 	var username string
	// 	var password string
	// 	err = rows.Scan(&uid, &username, &password)
	// 	checkErr(err)
	// 	fmt.Println(uid)
	// 	fmt.Println(username)
	// 	fmt.Println(password)
	// }

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
