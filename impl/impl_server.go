package impl

import (
	"context"
	pb "github.com/Diode222/GomokuGameImpl/proto"
	"sync"
)

type gomokuGameImplServer struct {}

var server *gomokuGameImplServer
var once sync.Once

func NewGGomokuGameImplServer() *gomokuGameImplServer {
	once.Do(func() {
		server = &gomokuGameImplServer{}
	})
	return server
}

func (s *gomokuGameImplServer) MakePiece(ctx context.Context, board *pb.Board) (*pb.PiecePosition, error) {
	// You need to implemente this function. `board` is server's response board state right now,
	// you can get all pieces' info of all positions in board, and the board's info like length and height,
	// You should make one piece as `*pb.PiecePosition` back to me.
	return nil, nil
}
