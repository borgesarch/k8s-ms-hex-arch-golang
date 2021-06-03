package data

import (
	u "clean-code-golang/core/domain/users"

	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	c "gorm.io/gorm"
)

type IConnection interface {
	Context() *c.DB
}

type Connection int

func (c *Connection) Context() *c.DB {

	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	DBNAME := os.Getenv("DB_NAME")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	SSLMOD := os.Getenv("DB_SSLMOD")
	TIMEZONE := os.Getenv("DB_TIMEZONE")

	var dns string = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		HOST, USER, PASS, DBNAME, PORT, SSLMOD, TIMEZONE)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	db.AutoMigrate(&u.User{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
