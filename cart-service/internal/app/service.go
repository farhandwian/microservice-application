package app

import (
	"github.com/farhandwian/microservice/internal/service"
	pb "github.com/farhandwian/microservice/pb"
)

type MicroserviceServer struct {
	// ini isinya service
	pb.UnimplementedCartServiceServer
	cartService service.CartService
}

func NewMicroserviceServer(cartService service.CartService) *MicroserviceServer {
	return &MicroserviceServer{cartService: cartService}
}
