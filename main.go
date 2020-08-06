package main

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // 包装
)

func main() {
	db, err := gorm.Open("mysql", "root:@(localhost)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	user := User{}
	user.ID = 1
	user.Name = sql.NullString{}
}

//User - a register user
type User struct {
	ID   int
	Name sql.NullString
}
