package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"pingshow/internal/core/domain/model"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLRepository เป็น adapter ขาออกสำหรับการจัดเก็บข้อมูลลงฐานข้อมูล MySQL
type MySQLRepository struct {
	db *sql.DB
}

// NewMySQLRepository สร้าง repository ใหม่
func NewMySQLRepository(dsn string) (*MySQLRepository, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("ไม่สามารถเชื่อมต่อกับฐานข้อมูล: %w", err)
	}

	// ตรวจสอบการเชื่อมต่อ
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ไม่สามารถ ping ฐานข้อมูล: %w", err)
	}

	// ตั้งค่าการเชื่อมต่อ
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	return &MySQLRepository{db: db}, nil
}

// Close ปิดการเชื่อมต่อกับฐานข้อมูล
func (r *MySQLRepository) Close() error {
	return r.db.Close()
}

// SaveMatchEvents บันทึกเหตุการณ์ทั้งหมดของการแข่งขันลงฐานข้อมูล
func (r *MySQLRepository) SaveMatchEvents(ctx context.Context, match *model.Match) error {
	// เริ่ม transaction
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("ไม่สามารถเริ่ม transaction: %w", err)
	}
	defer tx.Rollback() // จะถูกยกเลิกหากไม่มีการ commit

	// เตรียม statement สำหรับเพิ่มข้อมูลในตาราง all_matches
	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO all_matches (event_time, event_type, player, power, goroutine, match_number, turn)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("ไม่สามารถเตรียม statement: %w", err)
	}
	defer stmt.Close()

	// บันทึกเหตุการณ์ทั้งหมด
	for i, event := range match.Events {
		turn := i + 1
		_, err := stmt.ExecContext(
			ctx,
			event.Time,
			event.EventType,
			event.Player,
			event.Power,
			event.Goroutine,
			match.MatchNumber,
			turn,
		)
		if err != nil {
			return fmt.Errorf("ไม่สามารถบันทึกเหตุการณ์: %w", err)
		}
	}

	// อัปเดตแมตช์ล่าสุด
	_, err = tx.ExecContext(ctx, `
		UPDATE last_played_match SET match_number = ? WHERE id = 1
	`, match.MatchNumber)
	if err != nil {
		return fmt.Errorf("ไม่สามารถอัปเดตแมตช์ล่าสุด: %w", err)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("ไม่สามารถ commit transaction: %w", err)
	}

	return nil
}

// GetLatestMatchNumber ดึงหมายเลขการแข่งขันล่าสุด
func (r *MySQLRepository) GetLatestMatchNumber(ctx context.Context) (int32, error) {
	var matchNumber int32
	err := r.db.QueryRowContext(ctx, "SELECT match_number FROM last_played_match WHERE id = 1").Scan(&matchNumber)
	if err != nil {
		return 0, fmt.Errorf("ไม่สามารถดึงหมายเลขแมตช์ล่าสุด: %w", err)
	}
	return matchNumber, nil
}

// IncrementMatchNumber เพิ่มหมายเลขการแข่งขัน
func (r *MySQLRepository) IncrementMatchNumber(ctx context.Context) (int32, error) {
	// ดึงหมายเลขแมตช์ล่าสุด
	var matchNumber int32
	err := r.db.QueryRowContext(ctx, "SELECT match_number FROM last_played_match WHERE id = 1").Scan(&matchNumber)
	if err != nil {
		return 0, fmt.Errorf("ไม่สามารถดึงหมายเลขแมตช์ล่าสุด: %w", err)
	}

	// เพิ่มหมายเลขแมตช์
	matchNumber++

	// อัปเดตหมายเลขแมตช์ล่าสุด
	_, err = r.db.ExecContext(ctx, "UPDATE last_played_match SET match_number = ? WHERE id = 1", matchNumber)
	if err != nil {
		return 0, fmt.Errorf("ไม่สามารถอัปเดตหมายเลขแมตช์ล่าสุด: %w", err)
	}

	return matchNumber, nil
}

// SaveMatch บันทึกข้อมูลการแข่งขัน
func (r *MySQLRepository) SaveMatch(ctx context.Context, match *model.Match) error {
	return r.SaveMatchEvents(ctx, match)
}

// ImportCSVToDatabase นำเข้าข้อมูลจากไฟล์ CSV ลงฐานข้อมูล
func (r *MySQLRepository) ImportCSVToDatabase(ctx context.Context, csvData [][]string) error {
	// เริ่ม transaction
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("ไม่สามารถเริ่ม transaction: %w", err)
	}
	defer tx.Rollback() // จะถูกยกเลิกหากไม่มีการ commit

	// เตรียม statement สำหรับเพิ่มข้อมูลในตาราง all_matches
	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO all_matches (event_time, event_type, player, power, goroutine, match_number, turn)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("ไม่สามารถเตรียม statement: %w", err)
	}
	defer stmt.Close()

	// ตรวจสอบว่ามีข้อมูลหรือไม่
	if len(csvData) <= 1 {
		return fmt.Errorf("ไม่มีข้อมูลใน CSV")
	}

	// ข้ามแถวแรก (header)
	var latestMatchNumber int = 0
	for i, row := range csvData[1:] {
		// ตรวจสอบว่ามีข้อมูลครบหรือไม่
		if len(row) < 7 {
			log.Printf("ข้อมูลไม่ครบในแถวที่ %d: %v", i+1, row)
			continue
		}

		// แปลงข้อมูล
		eventTime, err := time.Parse(time.RFC3339, row[0])
		if err != nil {
			log.Printf("รูปแบบเวลาไม่ถูกต้องในแถวที่ %d: %v", i+1, err)
			continue
		}

		eventType := row[1]
		player := row[2]

		power := 0
		if row[3] != "" {
			if p, err := parseInt(row[3]); err == nil {
				power = p
			}
		}

		goroutine := row[4]

		matchNumber := 0
		if row[5] != "" {
			if m, err := parseInt(row[5]); err == nil {
				matchNumber = m
				if matchNumber > latestMatchNumber {
					latestMatchNumber = matchNumber
				}
			}
		}

		turn := 0
		if row[6] != "" {
			if t, err := parseInt(row[6]); err == nil {
				turn = t
			}
		}

		// บันทึกข้อมูล
		_, err = stmt.ExecContext(
			ctx,
			eventTime,
			eventType,
			player,
			power,
			goroutine,
			matchNumber,
			turn,
		)
		if err != nil {
			log.Printf("ไม่สามารถบันทึกข้อมูลในแถวที่ %d: %v", i+1, err)
			continue
		}
	}

	// อัปเดตแมตช์ล่าสุด
	_, err = tx.ExecContext(ctx, `
		UPDATE last_played_match SET match_number = ? WHERE id = 1
	`, latestMatchNumber)
	if err != nil {
		return fmt.Errorf("ไม่สามารถอัปเดตแมตช์ล่าสุด: %w", err)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("ไม่สามารถ commit transaction: %w", err)
	}

	return nil
}

// parseInt แปลงสตริงเป็นตัวเลข
func parseInt(s string) (int, error) {
	var i int
	_, err := fmt.Sscanf(s, "%d", &i)
	return i, err
}