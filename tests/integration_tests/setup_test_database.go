package integration_tests

import (
	"context"
	"fmt"
	"go-clean-architecture/internal/domain/entities"
	"go-clean-architecture/utils"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func setupTestContainer() (*gorm.DB, func(), error) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	log.Printf("Current working directory: %s", cwd)

	// Find the Go project root directory by looking for go.mod
	rootDir := utils.FindGoRootDirectory(cwd)
	if rootDir == "" {
		log.Fatalf("Could not find the root directory containing go.mod")
	}
	log.Printf("Go project root directory: %s", rootDir)

	// Load the .env file from the root directory
	if err := godotenv.Load(filepath.Join(rootDir, ".env")); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbEncrypt := os.Getenv("DB_ENCRYPT")

	// Set up the context for container
	ctx, cancelFunc := context.WithCancel(context.Background())

	// Set up the container request
	req := testcontainers.ContainerRequest{
		Image:        "mcr.microsoft.com/mssql/server:2022-latest",
		ExposedPorts: []string{"1433/tcp"},
		Env: map[string]string{
			"ACCEPT_EULA": "Y",
			"SA_PASSWORD": dbPassword, // Use the password from the .env file
		},
		WaitingFor: wait.ForLog("SQL Server is now ready for client connections").WithStartupTimeout(120 * time.Second),
	}

	// Start the container
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		cancelFunc()
		log.Fatalf("failed to start container: %v", err)
	}

	// Get the container host and port
	host, err := container.Host(ctx)
	if err != nil {
		cancelFunc()
		log.Fatalf("failed to get container host: %v", err)
	}
	port, err := container.MappedPort(ctx, "1433")
	if err != nil {
		cancelFunc()
		log.Fatalf("failed to get container port: %v", err)
	}

	// Add a short delay for database initialization
	time.Sleep(5 * time.Second)

	// Create the connection string
	connStr := fmt.Sprintf("sqlserver://%s:%s@%s:%s?encrypt=%s",
		dbUser, dbPassword, host, port.Port(), dbEncrypt)
	log.Printf("Connecting to SQL Server at %s:%s", host, port.Port())

	// Initialize GORM
	gormDB, err := gorm.Open(sqlserver.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
			NoLowerCase:   true,
		},
	})
	if err != nil {
		cancelFunc()
		log.Fatalf("failed to initialize GORM: %v", err)
	}

	// Verify the connection
	sqlDB, err := gormDB.DB()
	if err != nil {
		cancelFunc()
		log.Fatalf("failed to get sql.DB from GORM: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		cancelFunc()
		log.Fatalf("failed to ping database: %v", err)
	}

	log.Printf("Database is ready at %s:%s\n", host, port.Port())

	// Check if the GoCleanArchitecture database exists
	var result int
	err = gormDB.Raw("SELECT COUNT(*) FROM sys.databases WHERE name = ?", dbName).Scan(&result).Error
	if err != nil {
		cancelFunc()
		log.Fatalf("failed to query databases: %v", err)
	}

	// If the database doesn't exist, create it
	if result == 0 {
		log.Printf("Database %s does not exist. Creating the database...", dbName)
		err = gormDB.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName)).Error
		if err != nil {
			cancelFunc()
			log.Fatalf("failed to create database: %v", err)
		}
		log.Printf("Database %s created successfully.", dbName)
	}

	// Now that the database exists, close the current connection and reconnect to GoCleanArchitecture
	connStr = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=%s",
		dbUser, dbPassword, host, port.Port(), dbName, dbEncrypt)

	gormDB, err = gorm.Open(sqlserver.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
			NoLowerCase:   true,
		},
	})
	if err != nil {
		cancelFunc()
		log.Fatalf("failed to connect to %s database: %v", dbName, err)
	}

	// Perform auto-migration
	err = gormDB.AutoMigrate(&entities.Person{})
	if err != nil {
		cancelFunc()
		log.Fatalf("failed to auto-migrate: %v", err)
	}

	// Return teardown function
	teardown := func() {
		sqlDB.Close()
		container.Terminate(ctx)
		cancelFunc()
	}

	return gormDB, teardown, nil
}
