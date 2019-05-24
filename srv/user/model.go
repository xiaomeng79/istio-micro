package user

import (
	"context"
	"database/sql"
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/xiaomeng79/go-log"
	"github.com/xiaomeng79/istio-micro/cinit"
	"github.com/xiaomeng79/istio-micro/internal/utils"
)

type User struct {
	Id       int64  `json:"id" db:"id" valid:"int~用户id类型为int"`
	UserName string `json:"user_name" db:"user_name" valid:"required~用户名称必须存在"`
	Password string `json:"password" db:"password" valid:"required~密码必须存在"`
	Iphone   string `json:"iphone" db:"iphone" valid:"required~手机号码必须存在"`
	Sex      int32  `json:"sex" db:"sex" valid:"required~性别必须存在"`
	IsUsable int32  `json:"-" db:"is_usable"`
	Page     utils.Page
}

//性别
const (
	SexMan   = 1
	SexWoman = 2
	SexOther = 3
)

var (
	sexTypes = []int32{
		SexMan,
		SexWoman,
		SexOther,
	}
)

//从缓存获取数据
func (m *User) getFromCache(ctx context.Context) error {
	r, err := UserCacheGet(ctx, m.Id)
	if err != nil {
		return err
	}
	utils.Map2Struct(r, m)
	return nil
}

//验证参数
func (m *User) validate() error {
	_, err := govalidator.ValidateStruct(m)
	return err
}

//验证id
func (m *User) validateId() error {
	if m.Id <= 0 {
		return errors.New("id必须大于0")
	}
	return nil
}

//验证性别类型
func (m *User) validateSexType() error {
	b := false
	for _, v := range sexTypes {
		if m.Sex == v {
			b = true
			break
		}
	}
	if !b {
		return errors.New("性别类型不合法")
	}
	return nil
}

//添加之前
func (m *User) beforeAdd(ctx context.Context) error {
	//验证参数
	err := utils.V(m.validate)
	if err != nil {
		return err
	}
	return nil
}

//修改之前
func (m *User) beforeUpdate(ctx context.Context) error {
	err := utils.V(m.validate, m.validateId)
	if err != nil {
		return err
	}
	return nil
}

//删除之前
func (m *User) beforeDelete(ctx context.Context) error {
	err := utils.V(m.validateId)
	if err != nil {
		return err
	}
	return nil
}

//添加之后,异步操作
func (m *User) afterAdd(ctx context.Context) error {
	go msgNotify(ctx, "添加用户:"+m.UserName)
	return nil
}

//修改之后,异步操作
func (m *User) afterUpdate(ctx context.Context) error {
	//删除缓存
	go UserCacheDel(ctx, m.Id)
	go msgNotify(ctx, "修改用户:"+m.UserName)
	return nil
}

//删除之后,异步操作
func (m *User) afterDelete(ctx context.Context) error {
	//删除缓存
	go UserCacheDel(ctx, m.Id)
	return nil
}

//添加
func (m *User) Add(ctx context.Context) error {
	err := m.beforeAdd(ctx)
	if err != nil {
		log.Info(err.Error(), ctx)
		return err
	}
	r, err := cinit.Mysql.Exec(`INSERT INTO user (user_name,password,iphone,sex) VALUES (?,?,?,?)`,
		m.UserName, m.Password, m.Iphone, m.Sex)
	m.Id, err = utils.ID(r, err)
	if err != nil {
		log.Error(err.Error(), ctx)
		return err
	}
	m.afterAdd(ctx)
	return nil
}

//修改
func (m *User) Update(ctx context.Context) error {
	err := m.beforeUpdate(ctx)
	if err != nil {
		log.Info(err.Error(), ctx)
		return err
	}
	r, err := cinit.Mysql.Exec(`UPDATE user  SET user_name=?,password=?,iphone=?,sex=? WHERE id=?`,
		m.UserName, m.Password, m.Iphone, m.Sex, m.Id)
	err = utils.R(r, err)
	if err != nil {
		log.Error(err.Error(), ctx)
		return err
	}
	m.afterUpdate(ctx)
	return nil
}

//删除
func (m *User) Delete(ctx context.Context) error {
	err := m.beforeDelete(ctx)
	if err != nil {
		log.Info(err.Error(), ctx)
		return err
	}
	r, err := cinit.Mysql.Exec(`UPDATE user  SET is_usable=0 WHERE id=?`, m.Id)
	err = utils.R(r, err)
	if err != nil {
		log.Error(err.Error(), ctx)
		return err
	}
	m.afterDelete(ctx)
	return nil
}

//查询一个
func (m *User) QueryOne(ctx context.Context) error {
	err := utils.V(m.validateId)
	if err != nil {
		log.Info(err.Error(), ctx)
		return err
	}

	err = cinit.Mysql.Get(m, `SELECT id,user_name,password,iphone,sex FROM user WHERE id=? and is_usable=1 LIMIT 1`, m.Id)
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

//查询全部
func (m *User) QueryAll(ctx context.Context) ([]*User, utils.Page, error) {
	var err error
	all := make([]*User, 0, m.Page.PageSize)
	countQuery := `SELECT count(id) FROM user WHERE is_usable=1 `
	query := `SELECT id,user_name,password,iphone,sex FROM user WHERE is_usable=1 `

	/***************************使用IN**********************************/
	//_sql := `SELECT id,user_name,password,iphone,sex FROM user WHERE id IN(?)`
	//_args := make([]interface{}, 0)
	//_sql, _args, err = sqlx.In(_sql, args...)
	//if err != nil {
	//	log.Error(err.Error(), ctx)
	//}
	//_sql = cinit.Mysql.Rebind(_sql)

	/**********************************************************************/
	where := ""
	args := make([]interface{}, 0)
	if m.Sex > 0 {
		where += `AND sex=? `
		args = append(args, m.Sex)
	}
	if len(m.UserName) > 0 {
		where += `AND user_name=? `
		args = append(args, m.UserName)
	}

	var total int64
	//统计总页数
	err = cinit.Mysql.Get(&total, countQuery+where, args...)
	if err != nil {
		log.Error(err.Error(), ctx)
		return nil, m.Page, err
	}
	log.Debugf("总页数:%v", total, ctx)

	m.Page.InitPage(total)
	//加页数
	limit := `LIMIT ?,?`
	args = append(args, (m.Page.PageIndex-1)*m.Page.PageSize)
	args = append(args, m.Page.PageSize)
	err = cinit.Mysql.Select(&all, query+where+limit, args...)
	if err != nil {
		log.Error(err.Error(), ctx)
		return nil, m.Page, err
	}
	return all, m.Page, err
}
