package postgres

import (
	"fmt"
	"github.com/alexaxms/TopSecret/config"
	"github.com/alexaxms/TopSecret/postgres/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

func NewDB() *gorm.DB {
	DBMS := "postgres"
	dbConf := config.C.Database
	pgConf := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", dbConf.User, dbConf.Password, dbConf.DBName, dbConf.SSL)

	db, err := gorm.Open(DBMS, pgConf)

	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.Board{})

	return db
}
