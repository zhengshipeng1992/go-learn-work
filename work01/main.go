package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"log"
)

var Db *sql.DB

type User struct {
	Id string
	Name string
}

func init() {
	var err error
	dns := "root:111111@tcp(127.0.0.1:3306)/test?charset=uft8mb4$parseTime=True"
	Db, err = sql.Open("mysql", dns)
	if err != nil {
		// sql.Open函数没有进行数据库连接，只有在驱动未注册时才会返回err!=nil
		panic(err)
	}
}

func main()  {
	defer Db.Close()
	user, err := QueryUserNameById("1")

	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("data not found, %+v\n", err)
		return
	}

	if err != nil {
		log.Printf("unknown error:%+v", err)
		return
	}

	fmt.Printf("user is name:%s", user.Name)
}


func QueryUserNameById(id string) (User, error) {
	var user User
	sql := "SELECT id, name FROM user WHERE id = ?"
	row := Db.QueryRow(sql, id)
	err := row.Scan(&user.Id, &user.Name)
	if err != nil {
		return user, errors.Wrapf(err, "sql:%s,id:%s", sql, id)
		//return user, errors.WithStack(err)
	}
	return user, nil
}