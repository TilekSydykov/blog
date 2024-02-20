package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

var engines []*xorm.Engine

func main() {
	for i := 0; i < 10; i++ {
		engine, _ := xorm.NewEngine("sqlite3", fmt.Sprintf("./test%d.db", i))
		engines = append(engines, engine)
	}
}
