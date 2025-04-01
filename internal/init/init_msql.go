package init

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL() *gorm.DB {
	info := "root:2003@tcp(localhost:3306)/ecommerce?parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(info), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to MySQL database: ", err)
	}

	log.Println("Connected to MySQL database successfully!")
	return db
}
