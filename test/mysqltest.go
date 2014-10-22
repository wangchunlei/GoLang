package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/mydb")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Working fine!")

	stmtIns, err := db.Prepare("select * from user")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()
	var (
		id   int
		name string
	)
	for i := 1; i < 8; i++ {
		err = stmtIns.QueryRow().Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(id)
		fmt.Println(name)
	}
}
