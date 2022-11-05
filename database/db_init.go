package database

//import gorm and mysql
import (
	"fmt"
	"meditation/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBQuery *gorm.DB

func DatabaseMain() {
	dsn, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/meditation"), &gorm.Config{})

	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}
	fmt.Println(dsn)

	DBQuery = dsn
	MigrationModel()

}

func MigrationModel() {
	DBQuery.AutoMigrate(models.Music{}, models.User{})
}
