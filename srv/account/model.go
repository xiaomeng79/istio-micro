package account

import (
	"context"
	"database/sql"
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/xiaomeng79/go-log"
	"github.com/xiaomeng79/istio-micro/cinit"
	"github.com/xiaomeng79/istio-micro/internal/utils"
	pb "github.com/xiaomeng79/istio-micro/srv/account/proto"
)

type Account struct {
	base *pb.AccountBase
	Page utils.Page
}

// 验证参数
func (m *Account) validate() error {
	_, err := govalidator.ValidateStruct(m.base)
	return err
}

// 验证id
func (m *Account) validateId() error {
	if m.base.Id <= 0 {
		return errors.New("id必须大于0")
	}
	return nil
}

// 添加之前
func (m *Account) beforeAdd(ctx context.Context) error {
	// 验证参数
	err := utils.V(m.validate)
	if err != nil {
		return err
	}
	return nil
}

// 修改之前
func (m *Account) beforeUpdate(ctx context.Context) error {
	err := utils.V(m.validateId)
	if err != nil {
		return err
	}
	return nil
}

// 删除之前
func (m *Account) beforeDelete(ctx context.Context) error {
	err := utils.V(m.validateId)
	if err != nil {
		return err
	}
	return nil
}

// 添加 ,返回id
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
	//m.base.Id, err = utils.ID(r, err)
	if err != nil {
		log.Error(err.Error(), ctx)
		return err
	}
	return nil
}

// 修改
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

//
////删除
//func (m *Account) Delete(ctx context.Context) error {
//	err := m.beforeDelete(ctx)
//	if err != nil {
//		log.Info(err.Error(), ctx)
//		return err
//	}
//	r, err := cinit.Mysql.Exec(`UPDATE Account  SET is_usable=0 WHERE id=?`, m.Id)
//	err = utils.R(r, err)
//	if err != nil {
//		log.Error(err.Error(), ctx)
//		return err
//	}
//	m.afterDelete(ctx)
//	return nil
//}
//


// 查询一个
func (m *Account) QueryOne(ctx context.Context) error {
	err := utils.V(m.validateId)
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

////查询全部
//func (m *Account) QueryAll(ctx context.Context) ([]*Account, utils.Page, error) {
//	var err error
//	all := make([]*Account, 0, m.Page.PageSize)
//	countQuery := `SELECT count(id) FROM Account WHERE is_usable=1 `
//	query := `SELECT id,Account_name,password,iphone,sex FROM Account WHERE is_usable=1 `
//
//	/***************************使用IN**********************************/
//	//_sql := `SELECT id,Account_name,password,iphone,sex FROM Account WHERE id IN(?)`
//	//_args := make([]interface{}, 0)
//	//_sql, _args, err = sqlx.In(_sql, args...)
//	//if err != nil {
//	//	log.Error(err.Error(), ctx)
//	//}
//	//_sql = cinit.Mysql.Rebind(_sql)
//
//	/**********************************************************************/
//	where := ""
//	args := make([]interface{}, 0)
//	if m.Sex > 0 {
//		where += `AND sex=? `
//		args = append(args, m.Sex)
//	}
//	if len(m.AccountName) > 0 {
//		where += `AND Account_name=? `
//		args = append(args, m.AccountName)
//	}
//
//	var total int64
//	//统计总页数
//	err = cinit.Mysql.Get(&total, countQuery+where, args...)
//	if err != nil {
//		log.Error(err.Error(), ctx)
//		return nil, m.Page, err
//	}
//	log.Debugf("总页数:%v", total, ctx)
//
//	m.Page.InitPage(total)
//	//加页数
//	limit := `LIMIT ?,?`
//	args = append(args, (m.Page.PageIndex-1)*m.Page.PageSize)
//	args = append(args, m.Page.PageSize)
//	err = cinit.Mysql.Select(&all, query+where+limit, args...)
//	if err != nil {
//		log.Error(err.Error(), ctx)
//		return nil, m.Page, err
//	}
//	return all, m.Page, err
//}



