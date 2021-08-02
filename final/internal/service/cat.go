package service

import (
    "context"

    pb "shopp/api/shop/v1"
)

type CatService struct {
	pb.UnimplementedCatServer
}

func NewCatService() *CatService {
	return &CatService{}
}

func (s *CatService) CreateCat(ctx context.Context, req *pb.CreateCatRequest) (*pb.CreateCatReply, error) {
	return &pb.CreateCatReply{}, nil
}
func (s *CatService) UpdateCat(ctx context.Context, req *pb.UpdateCatRequest) (*pb.UpdateCatReply, error) {
	return &pb.UpdateCatReply{}, nil
}
func (s *CatService) DeleteCat(ctx context.Context, req *pb.DeleteCatRequest) (*pb.DeleteCatReply, error) {
	return &pb.DeleteCatReply{}, nil
}
func (s *CatService) GetCat(ctx context.Context, req *pb.GetCatRequest) (*pb.GetCatReply, error) {
	return &pb.GetCatReply{}, nil
}
func (s *CatService) ListCat(ctx context.Context, req *pb.ListCatRequest) (*pb.ListCatReply, error) {
	return &pb.ListCatReply{}, nil
}
