package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBMaster  *gorm.DB // master
	DBReplica *gorm.DB // replica
)

func ConnectDatabase() {
	// Kết nối master
	dsnMaster := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_MASTER_HOST"),
		os.Getenv("DB_MASTER_USER"),
		os.Getenv("DB_MASTER_PASSWORD"),
		os.Getenv("DB_MASTER_NAME"),
		os.Getenv("DB_MASTER_PORT"),
	)

	dbMaster, err := gorm.Open(postgres.Open(dsnMaster), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to master DB: " + err.Error())
	}

	// Kết nối replica
	dsnReplica := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_REPLICA_HOST"),
		os.Getenv("DB_REPLICA_USER"),
		os.Getenv("DB_REPLICA_PASSWORD"),
		os.Getenv("DB_REPLICA_NAME"),
		os.Getenv("DB_REPLICA_PORT"),
	)

	dbReplica, err := gorm.Open(postgres.Open(dsnReplica), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to replica DB: " + err.Error())
	}

	fmt.Println("✅ Connected to Master DB!")
	fmt.Println("✅ Connected to Replica DB!")

	DBMaster  = dbMaster
	DBReplica = dbReplica
}
