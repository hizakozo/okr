package infrastructure

import (
	_ "github.com/go-sql-driver/mysql" //コード内で直接参照するわけではないが、依存関係のあるパッケージには最初にアンダースコア_をつける
	"github.com/jinzhu/gorm"           //ここでパッケージをimport
	"log"
)

var Db *gorm.DB

var url = "root" + ":" + "root" + "@tcp(" + "127.0.0.1" + ")/okr"

func init() {
	var err error
	Db, err = gorm.Open(
		"mysql",
		url)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB connected !!")
	Db.SingularTable(true)
	Db.LogMode(true)
	return
}
