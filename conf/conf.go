package conf

import "github.com/Diode222/GomokuGameImpl/proto"

const (
	ETCD_ADDR = "127.0.0.1:2379"

	LISTEN_IP = "127.0.0.1"
)

var (
	IS_FIRST_HAND bool
	MY_PIECE_TYPE proto.PieceType
)
