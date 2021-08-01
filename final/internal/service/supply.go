package service

import (
	"context"

	pb "shopp/api/shop/v1"
)

type SupplyService struct {
	pb.UnimplementedSupplyServer
}

func NewSupplyService() *SupplyService {
	return &SupplyService{}
}

func (s *SupplyService) CreateSupply(ctx context.Context, req *pb.CreateSupplyRequest) (*pb.CreateSupplyReply, error) {
	return &pb.CreateSupplyReply{}, nil
}
func (s *SupplyService) UpdateSupply(ctx context.Context, req *pb.UpdateSupplyRequest) (*pb.UpdateSupplyReply, error) {
	return &pb.UpdateSupplyReply{}, nil
}
func (s *SupplyService) DeleteSupply(ctx context.Context, req *pb.DeleteSupplyRequest) (*pb.DeleteSupplyReply, error) {
	return &pb.DeleteSupplyReply{}, nil
}
func (s *SupplyService) GetSupply(ctx context.Context, req *pb.GetSupplyRequest) (*pb.GetSupplyReply, error) {
	return &pb.GetSupplyReply{}, nil
}
func (s *SupplyService) ListSupply(ctx context.Context, req *pb.ListSupplyRequest) (*pb.ListSupplyReply, error) {
	return &pb.ListSupplyReply{}, nil
}
