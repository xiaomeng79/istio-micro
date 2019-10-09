package account

import (
	"context"
	"database/sql"
	"errors"

	"github.com/xiaomeng79/istio-micro/cinit"
	"github.com/xiaomeng79/istio-micro/internal/utils"
	pb "github.com/xiaomeng79/istio-micro/srv/account/proto"

	"github.com/asaskevich/govalidator"
	"github.com/xiaomeng79/go-log"
)

type Account struct {
	base *pb.AccountBase
	Page utils.Page
}

//  验证参数
func (m *Account) validate() error {
	_, err := govalidator.ValidateStruct(m.base)
	return err
}

//  验证id
func (m *Account) validateID() error {
	if m.base.Id <= 0 {
		return errors.New("id必须大于0")
	}
	return nil
}

//  添加之前
func (m *Account) beforeAdd(ctx context.Context) error { // nolint: unparam
	//  验证参数
	err := utils.V(m.validate)
	if err != nil {
		return err
	}
	return nil
}

//  修改之前
// nolint: unparam
func (m *Account) beforeUpdate(ctx context.Context) error {
	err := utils.V(m.validateID)
	if err != nil {
		return err
	}
	return nil
}

//  删除之前
// nolint: unparam,unused
func (m *Account) beforeDelete(ctx context.Context) error {
	err := utils.V(m.validateID)
	if err != nil {
		return err
	}
	return nil
}

//  添加 ,返回id
func (m *Account) Add(ctx context.Context) error {
	err := m.beforeAdd(ctx)
	if err != nil {
		log.Info(err.Error(), ctx)
		return err
	}
	_sql := `INSERT INTO account (user_id,account_level,balance,account_status) VALUES (:user_id,:account_level,:balance,:account_status) RETURNING id`
	stmt, err := cinit.Pg.PrepareNamed(_sql)
	if err != nil {
		log.Error(err.Error(), ctx)
		return err
	}
	arg := map[string]interface{}{
		"user_id":        m.base.UserId,
		"account_level":  m.base.AccountLevel,
		"balance":        m.base.Balance,
		"account_status": m.base.AccountStatus,
	}
	err = stmt.Get(&m.base.Id, arg)
	defer stmt.Close()
	log.Debugf("m:%+v", m.base, ctx)
	// m.base.ID, err = utils.ID(r, err)
	if err != nil {
		log.Error(err.Error(), ctx)
		return err
	}
	return nil
}

//  修改
func (m *Account) Update(ctx context.Context) error {
	err := m.beforeUpdate(ctx)
	if err != nil {
		log.Info(err.Error(), ctx)
		return err
	}
	r, err := cinit.Pg.Exec(`UPDATE account  SET balance=$1 WHERE id=$2`,
		m.base.Balance, m.base.Id)
	err = utils.R(r, err)
	if err != nil {
		log.Error(err.Error(), ctx)
		return err
	}
	return nil
}

//  查询一个
func (m *Account) QueryOne(ctx context.Context) error {
	err := utils.V(m.validateID)
	if err != nil {
		log.Info(err.Error(), ctx)
		return err
	}

	err = cinit.Pg.Get(m.base, `SELECT id,user_id,account_level,balance,account_status FROM account WHERE id=$1  LIMIT 1`, m.base.Id)
	switch {
	case err == sql.ErrNoRows:
		log.Info(err.Error(), ctx)
		return nil
	case err != nil:
		log.Error(err.Error(), ctx)
		return err
	default:
		return nil
	}
}
