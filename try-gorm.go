package main

import _ "fmt"
import "os"
import "github.com/jinzhu/gorm"
import "github.com/davecgh/go-spew/spew"
import _ "github.com/jinzhu/gorm/dialects/sqlite"
import log "github.com/sirupsen/logrus"

// Product demo model
type Product struct {
	gorm.Model
	Code  string `gorm:"type:varchar(100)"`
	Price uint
}

func init() {
	log.SetFormatter(&log.TextFormatter{DisableColors: false, FullTimestamp: true})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "L1213", Price: 2500})
	log.Debug("created one product")

	var product Product
	db.First(&product)
	log.WithFields(log.Fields{
		"type":    "Database",
		"product": spew.Sdump(&product),
	}).Info("Query one product")

	db.First(&product, "code = ?", "L1213")

	db.Model(&product).Update("Price", 2640)

	db.Delete(&product)
}
