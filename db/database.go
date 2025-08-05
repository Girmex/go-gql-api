package db
import (
	"fmt"
	"gorm.io/driver/postgres"
	"github.com/Girmex/go-gql-api/graph/model"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=1234 dbname=go-gql port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	// Migrate models
	err = db.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Comment{},
	)
	if err != nil {
		log.Fatalf("failed to migrate DB: %v", err)
	}

	DB = db
	fmt.Println("Database connected and migrated")
	return DB
}
