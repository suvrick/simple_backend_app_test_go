package database

import (
	"fmt"
	"kiss_web/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectionDB ...
type ConnectionDB struct {
	NameDb    string
	LoginDb   string
	PassDb    string
	PortDb    uint16
	HostDb    string
	SslModeDb string
}

var (
	connection *gorm.DB
	errDB      error
	dsn        string
)

// Init ....
func init() {

	// TODO: описать иницилизацию структуры в функции , с возможностью загрузки из окружения
	cd := &ConnectionDB{
		LoginDb:   "suvrick",
		PassDb:    "bl69unn",
		HostDb:    "localhost",
		PortDb:    5432,
		NameDb:    "db1",
		SslModeDb: "disable",
	}

	dsn = fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=%v",
		cd.LoginDb,
		cd.PassDb,
		cd.NameDb,
		cd.HostDb,
		cd.PortDb,
		cd.SslModeDb)

	open()
}

func open() {
	connection, errDB = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if errDB != nil {
		log.Panicln(errDB)
		return
	}

	errDB = connection.AutoMigrate(&models.User{})
	if errDB != nil {
		log.Panicln(errDB)
	}

	log.Println("Connection db OK")
}

// GetDB return pointer connection db
func GetDB() *gorm.DB {

	if errDB != nil {
		return nil
	}

	return connection
}
