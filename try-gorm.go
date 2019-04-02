package main

import _ "fmt"
import "github.com/jinzhu/gorm"
import "github.com/davecgh/go-spew/spew"
import _ "github.com/jinzhu/gorm/dialects/sqlite"

// Product demo model
type Product struct {
	gorm.Model
	Code  string `gorm:"type:varchar(100)"`
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
	db.First(&product)
	spew.Dump(&product)
	db.First(&product, "code = ?", "L1213")

	db.Model(&product).Update("Price", 2640)

	db.Delete(&product)
}
