package models

import (
	"log"
)

func TableCreate() {
	db, err := DBConn()
	if err != nil {
		log.Print(err)
	}
	var account Account
	if !db.HasTable(&account) {
		db.CreateTable(&account)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Print(err)
		}
	}()

	row,err := db.Rows()
	if err !=nil{
		log.Print(err)
	}
	
}


