package model

import (
	"time"
)

// MatchEvent เป็นเหตุการณ์ที่เกิดขึ้นในการแข่งขัน
type MatchEvent struct {
	Time      time.Time
	Player    string
	Power     int32
	Goroutine string
	Duration  int64
	EventType string
}

// Match เป็นข้อมูลการแข่งขัน
type Match struct {
	MatchNumber  int32
	Events       []MatchEvent
	CurrentPower int32
	IsGameOver   bool
}

// NewMatch สร้างการแข่งขันใหม่
func NewMatch(matchNumber int32) *Match {
	return &Match{
		MatchNumber:  matchNumber,
		Events:       make([]MatchEvent, 0),
		CurrentPower: 0,
		IsGameOver:   false,
	}
}

// AddEvent เพิ่มเหตุการณ์ในการแข่งขัน
func (m *Match) AddEvent(event MatchEvent) {
	m.Events = append(m.Events, event)
}

// SetCurrentPower ตั้งค่าพลังปัจจุบัน
func (m *Match) SetCurrentPower(power int32) {
	m.CurrentPower = power
}

// EndGame จบเกม
func (m *Match) EndGame() {
	m.IsGameOver = true
}