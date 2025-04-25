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
	matchRepository := repository.NewMatchRepository()
	matchService := service.NewMatchService(matchRepository)
	matchHandler := handler.NewMatchHandler(matchService)

	port := 8889
	lis, err := net.Listen("tcp", ":8889")
	if err != nil {
		log.Fatalf("ไม่สามารถเปิดพอร์ต %v: %v", port, err)
	}

	s := grpc.NewServer()
	pb.RegisterTableServiceServer(s, matchHandler)
	reflection.Register(s)

	log.Printf("Table Service กำลังทำงานที่พอร์ต %d", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("ไม่สามารถให้บริการได้: %v", err)
	}
}