package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//Todo Model
//
type Todo struct {
	gorm.Model
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

//Todos Collection
//
type Todos []Todo

//InitDB Database initializer
//
// func InitDB() *gorm.DB {
//
// 	db, err := gorm.Open("sqlite3", "gorm.db")
//
// 	if err != nil {
// 		log.Fatalf("Got error when connect database, the error is '%v'", err)
// 	}
//
// 	// defer db.Close()
//
// 	// Migrate the schema
// 	db.AutoMigrate(&Todo{})
//
// 	Create
// 	seed := Todos{
// 		Todo{Name: "Going out"},
// 		Todo{Name: "Go to the Cinema"},
// 	}
//
// 	for _, todo := range seed {
// 		db.Create(&Todo{Name: todo.Name})
// 	}
//
// }
