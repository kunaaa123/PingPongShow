package in

import (
	"context"
	"pingshow/internal/core/domain/model"
)

// MatchUseCase เป็น port ขาเข้าสำหรับการจัดการการแข่งขัน
type MatchUseCase interface {
	// StartMatch เริ่มการแข่งขันใหม่และส่งเหตุการณ์ผ่าน channel
	StartMatch(ctx context.Context, newGame bool) (<-chan model.MatchEvent, error)
	
	// GetLatestMatch ดึงข้อมูลการแข่งขันล่าสุด
	GetLatestMatch(ctx context.Context) (int32, error)
}