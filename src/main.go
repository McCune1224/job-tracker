package main

import (
	// "github.com/gin-gonic/gin"
	"github.com/mccune1224/job-tracker/pkg/services/repository"
	"gorm.io/gorm"
    _ "github.com/joho/godotenv/autoload"
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
