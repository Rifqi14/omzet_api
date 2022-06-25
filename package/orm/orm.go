package orm

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connection struct {
	Host                    string
	DbName                  string
	User                    string
	Password                string
	Port                    string
	Location                *time.Location
	DBMaxConnection         int
	DBMAxIdleConnection     int
	DBMaxLifeTimeConnection int
}

func (c Connection) Conn() (*gorm.DB, error) {
	// dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.DbName)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.DbName)

	DB, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err.Error())
	}

	db, _ := DB.DB()

	db.Ping()
	db.SetMaxOpenConns(c.DBMaxConnection)
	db.SetMaxIdleConns(c.DBMAxIdleConnection)
	db.SetConnMaxLifetime(time.Duration(c.DBMaxLifeTimeConnection) * time.Second)

	return DB, err
}
