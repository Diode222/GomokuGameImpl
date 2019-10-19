package impl

import (
	"context"
	"errors"
	"github.com/Diode222/GomokuGameImpl/conf"
	pb "github.com/Diode222/GomokuGameImpl/proto"
	"github.com/sirupsen/logrus"
	"sync"
)

type gomokuGameImplServer struct{}

var server *gomokuGameImplServer
var once sync.Once

func NewGGomokuGameImplServer() *gomokuGameImplServer {
	once.Do(func() {
		server = &gomokuGameImplServer{}
	})
	return server
}

func (s *gomokuGameImplServer) Init(ctx context.Context, isFirst *pb.IsFirst) (*pb.Status, error) {
	// TODO
	// You need to implement this function. isFirst means if you are first-hand.
	// You can return any status to me, when your init has completed.
	conf.IS_FIRST_HAND = isFirst.GetIsFirst()
	if conf.IS_FIRST_HAND {
		conf.MY_PIECE_TYPE = pb.PieceType_BLANK
	} else {
		conf.MY_PIECE_TYPE = pb.PieceType_WHITE
	}
	// TODO init

	status := true
	return &pb.Status{
		Status: &status,
	}, nil
}

func (s *gomokuGameImplServer) MakePiece(ctx context.Context, board *pb.Board) (*pb.PiecePosition, error) {
	// TODO
	// You need to implement this function. `board` is server's response board state right now,
	// you can get all pieces' info of all positions in board, and the board's info like length and height,
	// You should make one piece as `*pb.PiecePosition` back to me.

	// TODO change to your implemention
	chessPositions := board.GetChessPositions()
	for _, chessPosition := range chessPositions {
		if chessPosition.GetType() == pb.PieceType_NONE {
			pieceType := conf.MY_PIECE_TYPE
			piecePositionX := chessPosition.GetPosition().GetX()
			piecePositionY := chessPosition.GetPosition().GetY()
			logrus.WithFields(logrus.Fields{
				"Type": pieceType,
				"X":    piecePositionX,
				"Y":    piecePositionY,
			}).Info("Piece make.")
			return &pb.PiecePosition{
				Type: &pieceType,
				Position: &pb.Position{
					X: &piecePositionX,
					Y: &piecePositionY,
				},
			}, nil
		}
	}

	logrus.Warn("what's wrong of referee???")
	return nil, errors.New("what's wrong of referee???")
}
