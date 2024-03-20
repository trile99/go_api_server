package databases

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rs/zerolog/log"
	"github.com/trile99/go_api_server/configs"
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
	p := configs.Config("DB_PORT")
	// because our config function returns a string, we are parsing our      str to int here
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println("Error parsing str to int")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", configs.Config("DB_HOST"), configs.Config("DB_USER"), configs.Config("DB_PASSWORD"), configs.Config("DB_NAME"), port)
	fmt.Printf(dsn)

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
