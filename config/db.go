package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

func InitDb() (*sql.DB, error){
cfg := mysql.NewConfig()
    cfg.User = os.Getenv("DBUSER")
    cfg.Passwd = os.Getenv("DBPASS")
    cfg.Net = "tcp"
    cfg.Addr = "centerbeam.proxy.rlwy.net:20618"
    cfg.DBName = "railway"
    cfg.ParseTime = true
    cfg.Loc = time.UTC

    db, err := sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
		return nil, err
    }

	pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
		return nil, err
    }
    fmt.Println("Connected!")

	return db, nil

}