package model

import (
	"database/sql"
	"fmt"
	"log"

	// use mysq
	_ "github.com/go-sql-driver/mysql"
	"github.com/spkills/spkills/config"
)

var db *sql.DB

func InitDB(conf config.Config) {
	fmt.Println(conf.Database.User)
	fmt.Println(conf.Database.Password)
	fmt.Println(conf.Database.Server)
	fmt.Println(conf.Database.Port)
	fmt.Println(conf.Database.DbName)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Server,
		conf.Database.Port,
		conf.Database.DbName,
	)

	fmt.Println(dsn)

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}
