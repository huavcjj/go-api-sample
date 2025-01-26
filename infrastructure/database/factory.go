package database

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	InstanceMySQL int = iota
	InstancePostgres
	InstanceSQLite
)

var (
	errNoDatabaseInstance = errors.New("no database instance")
)

func NewDatabase(instance int) (db *gorm.DB, err error) {
	switch instance {
	case InstanceMySQL:
		config := NewConfigMySQL()
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
			config.User, config.Password, config.Host, config.Port, config.Database)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case InstancePostgres:
		config := NewConfigPostgres()
		dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			config.Host, config.Port, config.User, config.Database, config.Password)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case InstanceSQLite:
		config := NewConfigSQLite()
		db, err = gorm.Open(sqlite.Open(config.Database), &gorm.Config{})
	default:
		return nil, errNoDatabaseInstance
	}
	return db, err
}

func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func Migrate(db *gorm.DB, models ...interface{}) error {
	return db.AutoMigrate(models...)
}
