// 这个包主要是增量更新sql

package sqlupdate

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"sort"
)

var (
	NoSqlNeedUpdate = errors.New("no sql need update")
)

// sql更新
type SqlUpdate struct {
	Project string
	Update  []UpdateRecord
}

// 更新记录
type UpdateRecord struct {
	Version string `json:"version"`
	Author  string `json:"author"`
	File    string `json:"file"`
	Date    string `json:"date"`
}

// 解析记录
func (s *SqlUpdate) decode(filename string) error {
	// 读取文件
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	// 解析数据
	err = json.Unmarshal(bs, s)
	if err != nil {
		return err
	}
	return nil
}

// 比较应该执行的sql
func (s *SqlUpdate) compareResult(lastVersion, currentVersion string) []UpdateRecord {
	res := make([]UpdateRecord, 0)
	for _, u := range s.Update {
		if compare(lastVersion, u.Version) && compare(u.Version, currentVersion)  {
			res = append(res, u)
		}
	}
	return res
}

// 比较获取需要更新的版本
// 判断旧版本是否大于要比较版本,true:小于等于(需要更新) false:大于(不需要更新)
func compare(ov, cv string) bool {
	return ov < cv
}

// 返回需要执行的sql
func getSql(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// 根据sql版本升级记录,新旧版本号,返回需要升级的sql
func (s *SqlUpdate) GetSqls(filename, oldVersion, newVersion string) (string, error) {
	// 解析更新文件
	err := s.decode(filename)
	if err != nil {
		return "", err
	}
	// 解析需要更新的sql
	res := s.compareResult(oldVersion, newVersion)
	if len(res) <= 0 {
		return "", NoSqlNeedUpdate
	}
	// 排序sql
	res = updateSqlSort(res)
	// 获取需要更新的sql
	result := make([]byte, 0)
	for _, v := range res {
		bs, err := getSql(v.File)
		if err != nil {
			return "", err
		}
		result = append(result, bs...)
	}

	return string(result), nil
}

// 排序
func updateSqlSort(ur []UpdateRecord) []UpdateRecord {
	sort.Slice(ur, func(i, j int) bool {
		return ur[i].Version < ur[j].Version
	})
	return ur
}
