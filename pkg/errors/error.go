package errors

import "github.com/jinzhu/gorm"

func Error(e error) {
	panic(e)
}

func DBErrorCheck(db *gorm.DB) {
	if err := db.Error; err != nil {
		Error(err)
	}
}

func ErrorCheck(e error) {
	if e != nil {
		Error(e)
	}
}
