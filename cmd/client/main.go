package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "pingshow/pkg/proto"
)

func main() {
	// เชื่อมต่อกับ Table Service
	tableConn, err := grpc.Dial("localhost:8889", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("ไม่สามารถเชื่อมต่อกับ Table Service: %v", err)
	}
	defer tableConn.Close()
	

	// เชื่อมต่อกับ Player Service
	playerConn, err := grpc.Dial("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("ไม่สามารถเชื่อมต่อกับ Player Service: %v", err)
	}
	defer playerConn.Close()
	playerClient := pb.NewPlayerServiceClient(playerConn)

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()


	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh
		cancel()
	}()

	// สร้างไฟล์ CSV สำหรับบันทึกข้อมูล
	csvFile, err := os.Create("match_log.csv")
	if err != nil {
		log.Fatalf("ไม่สามารถสร้างไฟล์ CSV: %v", err)
	}
	defer csvFile.Close()


	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	// เขียนส่วนหัวของไฟล์ CSV
	csvWriter.Write([]string{"Time", "Event", "Player", "Power", "Goroutine", "Match Number", "Turn"})

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				// เริ่มเกมใหม่ผ่าน Player Service
				latestMatch, err := playerClient.GetPlayerInfo(ctx, &pb.PlayerRequest{PlayerId: "A"})
				if err != nil {
					log.Printf("ไม่สามารถดึงข้อมูลผู้เล่น A ได้: %v", err)
					return
				}
				
				
				matchNumber := latestMatch.TotalGames + 1
				
				fmt.Printf("\n========================================\n")
				fmt.Printf("🎮 เริ่มเกมใหม่ หมายเลขแมตช์: %d\n", matchNumber)
				fmt.Printf("========================================\n\n")
				// ลบบรรทัดนี้: matchNumber := latestMatch.TotalGames
				
				// เริ่ม goroutine ของผู้เล่น A และ B
				playerAChan := make(chan int32, 1)
				playerBChan := make(chan int32, 1)
				
				// บันทึกข้อมูลการเริ่มเกม
				currentTime := time.Now().Format(time.RFC3339)
				csvWriter.Write([]string{
					currentTime,
					"start_game",
					"System",
					"0",
					"main",
					strconv.Itoa(int(matchNumber)),
					"0",
				})
				
				// เริ่ม goroutine ของผู้เล่น A
				go func(ctx context.Context, matchNum int32) {
					goroutineA := fmt.Sprintf("A-%d", time.Now().UnixNano())
					turn := 1
					
					// สุ่มพลังเริ่มต้น
					initialPower := int32(rand.Intn(50) + 50) // 50-100
					
					// บันทึกข้อมูลการเริ่ม goroutine A
					currentTime := time.Now().Format(time.RFC3339)
					csvWriter.Write([]string{
						currentTime,
						"wake_up",
						"A",
						"0",
						goroutineA,
						strconv.Itoa(int(matchNum)),
						strconv.Itoa(turn),
					})
					
					// ส่ง ping ครั้งแรกพร้อมพลังเริ่มต้น
					fmt.Printf("⏰ %s | 🎮 A | 💪 %d | 🧵 %s | #%d | รอบที่ %d\n",
						currentTime, initialPower, goroutineA, matchNum, turn)
					
					// บันทึกข้อมูลการส่ง ping
					csvWriter.Write([]string{
						currentTime,
						"ping",
						"A",
						strconv.Itoa(int(initialPower)),
						goroutineA,
						strconv.Itoa(int(matchNum)),
						strconv.Itoa(turn),
					})
					
					// ส่งพลังให้ Table Service
					playerAChan <- initialPower
					
					// รอรับพลังจาก B
					for {
						select {
						case <-ctx.Done():
							return
						case power := <-playerAChan:
							turn += 2 // เพิ่มรอบการเล่น (A เล่นรอบคี่)
							
							// บันทึกข้อมูลการตื่นของ A
							currentTime = time.Now().Format(time.RFC3339)
							csvWriter.Write([]string{
								currentTime,
								"wake_up",
								"A",
								strconv.Itoa(int(power)),
								goroutineA,
								strconv.Itoa(int(matchNum)),
								strconv.Itoa(turn),
							})
							
							// สุ่มพลังใหม่
							newPower := int32(rand.Intn(50) + 50) // 50-100
							
							fmt.Printf("⏰ %s | 🎮 A | 💪 %d | 🧵 %s | #%d | รอบที่ %d\n",
								currentTime, newPower, goroutineA, matchNum, turn)
							
							// บันทึกข้อมูลการส่ง ping
							csvWriter.Write([]string{
								currentTime,
								"ping",
								"A",
								strconv.Itoa(int(newPower)),
								goroutineA,
								strconv.Itoa(int(matchNum)),
								strconv.Itoa(turn),
							})
							
							// ส่งพลังให้ Table Service
							playerAChan <- newPower
						}
					}
				}(ctx, matchNumber)
				
				// เริ่ม goroutine ของผู้เล่น B
				gameFinished := make(chan struct{}, 1) // สร้าง channel สำหรับแจ้งเมื่อเกมจบ
				
				go func(ctx context.Context, matchNum int32) {
					goroutineB := fmt.Sprintf("B-%d", time.Now().UnixNano())
					turn := 2
					
					// บันทึกข้อมูลการเริ่ม goroutine B
					currentTime := time.Now().Format(time.RFC3339)
					csvWriter.Write([]string{
						currentTime,
						"wake_up",
						"B",
						"0",
						goroutineB,
						strconv.Itoa(int(matchNum)),
						strconv.Itoa(turn),
					})
					
					for {
						select {
						case <-ctx.Done():
							return
						case tablePower := <-playerBChan:
							// บันทึกข้อมูลการตื่นของ B
							currentTime = time.Now().Format(time.RFC3339)
							csvWriter.Write([]string{
								currentTime,
								"wake_up",
								"B",
								strconv.Itoa(int(tablePower)),
								goroutineB,
								strconv.Itoa(int(matchNum)),
								strconv.Itoa(turn),
							})
							
							// สุ่มพลังของ B
							bPower := int32(rand.Intn(100) + 1) // 1-100
							
							fmt.Printf("⏰ %s | 🎮 B | 💪 %d | 🧵 %s | #%d | รอบที่ %d\n",
								currentTime, bPower, goroutineB, matchNum, turn)
							
							// บันทึกข้อมูลการส่ง pong
							csvWriter.Write([]string{
								currentTime,
								"pong",
								"B",
								strconv.Itoa(int(bPower)),
								goroutineB,
								strconv.Itoa(int(matchNum)),
								strconv.Itoa(turn),
							})
							
							// ตรวจสอบว่า B ชนะหรือแพ้
							if bPower > tablePower {
								// B ชนะรอบนี้ ส่งพลังให้ A
								playerAChan <- bPower
								turn += 2 // เพิ่มรอบการเล่น (B เล่นรอบคู่)
							} else {
								// B แพ้ จบเกม
								currentTime = time.Now().Format(time.RFC3339)
								fmt.Printf("⏰ %s | 🎮 B | 💥 แพ้! | 🧵 %s | #%d | รอบที่ %d\n",
									currentTime, goroutineB, matchNum, turn)
								
								// บันทึกข้อมูลการแพ้
								csvWriter.Write([]string{
									currentTime,
									"lose",
									"B",
									strconv.Itoa(int(bPower)),
									goroutineB,
									strconv.Itoa(int(matchNum)),
									strconv.Itoa(turn),
								})
								
								// บันทึกข้อมูลการจบเกม
								csvWriter.Write([]string{
									currentTime,
									"game_over",
									"System",
									"0",
									"main",
									strconv.Itoa(int(matchNum)),
									strconv.Itoa(turn),
								})
								
								// อัปเดตสถิติผู้เล่น A (ชนะ)
								_, err = playerClient.UpdatePlayerStats(ctx, &pb.PlayerStats{
									PlayerId:  "A",
									Win:       true,
									PowerUsed: 50,
								})
								if err != nil {
									log.Printf("ไม่สามารถอัปเดตสถิติผู้เล่น A ได้: %v", err)
								}
								
								// อัปเดตสถิติผู้เล่น B (แพ้)
								_, err = playerClient.UpdatePlayerStats(ctx, &pb.PlayerStats{
									PlayerId:  "B",
									Win:       false,
									PowerUsed: 50,
								})
								if err != nil {
									log.Printf("ไม่สามารถอัปเดตสถิติผู้เล่น B ได้: %v", err)
								}
								
								// ส่งสัญญาณว่าเกมจบแล้ว
								gameFinished <- struct{}{}
								return
							}
						}
					}
				}(ctx, matchNumber)
				
				// จำลอง Table Service
				go func(ctx context.Context, matchNum int32) {
					for {
						select {
						case <-ctx.Done():
							return
						case power := <-playerAChan:
							// จำลองการตอบสนองของโต๊ะ (ลดพลัง 10-30%)
							reduction := float64(rand.Intn(20) + 10) / 100.0
							tablePower := int32(float64(power) * (1.0 - reduction))
							
							// บันทึกข้อมูลการตอบสนองของโต๊ะ
							currentTime := time.Now().Format(time.RFC3339)
							csvWriter.Write([]string{
								currentTime,
								"table_response",
								"Table",
								strconv.Itoa(int(tablePower)),
								"Table",
								strconv.Itoa(int(matchNum)),
								"0",
							})
							
							fmt.Printf("⏰ %s | 🎮 Table | 💪 %d | 🧵 Table | #%d\n",
								currentTime, tablePower, matchNum)
							
							// ส่งพลังให้ผู้เล่น B
							time.Sleep(time.Millisecond * 200) // จำลองการหน่วงเวลา
							playerBChan <- tablePower
						}
					}
				}(ctx, matchNumber)
				
				// รอให้เกมจบ
				select {
				case <-ctx.Done():
					return
				case <-gameFinished:
					fmt.Println("\n👋 เกมจบแล้ว B แพ้! กำลังเริ่มเกมใหม่...")
					// เกมจบแล้ว เริ่มเกมใหม่อัตโนมัติ
					csvWriter.Flush() // บันทึกข้อมูลลงไฟล์ CSV
					time.Sleep(2 * time.Second) // รอสักครู่ก่อนเริ่มเกมใหม่
				}
			}
		}
	}()

	wg.Wait()
}