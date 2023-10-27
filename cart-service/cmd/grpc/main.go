package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/farhandwian/microservice/internal/app"
	"github.com/farhandwian/microservice/internal/repository"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/farhandwian/microservice/helper/eureka"
	"github.com/farhandwian/microservice/internal/service"

	"github.com/farhandwian/microservice/pb"
	"github.com/rs/zerolog/log" // Update
)

var ipaddress string = "localhost"

func main() {
	// register all service
	repository.RDS()
	dao := repository.NewDAO()
	cartService := service.NewCartService(dao)
	// serviceDiscoveryIp := "172.20.0.4"

	go func() {
		listener, err := net.Listen("tcp", ipaddress+":9091") // Changed the port to 9091
		if err != nil {
			log.Fatal().Err(err).Msg("error listener")
		}

		grpcServer := grpc.NewServer()
		pb.RegisterCartServiceServer(grpcServer, app.NewMicroserviceServer(cartService))
		reflection.Register(grpcServer)
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Fatal().Err(err).Msg("error serve")
		}
	}()

	go startHTTPServer()

	// serviceRegistryURL := "http://192.168.18.200:8167/eureka/apps/"

	serviceRegistryURL := "http://discovery-server:8167/eureka/apps/"
	erm := eureka.EurekaRegistrationManager{}
	erm.RegisterWithSerivceRegistry(serviceRegistryURL)
	erm.SendHeartBeat(serviceRegistryURL)

	// Block...
	wg := sync.WaitGroup{} // Use a WaitGroup to block main() exit
	wg.Add(1)
	wg.Wait()
}

func startHTTPServer() {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()} // Add insecure option for simplicity

	err := pb.RegisterCartServiceHandlerFromEndpoint(context.Background(), mux, "localhost:9091", opts)
	if err != nil {
		fmt.Println("Error registering the gateway:", err)
		return
	}
	fmt.Println("tes")
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	err = http.ListenAndServe(":9090", mux) // This remains the same
	if err != nil {
		fmt.Println("Error starting HTTP server:", err)
	}
}

// func eurekaRegister() {
// 	serviceDiscoveryIp := "192.168.18.200"
// 	// Register with Eureka
// 	instance := fargo.Instance{
// 		InstanceId:       util.GetLocalIP(),
// 		HostName:         util.GetUUID(),
// 		App:              "cart-service",
// 		Port:             9090,
// 		IPAddr:           ipaddress,
// 		VipAddress:       "cart-service",
// 		SecureVipAddress: "cart-service",
// 		HealthCheckUrl:   "http://" + ipaddress + ":9090/health",
// 		StatusPageUrl:    "http://" + ipaddress + ":9090/status",
// 		HomePageUrl:      "http://" + ipaddress + ":9090",
// 	}
// 	eurekaConnection := fargo.NewConn("http://" + serviceDiscoveryIp + ":8167/eureka")
// 	err := eurekaConnection.RegisterInstance(&instance)
// 	if err != nil {
// 		fmt.Println("disini")
// 		fmt.Println(err.Error())
// 	}
// }
