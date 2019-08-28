package account

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/xiaomeng79/istio-micro/srv/account/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccountServer struct{}

func (s *AccountServer) AccountAdd(ctx context.Context, in *pb.AccountBase) (out *pb.AccountBase, outerr error) {
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

func (s *AccountServer) AccountUpdate(ctx context.Context, in *pb.AccountUpdateReq) (out *empty.Empty, outerr error) {
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

//
//func (s *AccountServer) AccountDelete(ctx context.Context, in *pb.AccountId) (out *pb.AccountId, outerr error) {
//	m := new(Account)
//	out = new(pb.AccountId)
//	err := copier.Copy(m, in)
//	if err != nil {
//		log.Error(err.Error(), ctx)
//		outerr = status.Error(codes.Internal, err.Error())
//		return
//	}
//	err = m.Delete(ctx)
//	if err != nil {
//		outerr = status.Error(codes.InvalidArgument, err.Error())
//		return
//	}
//	err = copier.Copy(out, m)
//	if err != nil {
//		log.Error(err.Error(), ctx)
//		outerr = status.Error(codes.Internal, err.Error())
//		return
//	}
//	return
//}

// 查询一个
func (s *AccountServer) AccountQueryOne(ctx context.Context, in *pb.AccountId) (out *pb.AccountBase, outerr error) {
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

//
//func (s *AccountServer) AccountQueryAll(ctx context.Context, in *pb.AccountAllOption) (*pb.AccountAll, error) {
//	m := new(Account)
//	err := copier.Copy(m, in)
//	if err != nil {
//		log.Error(err.Error(), ctx)
//		return &pb.AccountAll{}, status.Error(codes.Internal, err.Error())
//	}
//	ms, page, err := m.QueryAll(ctx)
//	if err != nil {
//		return &pb.AccountAll{}, status.Error(codes.InvalidArgument, err.Error())
//	}
//	var agt []*pb.AccountBase
//	err = copier.Copy(&agt, ms)
//	if err != nil {
//		log.Error(err.Error(), ctx)
//		return &pb.AccountAll{}, status.Error(codes.Internal, err.Error())
//	}
//	_page := new(pb.Page)
//	err = copier.Copy(_page, &page)
//	if err != nil {
//		log.Error(err.Error(), ctx)
//		return &pb.AccountAll{}, status.Error(codes.Internal, err.Error())
//	}
//	return &pb.AccountAll{All: agt, Page: _page}, nil
//}
