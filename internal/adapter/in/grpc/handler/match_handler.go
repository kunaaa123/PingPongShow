package handler

import (
	"pingshow/internal/adapter/in/grpc/mapper"
	"pingshow/internal/core/domain/service"
	pb "pingshow/pkg/proto"
)


type MatchHandler struct {
	pb.UnimplementedTableServiceServer
	matchService *service.MatchService
}


func NewMatchHandler(matchService *service.MatchService) *MatchHandler {
	return &MatchHandler{
		matchService: matchService,
	}
}


func (h *MatchHandler) StartMatch(req *pb.StartRequest, stream pb.TableService_StartMatchServer) error {
	
	match, eventChan, err := h.matchService.StartNewMatch(stream.Context())
	if err != nil {
		return err
	}
	
	// ส่งเหตุการณ์กลับไปยังไคลเอนต์
	for event := range eventChan {
		
		protoEvent := mapper.MapToProtoMatchEvent(event)
		protoEvent.MatchNumber = match.MatchNumber
		
		// ส่งเหตุการณ์
		if err := stream.Send(protoEvent); err != nil {
			return err
		}
	}
	
	return nil
}