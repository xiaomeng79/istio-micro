package account

import (
	"context"

	pb "github.com/xiaomeng79/istio-micro/srv/account/proto"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct{}

func (s *Server) AccountAdd(ctx context.Context, in *pb.AccountBase) (out *pb.AccountBase, outerr error) {
	m := new(Account)
	m.base = in
	out = new(pb.AccountBase)
	err := m.Add(ctx)
	if err != nil {
		outerr = status.Error(codes.InvalidArgument, err.Error())
		return
	}
	out = m.base
	return
}

func (s *Server) AccountUpdate(ctx context.Context, in *pb.AccountUpdateReq) (out *empty.Empty, outerr error) {
	m := &Account{base: &pb.AccountBase{
		Id:      in.Id,
		Balance: in.Balance,
	}}
	out = new(empty.Empty)
	err := m.Update(ctx)
	if err != nil {
		outerr = status.Error(codes.InvalidArgument, err.Error())
		return
	}
	return
}

//  查询一个
func (s *Server) AccountQueryOne(ctx context.Context, in *pb.AccountID) (out *pb.AccountBase, outerr error) {
	m := &Account{
		base: &pb.AccountBase{
			Id: in.Id,
		}}
	out = new(pb.AccountBase)
	err := m.QueryOne(ctx)
	if err != nil {
		outerr = status.Error(codes.InvalidArgument, err.Error())
		return
	}
	out = m.base
	return
}
