package main

import (
	"fmt"
	"github.com/Diode222/GomokuGameImpl/conf"
	"github.com/Diode222/GomokuGameImpl/impl"
	"github.com/Diode222/GomokuGameImpl/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listenIP := conf.LISTEN_IP
	// Do not know firsthand or backhand, so do not know port, should be dispatch by referee
	//servicePort, err := strconv.Atoi(os.Getenv("PORT"))
	//if err != nil {
	//	log.Println(fmt.Sprintf("Wrong port param, port: %s", os.Getenv("PORT")))
	//	return
	//}
	servicePort := 10001

	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()

	proto.RegisterMakePieceServiceServer(grpcServer, impl.NewGGomokuGameImplServer())
	listenAddr := fmt.Sprintf("%s:%d", listenIP, servicePort)
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Failed to start listen, err: %s", err.Error()))
		return
	} else {
		log.Println(fmt.Sprintf("Listening at %s in pod", listenAddr))
	}
	defer listener.Close()

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Println(fmt.Sprintf("Start GomokuGameImpl service failed, LISTEN_ADDR: %s", listenAddr))
		return
	}
}
