package user

import (
	"context"

	"github.com/xiaomeng79/istio-micro/cinit"
	"github.com/xiaomeng79/istio-micro/internal/utils"

	"github.com/xiaomeng79/go-log"
)

const (
	CacheIDPrefix = "ucid"
)

func CacheGet(ctx context.Context, id int64) (map[string]string, error) {
	k := getIDKey(CacheIDPrefix, id)
	// 获取全部
	r, err := cinit.RedisCli.HGetAll(k).Result()
	if err != nil {
		log.Info(err.Error(), ctx)
	}
	if len(r) == 0 {
		CacheSet(ctx, id)
	}
	return r, err
}

func CacheSet(ctx context.Context, id int64) {
	m := new(User)
	m.ID = id
	err := m.QueryOne(ctx)
	if err != nil {
		log.Info(err.Error(), ctx)
		return
	}
	_m := utils.Struct2Map(*m)
	k := getIDKey(CacheIDPrefix, id)
	err = cinit.RedisCli.HMSet(k, _m).Err()
	if err != nil {
		log.Error(err.Error(), ctx)
		return
	}
	setKeyExpire(ctx, k)
}

func CacheDel(ctx context.Context, id int64) {
	k := getIDKey(CacheIDPrefix, id)
	err := cinit.RedisCli.Del(k).Err()
	if err != nil {
		log.Info(err.Error(), ctx)
	}
}
