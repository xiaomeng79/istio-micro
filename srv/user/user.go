package user

import (
	"context"

	pb "github.com/xiaomeng79/istio-micro/srv/user/proto"

	"github.com/jinzhu/copier"
	"github.com/xiaomeng79/go-log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct{}

func (s *Server) UserAdd(ctx context.Context, in *pb.UserBase) (out *pb.UserBase, outerr error) {
	m := new(User)
	out = new(pb.UserBase)
	err := copier.Copy(m, in)
	if err != nil {
		log.Error(err.Error(), ctx)
		outerr = status.Error(codes.Internal, err.Error())
		return
	}
	err = m.Add(ctx)
	if err != nil {
		outerr = status.Error(codes.InvalidArgument, err.Error())
		return
	}
	err = copier.Copy(out, m)
	if err != nil {
		log.Error(err.Error(), ctx)
		outerr = status.Error(codes.Internal, err.Error())
		return
	}
	return
}

func (s *Server) UserUpdate(ctx context.Context, in *pb.UserBase) (out *pb.UserBase, outerr error) {
	m := new(User)
	out = new(pb.UserBase)
	err := copier.Copy(m, in)
	if err != nil {
		log.Error(err.Error(), ctx)
		outerr = status.Error(codes.Internal, err.Error())
		return
	}
	err = m.Update(ctx)
	if err != nil {
		outerr = status.Error(codes.InvalidArgument, err.Error())
		return
	}
	err = copier.Copy(out, m)
	if err != nil {
		log.Error(err.Error(), ctx)
		outerr = status.Error(codes.Internal, err.Error())
		return
	}
	return
}

func (s *Server) UserDelete(ctx context.Context, in *pb.UserID) (out *pb.UserID, outerr error) {
	m := new(User)
	out = new(pb.UserID)
	err := copier.Copy(m, in)
	if err != nil {
		log.Error(err.Error(), ctx)
		outerr = status.Error(codes.Internal, err.Error())
		return
	}
	err = m.Delete(ctx)
	if err != nil {
		outerr = status.Error(codes.InvalidArgument, err.Error())
		return
	}
	err = copier.Copy(out, m)
	if err != nil {
		log.Error(err.Error(), ctx)
		outerr = status.Error(codes.Internal, err.Error())
		return
	}
	return
}

func (s *Server) UserQueryOne(ctx context.Context, in *pb.UserID) (out *pb.UserBase, outerr error) {
	m := new(User)
	out = new(pb.UserBase)
	err := copier.Copy(m, in)
	if err != nil {
		log.Error(err.Error(), ctx)
		outerr = status.Error(codes.Internal, err.Error())
		return
	}
	err = m.QueryOne(ctx)
	if err != nil {
		outerr = status.Error(codes.InvalidArgument, err.Error())
		return
	}
	err = copier.Copy(out, m)
	if err != nil {
		log.Error(err.Error(), ctx)
		outerr = status.Error(codes.Internal, err.Error())
		return
	}
	return
}

func (s *Server) UserQueryAll(ctx context.Context, in *pb.UserAllOption) (*pb.UserAll, error) {
	m := new(User)
	err := copier.Copy(m, in)
	if err != nil {
		log.Error(err.Error(), ctx)
		return &pb.UserAll{}, status.Error(codes.Internal, err.Error())
	}
	ms, page, err := m.QueryAll(ctx)
	if err != nil {
		return &pb.UserAll{}, status.Error(codes.InvalidArgument, err.Error())
	}
	var agt []*pb.UserBase
	err = copier.Copy(&agt, ms)
	if err != nil {
		log.Error(err.Error(), ctx)
		return &pb.UserAll{}, status.Error(codes.Internal, err.Error())
	}
	_page := new(pb.Page)
	err = copier.Copy(_page, &page)
	if err != nil {
		log.Error(err.Error(), ctx)
		return &pb.UserAll{}, status.Error(codes.Internal, err.Error())
	}
	return &pb.UserAll{All: agt, Page: _page}, nil
}
