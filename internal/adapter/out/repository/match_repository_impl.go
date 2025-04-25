package repository

import (
	"context"
	"pingshow/internal/core/domain/model"
	"sync"
)

// MatchRepositoryImpl เป็น adapter ขาออกสำหรับการจัดเก็บข้อมูลการแข่งขัน
type MatchRepositoryImpl struct {
	mu           sync.Mutex
	matchNumber  int32
	matchHistory map[int32]*model.Match
}

// NewMatchRepository สร้าง repository ใหม่
func NewMatchRepository() *MatchRepositoryImpl {
	return &MatchRepositoryImpl{
		matchNumber:  0,
		matchHistory: make(map[int32]*model.Match),
	}
}

// GetLatestMatchNumber ดึงหมายเลขการแข่งขันล่าสุด
func (r *MatchRepositoryImpl) GetLatestMatchNumber(ctx context.Context) (int32, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.matchNumber, nil
}

// IncrementMatchNumber เพิ่มหมายเลขการแข่งขัน
func (r *MatchRepositoryImpl) IncrementMatchNumber(ctx context.Context) (int32, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.matchNumber++
	return r.matchNumber, nil
}

// SaveMatch บันทึกข้อมูลการแข่งขัน
func (r *MatchRepositoryImpl) SaveMatch(ctx context.Context, match *model.Match) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.matchHistory[match.MatchNumber] = match
	return nil
}