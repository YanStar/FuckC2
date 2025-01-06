package common

import pb "SimpleC2RpcTest/protobuf"

// HostList 结构体，定义服务返回的数据结构
type HostInfo struct {
	client_id string
	hostname string
	ip string
	conn_port string
	privilege string
	version string
	remarks string
}


// ToProtobuf - Get the protobuf version of the object
func (h *HostInfo) ToProtobuf() *pb.HostInfo {
	return &pb.HostInfo{
		ClientId:          h.client_id,
		Hostname:        h.hostname,
		Ip: h.ip,
		ConnPort:    h.conn_port,
		Privilege:        h.privilege,
		Version:     h.version,
		Remarks: h.remarks,
	}
}