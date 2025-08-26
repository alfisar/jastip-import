package config

import (
	"fmt"
	"strings"

	"log"
	"os"
	"sync"

	"github.com/alfisar/jastip-import/database"
	"github.com/alfisar/jastip-import/domain"
	"github.com/joho/godotenv"

	"gorm.io/gorm"
)

var (
	_ = godotenv.Load(".env")
)

// configuration init config
func InitSch() {
	// Initialize connection
	dbSQL := initDBSCH()
	// Initialize sync.Pool
	domain.DataPool = sync.Pool{
		New: func() interface{} {
			return &domain.ConfigSch{
				DBSql: dbSQL,
			}
		},
	}
}

// Function to initialize DB
func initDBSCH() map[string]*gorm.DB {
	fmt.Println("DB_USE : " + os.Getenv("DB_USE"))
	fmt.Println("DB_HOST : " + os.Getenv("DB_HOST"))
	switch os.Getenv("DB_USE") {
	case "MySQL":
		if os.Getenv("DB_DESTINATION") == "" {
			log.Fatalf("Destinmation DB cannit empty")
		}

		destinations := strings.Split(os.Getenv("DB_DESTINATION"), ",")
		db, err := database.NewConnSQLs(destinations)
		if err != nil {
			log.Fatalf("Failed to connect to MySQL: %v", err)
		}
		return db
	default:
		log.Fatalln("Invalid DB_USE specified in environment variables")
	}
	return nil
}
