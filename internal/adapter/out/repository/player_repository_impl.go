package repository

import (
	"context"
	"fmt"
	"pingshow/internal/core/domain/model"
	"sync"
)

// PlayerRepositoryImpl เป็น adapter ขาออกสำหรับการจัดเก็บข้อมูลผู้เล่น
type PlayerRepositoryImpl struct {
	mu      sync.Mutex
	players map[string]*model.Player
}

// NewPlayerRepository สร้าง repository ใหม่
func NewPlayerRepository() *PlayerRepositoryImpl {
	// สร้างข้อมูลผู้เล่นตัวอย่าง
	players := map[string]*model.Player{
		"A": {
			PlayerID:   "A",
			Name:       "Player A",
			TotalGames: 10,
			Wins:       5,
		},
		"B": {
			PlayerID:   "B",
			Name:       "Player B",
			TotalGames: 8,
			Wins:       3,
		},
	}
	
	return &PlayerRepositoryImpl{
		players: players,
	}
}

// GetPlayer ดึงข้อมูลผู้เล่น
func (r *PlayerRepositoryImpl) GetPlayer(ctx context.Context, playerID string) (*model.Player, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	player, ok := r.players[playerID]
	if !ok {
		return nil, fmt.Errorf("ไม่พบผู้เล่น ID: %s", playerID)
	}
	
	return player, nil
}

// SavePlayer บันทึกข้อมูลผู้เล่น
func (r *PlayerRepositoryImpl) SavePlayer(ctx context.Context, player *model.Player) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	r.players[player.PlayerID] = player
	return nil
}