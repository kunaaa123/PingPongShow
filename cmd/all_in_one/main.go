package main

import (
	"log"
	"net"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"pingshow/internal/adapter/in/grpc/handler"
	"pingshow/internal/adapter/out/repository"
	"pingshow/internal/core/domain/service"
	pb "pingshow/pkg/proto"
	"os/exec"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	
	// เชื่อมต่อกับฐานข้อมูล MySQL
	dsn := "root:@tcp(localhost:3306)/pingshow?parseTime=true"
	mysqlRepo, err := repository.NewMySQLRepository(dsn)
	if err != nil {
		log.Fatalf("ไม่สามารถเชื่อมต่อกับฐานข้อมูล MySQL: %v", err)
	}
	defer mysqlRepo.Close()
	
	// Player Service
	go func() {
		playerRepository := repository.NewPlayerRepository()
		playerService := service.NewPlayerService(playerRepository)
		playerHandler := handler.NewPlayerHandler(playerService)

		lis, err := net.Listen("tcp", ":8888")
		if err != nil {
			log.Fatalf("ไม่สามารถเปิดพอร์ต 8888: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterPlayerServiceServer(s, playerHandler)
		reflection.Register(s)
		log.Println("Player Service กำลังทำงานที่พอร์ต 8888")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Player Service error: %v", err)
		}
	}()


	go func() {
	
		matchService := service.NewMatchService(mysqlRepo)
		matchHandler := handler.NewMatchHandler(matchService)

		lis, err := net.Listen("tcp", ":8889")
		if err != nil {
			log.Fatalf("ไม่สามารถเปิดพอร์ต 8889: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterTableServiceServer(s, matchHandler)
		reflection.Register(s)
		log.Println("Table Service กำลังทำงานที่พอร์ต 8889")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Table Service error: %v", err)
		}
	}()

	
	time.Sleep(2 * time.Second)


	go func() {
	
		cmd := exec.Command("go", "run", "./cmd/client/main.go")
		cmd.Stdout = log.Writer()
		cmd.Stderr = log.Writer()
		if err := cmd.Run(); err != nil {
			log.Fatalf("Client error: %v", err)
		}
		wg.Done() 
	}()

	wg.Wait() 
}