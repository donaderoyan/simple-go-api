package config

import (
	"os"

	model "github.com/donaderoyan/simple-go-api/models"
	util "github.com/donaderoyan/simple-go-api/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/postgres"
)

func Connection() *gorm.DB {
	databaseURI := make(chan string, 1)

	if os.Getenv("GO_ENV") != "production" {
		databaseURI <- util.GodotEnv("DATABASE_URI_DEV")
	} else {
		databaseURI <- os.Getenv("DATABASE_URI_PROD")
	}

	db, err := gorm.Open(postgres.Open(<-databaseURI), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error(), databaseURI)
	}

	if os.Getenv("GO_ENV") != "production" {
		logrus.Info("Connection to Database Successfully")
	}

	for _, model := range model.RegisterModels() {
		err := db.Debug().AutoMigrate(model.Model)

		if err != nil {
			logrus.Fatal(err.Error())
		}
	}

	return db
}
