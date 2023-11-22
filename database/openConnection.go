package database

import (
	entitiesauth "portofolio-api/auth/entities"
	entitieshortlink "portofolio-api/shortlink/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	dialect := postgres.Open("host=localhost user=root password=secret dbname=portofolio_api port=5432 sslmode=disable TimeZone=Asia/Jakarta")
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic("failed Connecting database")
	}
	return db
}

func AutoMigration() {
	db := OpenConnection()
	db.AutoMigrate((&entitieshortlink.Shortlink_tab{}))
	db.AutoMigrate((&entitiesauth.User_tab{}))
}
