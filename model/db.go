package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// use mysq
	_ "github.com/go-sql-driver/mysql"
	"github.com/spkills/spkills/config"
	"github.com/volatiletech/sqlboiler/boil"
)

var db *sql.DB

func InitDB(conf config.Config) *sql.DB {
	var connectTo string
	_, err := os.Stat(conf.Database.SocketFile)

	if err == nil {
		connectTo = fmt.Sprintf("unix(%s)", conf.Database.SocketFile)
	} else {
		connectTo = fmt.Sprintf("tcp(%s:%s)", conf.Database.Server, conf.Database.Port)
	}

	dsn := fmt.Sprintf(
		"%s:%s@%s/%s?charset=%s&parseTime=%s&loc=%s",
		conf.Database.User,
		conf.Database.Password,
		connectTo,
		conf.Database.DbName,
		"utf8mb4",
		"true",
		"Local",
	)

	//log.Println(dsn)

	db, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	db.SetMaxIdleConns(conf.Database.MaxIdleConns)
	boil.SetDB(db)
	return db
}
