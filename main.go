package main

import (
	"fmt"
	"github.com/Diode222/GomokuGameImpl/conf"
	"github.com/Diode222/GomokuGameImpl/impl"
	"github.com/Diode222/GomokuGameImpl/proto"
	"github.com/Diode222/etcd_service_discovery/etcdservice"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
)

func main() {
	serviceName := os.Getenv("SERVICE_NAME")
	listenIP := conf.LISTEN_IP
	serviceIP := conf.SERVICE_IP
	servicePort, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Println("Wrong port param, port: %s", os.Getenv("SERVICE_PORT"))
	}
	ttl, err := strconv.Atoi(os.Getenv("TTL"))
	if err != nil {
		log.Println("Wrong ttl param, port: %s", os.Getenv("TTL"))
	}
	etcdAddr := conf.ETCD_ADDR

	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()

	log.Println(fmt.Sprintf("serviceName: %s, listenIP: %s, serviceIP: %s, servicePort: %s, ttl: %s", serviceName, listenIP, serviceIP, servicePort, ttl))
	proto.RegisterMakePieceServiceServer(grpcServer, impl.NewGGomokuGameImplServer())
	err = etcdservice.NewServiceManager(etcdAddr).Register(serviceName, listenIP, serviceIP, servicePort, grpcServer, int64(ttl)) // 修改服务发现这里ttl参数类型为int64不统一
	if err != nil {
		log.Println("Start GomokuGameImpl service failed, SERVICE_NAME: %s, LISTEN_IP: %s, SERVICE_IP: %s, SERVICE_PORT: %d")
	}
}
