package databases

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func ConnectDB() {
	sqlDB, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("connected to db successfuly")
	fmt.Printf("%v", sqlDB)
}
