package main

import (
	"Go/gorm_models"
	"fmt"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	// GORM examples

	db, err := gorm.Open("mysql", "root:root@/dev_themiseum?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Print(err)
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	//db.AutoMigrate(&Product{})

	// Create
	db.Create(&model.File{ID: 1, FileHash:"0x1234", Signature:"0x1234"})

	// Read
	var file model.File
	db.First(&file, 1) // find product with id 1
	db.First(&file, "file_hash = ?", "0x1234") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&file).Update("file_hash", "0x2345")

	// Delete - delete product
	// db.Delete(&file)
}
