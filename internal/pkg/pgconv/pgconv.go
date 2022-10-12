package pgconv

import (
	"exusiai.dev/livehouse/internal/constant"
	"exusiai.dev/livehouse/internal/model/pb"
)

func ServerID(s string) uint8 {
	if i, ok := constant.ServerIDMapping[s]; ok {
		return i
	}
	return constant.ServerIDMapping[constant.DefaultServer]
}

// ServerIDFPBE converts a server ID from protobuf enum to uint8
func ServerIDFPBE(p pb.Server) uint8 {
	return ServerID(p.String())
}

func ServerIDTPBE(i uint8) pb.Server {
	return pb.Server(i)
}
