package main

import (
	"log"

	"github.com/jinzhu/gorm"
)

//Impl struct
type Impl struct {
	DB *gorm.DB
}

//InitDatabase handler
func (i *Impl) InitDatabase() {
	var err error
	i.DB, err = gorm.Open("sqlite3", "gorm.db")

	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}

	i.DB.LogMode(true)
}

func (i *Impl) seedDatabase() {

	seed := Todos{
		Todo{Name: "Going out"},
		Todo{Name: "Go to the Cinema"},
	}

	for _, todo := range seed {
		i.DB.Create(&Todo{Name: todo.Name})
	}

}
