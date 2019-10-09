package account

import (
	"github.com/xiaomeng79/istio-micro/cinit"
	"github.com/xiaomeng79/istio-micro/internal/sqlupdate"
	"github.com/xiaomeng79/istio-micro/version"

	"github.com/xiaomeng79/go-log"
)

//  获取旧的版本号
func getOldVersion() (string, error) {
	var oldversion string
	err := cinit.Pg.Get(&oldversion, `select version from sys_info where id =1 limit 1`)
	if err != nil {
		log.Errorf("获取旧版本号失败:%+v", err)
		return "", err
	}
	return oldversion, nil
}

//  更新新的版本号
func updateVersion() error {
	_, err := cinit.Pg.Exec(`update sys_info set version=$1 where id=1`, version.Version)
	if err != nil {
		log.Errorf("更新版本号失败:%+v", err)
		return err
	}
	return nil
}

//  获取需要执行的sql
func getSQL() (string, error) {
	oldVersion, err := getOldVersion()
	if err != nil {
		return "", err
	}
	s := new(sqlupdate.SQLUpdate)
	sqls, err := s.GetSqls("./sqlupdate/record.json", oldVersion, version.Version)
	if err != nil {
		log.Errorf("获取执行sql失败:%+v", err)
		return "", err
	}
	return sqls, nil
}

//  执行sql
func execUpdateSQL() error {
	sqls, err := getSQL()
	if err != nil {
		if err == sqlupdate.ErrNoSQLNeedUpdate {
			return nil
		}
		return err
	}
	tx, err := cinit.Pg.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(sqls)
	if err != nil {
		log.Errorf("执行sql失败:%+v", err)
		_ = tx.Rollback()
		return err
	}
	_ = tx.Commit()
	return nil
}
