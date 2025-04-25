package mapper

import (
	"pingshow/internal/core/domain/model"
	pb "pingshow/pkg/proto"
)

// MapToProtoMatchEvent แปลง domain model เป็น protobuf message
func MapToProtoMatchEvent(event model.MatchEvent) *pb.MatchEvent {
	return &pb.MatchEvent{
		Time:        event.Time.Format("15:04:05.000"),
		Player:      event.Player,
		Power:       event.Power,
		Goroutine:   event.Goroutine,
		MatchNumber: 0, // จะถูกกำหนดโดย handler
		Duration:    event.Duration,
		EventType:   event.EventType,
	}
}

// MapToProtoPlayerInfo แปลง domain model เป็น protobuf message
func MapToProtoPlayerInfo(player *model.Player) *pb.PlayerInfo {
	return &pb.PlayerInfo{
		PlayerId:   player.PlayerID,
		Name:       player.Name,
		TotalGames: player.TotalGames,
		Wins:       player.Wins,
	}
}

// MapFromProtoPlayerStats แปลง protobuf message เป็น domain model
func MapFromProtoPlayerStats(stats *pb.PlayerStats) (string, bool, int32) {
	return stats.PlayerId, stats.Win, stats.PowerUsed
}