package config

import (
	"github.com/bantheus/API-Empregos-Golang/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgreSQL() (*gorm.DB, error) {
	logger := GetLogger("postgresql")

	dbURL := "postgres://admin:admin@localhost:5432/db_api_empregos"

	// _, err := os.Stat("./postgres-data:/var/lib/postgresql/data")

	// if os.IsNotExist(err) {
	// 	logger.Info("database file not found, creating...")

	// 	err = os.MkdirAll("./db", os.ModePerm)

	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	file, err := os.Create("./db:/var/lib/postgresql/data")

	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	file.Close()
	// }

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		logger.Errorf("postgresql opening error %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("postgresql automigration error %v", err)
	}

	return db, nil
}
