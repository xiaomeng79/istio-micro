// 这个包主要是增量更新sql

package sqlupdate

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"sort"
)

var (
	ErrNoSQLNeedUpdate = errors.New("no sql need update")
)

//  sql更新
type SQLUpdate struct {
	Project string
	Update  []UpdateRecord
}

//  更新记录
type UpdateRecord struct {
	Version string `json:"version"`
	Author  string `json:"author"`
	File    string `json:"file"`
	Date    string `json:"date"`
}

//  解析记录
func (s *SQLUpdate) decode(filename string) error {
	//  读取文件
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	//  解析数据
	err = json.Unmarshal(bs, s)
	if err != nil {
		return err
	}
	return nil
}

//  比较应该执行的sql
func (s *SQLUpdate) compareResult(lastVersion, currentVersion string) []UpdateRecord {
	res := make([]UpdateRecord, 0)
	for _, u := range s.Update {
		if compare(lastVersion, u.Version) && compare(u.Version, currentVersion) {
			res = append(res, u)
		}
	}
	return res
}

//  比较获取需要更新的版本//  判断旧版本是否大于要比较版本,true:小于等于(需要更新) false:大于(不需要更新)
func compare(ov, cv string) bool {
	return versionOrdinal(ov) < versionOrdinal(cv)
}

// 版本号顺序
func versionOrdinal(version string) string {
	// ISO/IEC 14651:2011
	const maxByte = 1<<8 - 1
	vo := make([]byte, 0, len(version)+8)
	j := -1
	for i := 0; i < len(version); i++ {
		b := version[i]
		if '0' > b || b > '9' {
			vo = append(vo, b)
			j = -1
			continue
		}
		if j == -1 {
			vo = append(vo, 0x00)
			j = len(vo) - 1
		}
		if vo[j] == 1 && vo[j+1] == '0' {
			vo[j+1] = b
			continue
		}
		if vo[j]+1 > maxByte {
			panic("VersionOrdinal: invalid version")
		}
		vo = append(vo, b)
		vo[j]++
	}
	return string(vo)
}

//  返回需要执行的sql
func getSQL(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

//  根据sql版本升级记录,新旧版本号,返回需要升级的sql
func (s *SQLUpdate) GetSqls(filename, oldVersion, newVersion string) (string, error) {
	//  解析更新文件
	err := s.decode(filename)
	if err != nil {
		return "", err
	}
	//  解析需要更新的sql
	res := s.compareResult(oldVersion, newVersion)
	if len(res) == 0 {
		return "", ErrNoSQLNeedUpdate
	}
	//  排序sql
	res = updateSQLSort(res)
	//  获取需要更新的sql
	result := make([]byte, 0)
	for _, v := range res {
		bs, err := getSQL(v.File)
		if err != nil {
			return "", err
		}
		result = append(result, bs...)
	}

	return string(result), nil
}

//  排序
func updateSQLSort(ur []UpdateRecord) []UpdateRecord {
	sort.Slice(ur, func(i, j int) bool {
		return ur[i].Version < ur[j].Version
	})
	return ur
}
