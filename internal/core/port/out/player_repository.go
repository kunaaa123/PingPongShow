package out

import (
	"context"
	"pingshow/internal/core/domain/model"
)

// PlayerRepository เป็น port ขาออกสำหรับการจัดเก็บข้อมูลผู้เล่น
type PlayerRepository interface {
	// GetPlayer ดึงข้อมูลผู้เล่น
	GetPlayer(ctx context.Context, playerID string) (*model.Player, error)
	
	// SavePlayer บันทึกข้อมูลผู้เล่น
	SavePlayer(ctx context.Context, player *model.Player) error
}