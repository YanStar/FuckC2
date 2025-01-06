package main

import (
	"fmt"
	sqlmt "SimpleC2RpcTest/server/db_test"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	//"strings"

	//"gorm.io/gorm"
)


func PrintTable(hostlists []sqlmt.HostList, identity int) {
	var data [][]string
	table := tablewriter.NewWriter(os.Stdout)
	table.SetColWidth(48)

	table.SetHeader([]string{"id", "Hostname", "Ip", "Os", "Arch", "Privilege", "Port"})
	for i, host := range hostlists {
		//verSplit := strings.Split(host.Version, ":")
		//data = append(data, []string{strconv.Itoa(i + 1), host.Hostname, host.Ip, host.Os, verSplit[2], host.Privilege, host.ConnPort})
		data = append(data, []string{strconv.Itoa(i + 1), host.Hostname, host.Ip, host.Os, host.Version, host.Privilege, host.ConnPort})
	}

	for _, raw := range data {
		if identity != -1 {
			if raw[0] == strconv.Itoa(identity) {
				table.Append(raw)
				break
			}
		} else {
			table.Append(raw)
		}
	}
	table.Render()
	return
}





func main()  {
	fmt.Println("test")
	//sqlmt.InitHost()


	sqlmt.Db = sqlmt.GetDb()
	//host_list := sqlmt.HostList{
	//	Model:     gorm.Model{},
	//	ClientId:  "implant_1",
	//	Hostname:  "shuyang",
	//	Ip:        "127.0.0.1",
	//	ConnPort:  "8888",
	//	Privilege: "admin",
	//	Os:        "windows",
	//	Version:   "0.1",
	//	Remarks:   "",
	//}
	//
	//sqlmt.Db.Create(&host_list)

	var hosts []sqlmt.HostList
	sqlmt.Db.Find(&hosts)

	PrintTable(hosts,-1)

	for _,host_info := range hosts{
		fmt.Println(host_info.Hostname)
	}

	fmt.Println(hosts)


}
