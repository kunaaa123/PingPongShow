package model

// Player เป็นข้อมูลผู้เล่น
type Player struct {
	PlayerID   string
	Name       string
	TotalGames int32
	Wins       int32
}

// NewPlayer สร้างผู้เล่นใหม่
func NewPlayer(id string, name string) *Player {
	return &Player{
		PlayerID:   id,
		Name:       name,
		TotalGames: 0,
		Wins:       0,
	}
}

// AddGame เพิ่มจำนวนเกมที่เล่น
// UpdateStats อัปเดตสถิติผู้เล่น
func (p *Player) UpdateStats(win bool) {
	p.TotalGames++
	if win {
		p.Wins++
	}
}