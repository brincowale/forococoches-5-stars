package utils

import (
	"forococoches-5-stars/models"
	"github.com/getsentry/sentry-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var err error

func CreateConnectionDB(database string) {
	DB, err = gorm.Open("mysql", database)
	if err != nil {
		sentry.CaptureException(err)
	}
}

func IsNewThread(thread models.Thread) bool {
	var t models.Thread
	DB.Where(&models.Thread{Id: thread.Id}).First(&t)
	return t.Id == 0
}

func InsertThread(thread models.Thread) {
	DB.Create(&thread)
}
