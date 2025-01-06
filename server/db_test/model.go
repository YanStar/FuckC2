package sqlmt

import (
	//"Orca_Server/define/config"
	"fmt"

	//"Orca_Server/define/config"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
	"gorm.io/driver/sqlite"
)

var Db *gorm.DB

func GetDb() *gorm.DB {
	var err error
	pwd, _ := os.Getwd()
	s := strings.Replace(pwd, "\\", "/", -1)

	db_path := s+"/db/test.db"

	fmt.Println("ssssss path " + db_path)
	Db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	//defer Db.Close()
	if err != nil {
		log.Println("数据库连接失败！", err)
	}
	if err != nil {
		panic(err)
	}
	//Db.LogMode(true)  //sql调试模式
	return Db
}

type HostList struct {
	gorm.Model
	ClientId  string
	Hostname  string
	Ip        string
	ConnPort  string
	Os        string
	Privilege string
	Version   string
	Remarks   string
}