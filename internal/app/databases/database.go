package databases

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/trile99/go_api_server/internal/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database instance
type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

// Connect function
func Connect() {
	// p := configs.Config("DB_PORT")
	// because our config function returns a string, we are parsing our      str to int here
	// port, err := strconv.ParseUint(p, 10, 32)
	// if err != nil {
	// 	fmt.Println("Error parsing str to int")
	// }
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")

	fmt.Println("DATABASE_USER:", dbUser)
	fmt.Println("DB_PASSWORD:", dbPassword)
	fmt.Println("DATABASE_NAME:", dbName)
	fmt.Println("DATABASE_HOST:", dbHost)
	fmt.Println("DATABASE_PORT:", dbPort)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	log.Print(dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database. \n")
		os.Exit(2)
	}
	log.Print("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Print("running migrations")
	db.AutoMigrate(&models.User{})
	DB = Dbinstance{
		Db: db,
	}
}
