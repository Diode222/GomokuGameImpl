package main

import (
	"errors"
	"fmt"
	"github.com/Diode222/GomokuGameImpl/conf"
	"github.com/Diode222/GomokuGameImpl/impl"
	"github.com/Diode222/GomokuGameImpl/proto"
	"github.com/Diode222/logS"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"strconv"
)

func main() {
	listenIP := conf.LISTEN_IP
	// Log addr should be setted by backend, player1 and player2's log files are different
	logFileAddr := os.Getenv("LOG_VOLUME_ADDR")
	initLogHook(logFileAddr)

	// Do not know firsthand or backhand, so do not know port, should be setted by backend
	// Port is fixed: firstHand---10001, backHand---10002
	servicePort, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"port": os.Getenv("PORT"),
		}).Info("Wrong port env.")
		return
	}
	//servicePort := 10001

	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()

	proto.RegisterMakePieceServiceServer(grpcServer, impl.NewGGomokuGameImplServer())
	listenAddr := fmt.Sprintf("%s:%d", listenIP, servicePort)
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Fatal("Failed to start listen.")
		return
	} else {
		logrus.WithFields(logrus.Fields{
			"listenAddr": listenAddr,
		}).Info("listening failed in pod.")
	}
	defer listener.Close()

	err = grpcServer.Serve(listener)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"LISTEN_ADDR": listenAddr,
		}).Fatal("Start GomokuGameImpl service failed.")
		return
	}
}

func initLogHook(logFileAddr string) {
	hook := logS.NewHook(logFileAddr)
	if hook == nil {
		logrus.WithFields(logrus.Fields{
			"logFileAddr": logFileAddr,
		}).Fatal("File log hook created failed.")
		panic(errors.New("File log hook created failed."))
	}
	logrus.AddHook(hook)
}
