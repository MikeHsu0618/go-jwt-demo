package pg

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type Config struct {
	Host string
	User string
	Db   string
	Pwd  string
	Port string
}

func NewPgClient(master Config) *gorm.DB {
	db, err = gorm.Open(postgres.Open(GetPgDns(master)), &gorm.Config{})
	if err != nil || db.Error != nil {
		log.Fatalln("db error")
	}

	return db
}

func GetPgDns(conf Config) (dsn string) {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=Asia/Taipei",
		conf.Host, conf.User, conf.Pwd, conf.Db, conf.Port)
}
