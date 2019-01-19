package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/xiaomeng79/go-log"
	pb "github.com/xiaomeng79/istio-micro/srv/user/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct{}

func (s *UserServer) UserAdd(ctx context.Context, in *pb.UserBase) (out *pb.UserBase, outerr error) {
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

func (s *UserServer) UserUpdate(ctx context.Context, in *pb.UserBase) (out *pb.UserBase, outerr error) {
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

func (s *UserServer) UserDelete(ctx context.Context, in *pb.UserId) (out *pb.UserId, outerr error) {
	m := new(User)
	out = new(pb.UserId)
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

func (s *UserServer) UserQueryOne(ctx context.Context, in *pb.UserId) (out *pb.UserBase, outerr error) {
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

func (s *UserServer) UserQueryAll(ctx context.Context, in *pb.UserAllOption) (*pb.UserAll, error) {
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
