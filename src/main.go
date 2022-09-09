package main

import (
	// "github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mccune1224/job-tracker/pkg/services/repository"
	"gorm.io/gorm"
	"os"
)

var (
	POSTGRESQL_URL = os.Getenv("POSTGRESQL_URL")
)

type DBServer struct {
	db *gorm.DB
}

func main() {
	DB := repository.LoadDatabase(POSTGRESQL_URL)
	repository.MigrateDatabase(DB)
}


