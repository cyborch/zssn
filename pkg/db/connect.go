package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"cyborch.com/apocalypse/pkg/user"
)

type PostgresSql struct {
	Client *gorm.DB
}

// NewDatabase creates a new database connection
// and returns a pointer to the database
// or panics if the connection fails
func NewDatabase(connection gorm.Dialector) *PostgresSql {
	database, err := gorm.Open(connection, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(&user.User{}, &user.UserItem{})

	return &PostgresSql{
		Client: database,
	}
}

// Connect connects to a postgres database
// and returns a pointer to the database
func Connect(address string) *PostgresSql {
	dictator := postgres.Open(address)
	return NewDatabase(dictator)
}
