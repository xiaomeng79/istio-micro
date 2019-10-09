package utils

import (
	"fmt"
	"math/rand"
	"time"
)

type vFn func() error

// 验证
func V(fns ...vFn) error {
	for _, fn := range fns {
		err := fn()
		if err != nil {
			return err
		}
	}
	return nil
}

func EightBitRand() string {
	_rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%08d", _rand.Int31n(100000000))
}
