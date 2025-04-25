package service

import (
	"context"
	"fmt"
	"math/rand"
	"pingshow/internal/core/domain/model"
	"pingshow/internal/core/port/out"
	"time"
)

// MatchService จัดการการแข่งขัน
type MatchService struct {
	matchRepository out.MatchRepository
	playerAChan     chan int32
	playerBChan     chan int32
	match           *model.Match
}

// NewMatchService สร้าง service ใหม่
func NewMatchService(matchRepository out.MatchRepository) *MatchService {
	return &MatchService{
		matchRepository: matchRepository,
		playerAChan:     make(chan int32, 1),
		playerBChan:     make(chan int32, 1),
	}
}

// StartNewMatch เริ่มการแข่งขันใหม่
func (s *MatchService) StartNewMatch(ctx context.Context) (*model.Match, <-chan model.MatchEvent, error) {
	// เพิ่มหมายเลขการแข่งขัน
	matchNumber, err := s.matchRepository.IncrementMatchNumber(ctx)
	if err != nil {
		return nil, nil, err
	}

	// สร้างการแข่งขันใหม่
	s.match = model.NewMatch(matchNumber)
	
	// สร้างช่องทางสำหรับส่งเหตุการณ์
	eventChan := make(chan model.MatchEvent, 100)
	
	// เริ่มเกม
	go s.runGame(ctx, eventChan)
	
	return s.match, eventChan, nil
}

// runGame ดำเนินการแข่งขัน
func (s *MatchService) runGame(ctx context.Context, eventChan chan<- model.MatchEvent) {
	defer close(eventChan)
	
	// สุ่มพลังเริ่มต้นของผู้เล่น A
	initialPower := int32(rand.Intn(50) + 50) // 50-100
	s.match.SetCurrentPower(initialPower)
	
	// ผู้เล่น A เริ่มเกม
	go s.playerA(ctx)
	
	// ส่งพลังเริ่มต้นให้ผู้เล่น A
	s.playerAChan <- initialPower
	
	// เริ่มผู้เล่น B
	go s.playerB(ctx)
	
	// รอจนกว่าเกมจะจบ
	for !s.match.IsGameOver {
		select {
		case <-ctx.Done():
			// ยกเลิกเกมเมื่อบริบทถูกยกเลิก
			event := model.MatchEvent{
				Time:      time.Now(),
				EventType: "game_cancelled",
				Player:    "System",
			}
			s.match.AddEvent(event)
			eventChan <- event
			return
		case <-time.After(time.Millisecond * 100):
			// ตรวจสอบสถานะเกม
			continue
		}
	}
	
	// ส่งเหตุการณ์จบเกม
	event := model.MatchEvent{
		Time:      time.Now(),
		EventType: "game_over",
		Player:    "System",
	}
	s.match.AddEvent(event)
	eventChan <- event
	
	// บันทึกการแข่งขัน
	s.matchRepository.SaveMatch(ctx, s.match)
}

// playerA จำลองการเล่นของผู้เล่น A
func (s *MatchService) playerA(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case power := <-s.playerAChan:
			// สร้างเหตุการณ์
			event := model.MatchEvent{
				Time:      time.Now(),
				Player:    "A",
				Power:     power,
				Goroutine: fmt.Sprintf("A-%d", time.Now().UnixNano()),
				Duration:  time.Now().UnixMilli(),
				EventType: "ping",
			}
			
			// เพิ่มเหตุการณ์ในการแข่งขัน
			s.match.AddEvent(event)
			
			// จำลองการตอบสนองของโต๊ะ (ลดพลัง 10-30%)
			reduction := float64(rand.Intn(20) + 10) / 100.0
			tablePower := int32(float64(power) * (1.0 - reduction))
			s.match.SetCurrentPower(tablePower)
			
			// ส่งพลังให้ผู้เล่น B
			time.Sleep(time.Millisecond * 200) // จำลองการหน่วงเวลา
			s.playerBChan <- tablePower
		}
	}
}

// playerB จำลองการเล่นของผู้เล่น B
func (s *MatchService) playerB(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case tablePower := <-s.playerBChan:
			// สุ่มพลังของผู้เล่น B
			bPower := int32(rand.Intn(100) + 1) // 1-100
			
			// สร้างเหตุการณ์
			event := model.MatchEvent{
				Time:      time.Now(),
				Player:    "B",
				Power:     bPower,
				Goroutine: fmt.Sprintf("B-%d", time.Now().UnixNano()),
				Duration:  time.Now().UnixMilli(),
				EventType: "pong",
			}
			
			// เพิ่มเหตุการณ์ในการแข่งขัน
			s.match.AddEvent(event)
			
			// ตรวจสอบว่าผู้เล่น B ชนะหรือแพ้
			if bPower > tablePower {
				// ผู้เล่น B ชนะรอบนี้ ส่งพลังให้ผู้เล่น A
				time.Sleep(time.Millisecond * 200) // จำลองการหน่วงเวลา
				s.playerAChan <- bPower
			} else {
				// ผู้เล่น B แพ้ จบเกม
				s.match.EndGame()
				
				// สร้างเหตุการณ์แพ้
				loseEvent := model.MatchEvent{
					Time:      time.Now(),
					Player:    "B",
					Power:     bPower,
					Goroutine: fmt.Sprintf("B-%d", time.Now().UnixNano()),
					Duration:  time.Now().UnixMilli(),
					EventType: "lose",
				}
				s.match.AddEvent(loseEvent)
				return
			}
		}
	}
}

// GetCurrentMatch ดึงข้อมูลการแข่งขันปัจจุบันสำหรับแสดงผล
func (s *MatchService) GetCurrentMatch() *struct {
	Rally         int32
	Player1Score  int32
	Player2Score  int32
} {
	if s.match == nil {
		return nil
	}
	
	// นับจำนวนเหตุการณ์ทั้งหมดเป็น Rally
	rally := int32(len(s.match.Events))
	
	// นับคะแนนจากเหตุการณ์
	var player1Score, player2Score int32
	for _, event := range s.match.Events {
		if event.EventType == "ping" {
			player1Score++
		} else if event.EventType == "pong" {
			player2Score++
		}
	}
	
	return &struct {
		Rally         int32
		Player1Score  int32
		Player2Score  int32
	}{
		Rally:         rally,
		Player1Score:  player1Score,
		Player2Score:  player2Score,
	}
}