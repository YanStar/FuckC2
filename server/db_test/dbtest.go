package sqlmt

import (

	"fmt"
)


func InitHost() {
	Db = GetDb()
	Db.Exec("DELETE FROM host_lists")
	err := Db.AutoMigrate(&HostList{})
	if err != nil {
		fmt.Printf("Error migrating database: %v\n", err)
		return
	}
}

func main()  {

	fmt.Println("db test")
}
