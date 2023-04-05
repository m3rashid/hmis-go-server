package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupGorm() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)

	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	DB_SSLMODE := os.Getenv("DB_SSLMODE")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")

	dataSourceName := fmt.Sprintf("host=%s user=%s port=%s password=%s dbname=%s sslmode=%s", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT, DB_SSLMODE)

	db, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN:                  dataSourceName,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}),
		&gorm.Config{
			Logger:                                   newLogger,
			DisableForeignKeyConstraintWhenMigrating: true,
		},
	)

	// TODO: setup prometheus metrics
	// db.Use(prometheus.New(prometheus.Config{
	// 	DBName:          "db1",
	// 	RefreshInterval: 15,                          // Refresh metrics interval (default 15 seconds)
	// 	PushAddr:        "prometheus pusher address", // push metrics if `PushAddr` configured
	// 	StartServer:     true,
	// 	HTTPServerPort:  4001,
	// 	MetricsCollector: []prometheus.MetricsCollector{
	// 		// user defined metrics
	// 		&prometheus.Postgres{},
	// 	},
	// }))

	db.AutoMigrate(
		&Address{},
		&Resource{},
		&Permission{},
		&Role{},
		&Availability{},
		&User{},
		&Profile{},
		&Leave{},
		&Appointment{},
		&Consumables{},
		&NonConsumables{},
		&PrescribedMedicine{},
		&Prescription{},
	)

	return db, err
}
