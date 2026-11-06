package main

import (
	"fmt"
	"log"
	"os"

	"github.com/example/person-crud-go/internal/handler"
	"github.com/example/person-crud-go/internal/repository"
	"github.com/example/person-crud-go/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5433"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "person_db"),
	)

	runMigrations(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	repo := repository.NewPersonRepository(db)
	svc := service.NewPersonService(repo)
	h := handler.NewPersonHandler(svc)

	r := gin.Default()
	h.RegisterRoutes(r)

	log.Println("Server running on :8080")
	r.Run(":8080")
}

func runMigrations(dsn string) {
	// golang-migrate dùng URL format khác với GORM DSN
	migrateURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5433"),
		getEnv("DB_NAME", "person_db"),
	)

	m, err := migrate.New("file://db/migration", migrateURL)
	if err != nil {
		log.Fatal("migration init error:", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("migration failed:", err)
	}

	log.Println("migrations applied successfully")
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
