package service

import (
	"context"
	"pingshow/internal/core/domain/model"
	"pingshow/internal/core/port/out"
)

// PlayerService จัดการข้อมูลผู้เล่น
type PlayerService struct {
	playerRepository out.PlayerRepository
}

// NewPlayerService สร้าง service ใหม่
func NewPlayerService(playerRepository out.PlayerRepository) *PlayerService {
	return &PlayerService{
		playerRepository: playerRepository,
	}
}

// GetPlayerInfo ดึงข้อมูลผู้เล่น
func (s *PlayerService) GetPlayerInfo(ctx context.Context, playerID string) (*model.Player, error) {
	return s.playerRepository.GetPlayer(ctx, playerID)
}

// UpdatePlayerStats อัปเดตสถิติผู้เล่น
func (s *PlayerService) UpdatePlayerStats(ctx context.Context, playerID string, win bool, powerUsed int32) error {
	// ดึงข้อมูลผู้เล่น
	player, err := s.playerRepository.GetPlayer(ctx, playerID)
	if err != nil {
		return err
	}
	
	// อัปเดตสถิติ
	player.UpdateStats(win)
	
	// บันทึกข้อมูลผู้เล่น
	return s.playerRepository.SavePlayer(ctx, player)
}