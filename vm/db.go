package vm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

var Db *xorm.Engine

func CreateEngine() {
	engine, err := xorm.NewEngine("mysql", printUrl())
	if err != nil {
		panic(err)
	}

	engine.SetMaxIdleConns(10000)
	engine.SetMaxOpenConns(10000)
	engine.ShowSQL(true)

	setEngine(engine)
}

func setEngine(_engine *xorm.Engine) {
	Db = _engine
	go func() {
		for {
			err := Db.Ping()
			if err != nil {
				panic(err)
			}
			time.Sleep(1 * time.Hour)
		}
	}()

}

func printUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		"root", "123456", "127.0.0.1", "3306",
		"demo", "utf8mb4")
}
