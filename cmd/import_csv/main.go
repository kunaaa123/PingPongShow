package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"pingshow/internal/adapter/out/repository"
	"time"
)

func main() {
	// รับพารามิเตอร์จาก command line
	csvPath := flag.String("csv", "match_log.csv", "เส้นทางไปยังไฟล์ CSV")
	dsn := flag.String("dsn", "root:@tcp(localhost:3306)/pingshow?parseTime=true", "Data Source Name สำหรับเชื่อมต่อ MySQL")
	flag.Parse()

	// เปิดไฟล์ CSV
	file, err := os.Open(*csvPath)
	if err != nil {
		log.Fatalf("ไม่สามารถเปิดไฟล์ CSV: %v", err)
	}
	defer file.Close()

	// อ่านข้อมูลจาก CSV
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("ไม่สามารถอ่านข้อมูลจาก CSV: %v", err)
	}

	// สร้าง repository
	repo, err := repository.NewMySQLRepository(*dsn)
	if err != nil {
		log.Fatalf("ไม่สามารถสร้าง repository: %v", err)
	}
	defer repo.Close()

	// กำหนดเวลาหมดเวลาสำหรับการนำเข้าข้อมูล
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// นำเข้าข้อมูล
	fmt.Println("กำลังนำเข้าข้อมูลจาก CSV...")
	start := time.Now()
	err = repo.ImportCSVToDatabase(ctx, records)
	if err != nil {
		log.Fatalf("ไม่สามารถนำเข้าข้อมูล: %v", err)
	}

	// แสดงผลลัพธ์
	fmt.Printf("นำเข้าข้อมูลสำเร็จ ใช้เวลา %v\n", time.Since(start))

	// ดึงหมายเลขแมตช์ล่าสุด
	latestMatch, err := repo.GetLatestMatchNumber(ctx)
	if err != nil {
		log.Fatalf("ไม่สามารถดึงหมายเลขแมตช์ล่าสุด: %v", err)
	}
	fmt.Printf("หมายเลขแมตช์ล่าสุด: %d\n", latestMatch)
}