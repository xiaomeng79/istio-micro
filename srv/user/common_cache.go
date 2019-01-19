package user

import (
	"context"
	"github.com/xiaomeng79/go-log"
	"github.com/xiaomeng79/istio-micro/cinit"
	"math/rand"
	"strconv"
	"time"
)

const (
	KeyMaxExpire     = 500 //秒
	AgainGetStopTime = 100 * time.Millisecond
)

func getIdKey(prefix string, ids ...int64) string {
	s := prefix
	for _, id := range ids {
		s += "_" + strconv.FormatInt(id, 10)
	}
	return s
}

func setKeyExpire(ctx context.Context, ks ...string) {
	rand.Seed(time.Now().UnixNano())
	t := time.Second * time.Duration(rand.Intn(KeyMaxExpire))
	for _, k := range ks {
		err := cinit.RedisCli.Expire(k, t).Err()
		if err != nil {
			log.Error(err.Error(), ctx)
		}
	}
	return
}
