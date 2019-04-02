package main

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/sqlite"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "L1213", Price: 2500})

	var product Product
	db.First(&product, 1)
	db.First(&product, "code = ?", "L1213")

	db.Model(&product).Update("Price", 2640)

	db.Delete(&product)
}
