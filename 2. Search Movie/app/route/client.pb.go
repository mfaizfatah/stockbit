package route

import (
	"fmt"
	"log"
	"net"
	"project/app/config"
	"project/app/delivery/pb"
	"project/app/usecase"

	"google.golang.org/grpc"
)

func PbClient(uc usecase.Usecase) {
	s := grpc.NewServer()
	pb.RegisterSearchServiceServer(s, &pb.Transport{Uc: uc})
	lis, err := net.Listen("tcp", config.Server.PORTPROTO)
	if err != nil {
		panic(err)
	}

	log.Println(fmt.Sprintf("%v 🚀 GRPC server started... ", config.Server.PORTPROTO))
	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
