package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"pingshow/internal/adapter/in/grpc/handler"
	"pingshow/internal/adapter/out/repository"
	"pingshow/internal/core/domain/service"
	pb "pingshow/pkg/proto"
)

func main() {
	playerRepository := repository.NewPlayerRepository()
	playerService := service.NewPlayerService(playerRepository)
	playerHandler := handler.NewPlayerHandler(playerService)

	port := 8888
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("ไม่สามารถเปิดพอร์ต %v: %v", port, err)
	}

	s := grpc.NewServer()
	pb.RegisterPlayerServiceServer(s, playerHandler)
	reflection.Register(s)

	log.Printf("Player Service กำลังทำงานที่พอร์ต %d", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("ไม่สามารถให้บริการได้: %v", err)
	}
}