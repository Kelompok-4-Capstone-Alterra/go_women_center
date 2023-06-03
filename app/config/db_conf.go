package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBconf struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func (dbc *DBconf) InitDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbc.DB_Username,
		dbc.DB_Password,
		dbc.DB_Host,
		dbc.DB_Port,
		dbc.DB_Name)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		TranslateError: true,
	})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	db.AutoMigrate(
		&entity.User{},
		&entity.Admin{},
	)

	return db
}
