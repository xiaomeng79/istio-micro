package user

import (
	"context"
	"github.com/xiaomeng79/go-log"
	"github.com/xiaomeng79/istio-micro/cinit"
	"github.com/xiaomeng79/istio-micro/internal/utils"
	"sync"
	"time"
)

const (
	UserCacheIdPrefix = "ucid"
)

func UserCacheGet(ctx context.Context, id int64) (map[string]string, error) {
	k := getIdKey(UserCacheIdPrefix, id)
	var mu sync.RWMutex
	var b bool
	//获取全部
	r, err := cinit.RedisCli.HGetAll(k).Result()
	if err != nil {
		log.Info(err.Error(), ctx)
	}
	if len(r) == 0 {
		if !b {
			mu.RLock()
			b = true
			UserCacheSet(ctx, id)
			mu.RUnlock()
		} else {
			time.Sleep(AgainGetStopTime)
		}
		if _, ok := ctx.Value(k).(int); !ok {
			ctx = context.WithValue(ctx, k, 1)
			r, err = UserCacheGet(ctx, id)
		}
	}
	return r, err
}

func UserCacheSet(ctx context.Context, id int64) {
	m := new(User)
	m.Id = id
	err := m.QueryOne(ctx)
	if err != nil {
		log.Info(err.Error(), ctx)
		return
	}
	_m := utils.Struct2Map(*m)
	k := getIdKey(UserCacheIdPrefix, id)
	err = cinit.RedisCli.HMSet(k, _m).Err()
	if err != nil {
		log.Error(err.Error(), ctx)
		return
	}
	setKeyExpire(ctx, k)
	return
}

func UserCacheDel(ctx context.Context, id int64) {
	k := getIdKey(UserCacheIdPrefix, id)
	err := cinit.RedisCli.Del(k).Err()
	if err != nil {
		log.Info(err.Error(), ctx)
	}
	return
}
