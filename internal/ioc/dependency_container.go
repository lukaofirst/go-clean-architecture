package ioc

import (
	"fmt"
	"go-clean-architecture/internal/application/services"
	"go-clean-architecture/internal/domain/entities"
	infrastructure "go-clean-architecture/internal/infrastructure/persistence"
	"go-clean-architecture/utils"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitializeDB() *gorm.DB {
	loadEnvironmentFile()

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbEncrypt := os.Getenv("DB_ENCRYPT")

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?encrypt=%s",
		dbUser, dbPassword, dbHost, dbPort, dbEncrypt)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
			NoLowerCase:   true,
		},
	})
	if err != nil {
		log.Fatalf("Failed to connect to SQL Server: %v", err)
	}

	var result int
	err = db.Raw("SELECT COUNT(*) FROM sys.databases WHERE name = ?", dbName).Scan(&result).Error
	if err != nil {
		log.Fatalf("Failed to query databases: %v", err)
	}

	if result == 0 {
		log.Printf("Database %s does not exist. Creating the database...", dbName)
		err = db.Exec("CREATE DATABASE " + dbName).Error
		if err != nil {
			log.Fatalf("Failed to create database: %v", err)
		}
		log.Printf("Database %s created successfully.", dbName)
	}

	dsn = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbEncrypt)
	db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
			NoLowerCase:   true,
		},
	})
	if err != nil {
		log.Fatalf("Failed to connect to %s database: %v", dbName, err)
	}

	err = db.AutoMigrate(&entities.Person{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	return db
}

func loadEnvironmentFile() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	log.Printf("Current working directory: %s", cwd)

	rootDir := utils.FindGoRootDirectory(cwd)
	if rootDir == "" {
		log.Fatalf("Could not find the root directory containing go.mod")
	}
	log.Printf("Go project root directory: %s", rootDir)

	if err := godotenv.Load(filepath.Join(rootDir, ".env")); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func AddPersonService(personRepository infrastructure.PersonRepository) services.PersonService {
	personService := services.NewPersonService(personRepository)

	return personService
}

func AddPersonRepository(db *gorm.DB) infrastructure.PersonRepository {
	return infrastructure.NewPersonRepository(db)
}
