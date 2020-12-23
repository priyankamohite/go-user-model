package databases

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID        uint   `gorm:"not null"`
	Fname     string `gorm:"not null"`
	Lname     string
	Email     *string
	DOB       *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateUsersTable() {
	//ConnectDB()

	sqlDB.LogMode(true)

	sqlDB, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")

	if err != nil {
		log.Panic(err)
	}
	fmt.Println("connected to db successfuly")

	// Create table for `User`
	sqlDB.AutoMigrate().DropTable(&User{})

	// Create table for `User`
	sqlDB.AutoMigrate().CreateTable(&User{})

	// Append "ENGINE=InnoDB" to the creating table SQL for `User`
	sqlDB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate().CreateTable(&User{})

}
