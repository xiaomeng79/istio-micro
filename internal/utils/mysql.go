package utils

import (
	"database/sql"
	"errors"
)

// 影响结果判断
func R(result sql.Result, err error) error {
	if err != nil {
		return err
	}
	res, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if res <= 0 {
		return errors.New("更新失败")
	}
	return nil
}

// ID
func ID(result sql.Result, err error) (int64, error) {
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
