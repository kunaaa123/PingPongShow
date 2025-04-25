package handler

import (
	"context"
	"pingshow/internal/adapter/in/grpc/mapper"
	"pingshow/internal/core/port/in"
	pb "pingshow/pkg/proto"
)

// PlayerHandler เป็น adapter ขาเข้าสำหรับ gRPC
type PlayerHandler struct {
	pb.UnimplementedPlayerServiceServer
	playerUseCase in.PlayerUseCase
}

// NewPlayerHandler สร้าง handler ใหม่
func NewPlayerHandler(playerUseCase in.PlayerUseCase) *PlayerHandler {
	return &PlayerHandler{
		playerUseCase: playerUseCase,
	}
}

// GetPlayerInfo ดึงข้อมูลผู้เล่น
func (h *PlayerHandler) GetPlayerInfo(ctx context.Context, req *pb.PlayerRequest) (*pb.PlayerInfo, error) {
	player, err := h.playerUseCase.GetPlayerInfo(ctx, req.PlayerId)
	if err != nil {
		return nil, err
	}
	
	return mapper.MapToProtoPlayerInfo(player), nil
}

// UpdatePlayerStats อัปเดตสถิติผู้เล่น
func (h *PlayerHandler) UpdatePlayerStats(ctx context.Context, req *pb.PlayerStats) (*pb.Empty, error) {
	playerID, win, powerUsed := mapper.MapFromProtoPlayerStats(req)
	
	err := h.playerUseCase.UpdatePlayerStats(ctx, playerID, win, powerUsed)
	if err != nil {
		return nil, err
	}
	
	return &pb.Empty{}, nil
}