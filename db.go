package lib

import (
	"database/sql"
	"fmt"
	"log"

	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MustGetEnv(envName string) string {

	envVal := os.Getenv(envName)
	if envVal == "" {
		log.Fatalf("An environment variable %v is not set!", envName)
	}

	return envVal
}
func ConnectToDB() (*gorm.DB, error) {

	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("INSTANCE_UNIX_SOCKET"))
	dbPool, err := sql.Open("pgx", dbURI)
	db, err := gorm.Open(postgres.New(postgres.Config{Conn: dbPool}), &gorm.Config{})

	if err != nil {

		log.Println("SOCKET CONNECTION NOT AVAILABLE")
		dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable",
			os.Getenv("HOST_ADDRESS"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		errorCheck(err)

		if err == nil {
			log.Printf("DB CONNECTED SUCCESSFULLY")
		}
	}

	return db, err
}
