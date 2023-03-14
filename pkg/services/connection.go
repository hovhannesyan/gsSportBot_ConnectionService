package services

import (
	"context"
	"github.com/hovhannesyan/gsSportBot_ConnectionService/pkg/db"
	"github.com/hovhannesyan/gsSportBot_ConnectionService/pkg/pb"
	"github.com/hovhannesyan/gsSportBot_ConnectionService/pkg/utils"
)

type Server struct {
	DbHandler db.Handler
	pb.UnimplementedConnectionServer
}

func (s *Server) CreateSet(ctx context.Context, req *pb.CreateSetRequest) (*pb.CreateSetResponse, error) {
	dummyElement := "dummy"
	_, err := s.DbHandler.DB.SAdd(ctx, utils.SetKeyToString(req.Set), dummyElement).Result()
	if err != nil {
		return &pb.CreateSetResponse{
			Success: false,
		}, err
	}
	_, err = s.DbHandler.DB.SRem(ctx, utils.SetKeyToString(req.Set), dummyElement).Result()
	if err != nil {
		return &pb.CreateSetResponse{
			Success: false,
		}, err
	}
	return &pb.CreateSetResponse{
		Success: true,
	}, nil
}

func (s *Server) GetSet(ctx context.Context, req *pb.GetSetRequest) (*pb.GetSetResponse, error) {
	res, err := s.DbHandler.DB.SMembers(ctx, utils.SetKeyToString(req.Set)).Result()
	if err != nil {
		return nil, err
	}
	return &pb.GetSetResponse{
		Items: res,
	}, nil
}

func (s *Server) DeleteSet(ctx context.Context, req *pb.DeleteSetRequest) (*pb.DeleteSetResponse, error) {
	_, err := s.DbHandler.DB.Del(ctx, utils.SetKeyToString(req.Set)).Result()
	if err != nil {
		return &pb.DeleteSetResponse{
			Success: false,
		}, err
	}
	return &pb.DeleteSetResponse{
		Success: true,
	}, nil
}

func (s *Server) AddToSet(ctx context.Context, req *pb.AddToSetRequest) (*pb.AddToSetResponse, error) {
	_, err := s.DbHandler.DB.SAdd(ctx, utils.SetKeyToString(req.Set), req.Items).Result()
	if err != nil {
		return &pb.AddToSetResponse{
			Success: false,
		}, err
	}
	return &pb.AddToSetResponse{}, nil
}

func (s *Server) RemoveFromSet(ctx context.Context, req *pb.RemoveFromSetRequest) (*pb.RemoveFromSetResponse, error) {
	_, err := s.DbHandler.DB.SRem(ctx, utils.SetKeyToString(req.Set), req.Items).Result()
	if err != nil {
		return &pb.RemoveFromSetResponse{
			Success: false,
		}, err
	}
	return &pb.RemoveFromSetResponse{}, nil
}
