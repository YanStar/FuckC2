package common

import (
	pb "SimpleC2RpcTest/protobuf"
	//pb "SimpleC2RpcTest/protobuf"
	//sqlmt "SimpleC2RpcTest/server/db_test"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

func PrintTable(host_list []*pb.HostInfo, identity int) {
	var data [][]string
	table := tablewriter.NewWriter(os.Stdout)
	table.SetColWidth(48)

	table.SetHeader([]string{"id", "Hostname", "Ip", "Os", "Arch", "Privilege", "Port"})
	for i, host := range host_list {
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