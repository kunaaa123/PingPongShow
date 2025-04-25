package in

import (
	"context"
	"pingshow/internal/core/domain/model"
)

// PlayerUseCase เป็น port ขาเข้าสำหรับการจัดการผู้เล่น
type PlayerUseCase interface {
	// GetPlayerInfo ดึงข้อมูลผู้เล่น
	GetPlayerInfo(ctx context.Context, playerID string) (*model.Player, error)
	
	// UpdatePlayerStats อัปเดตสถิติผู้เล่น
	UpdatePlayerStats(ctx context.Context, playerID string, win bool, powerUsed int32) error
}