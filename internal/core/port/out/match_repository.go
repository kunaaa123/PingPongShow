package out

import (
	"context"
	"pingshow/internal/core/domain/model"
)

// MatchRepository เป็น port ขาออกสำหรับการจัดเก็บข้อมูลการแข่งขัน
type MatchRepository interface {
	// GetLatestMatchNumber ดึงหมายเลขการแข่งขันล่าสุด
	GetLatestMatchNumber(ctx context.Context) (int32, error)
	
	// IncrementMatchNumber เพิ่มหมายเลขการแข่งขัน
	IncrementMatchNumber(ctx context.Context) (int32, error)
	
	// SaveMatch บันทึกข้อมูลการแข่งขัน
	SaveMatch(ctx context.Context, match *model.Match) error
}