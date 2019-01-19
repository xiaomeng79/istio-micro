package utils

import (
	"errors"
	"github.com/xiaomeng79/go-log"
	"github.com/xiaomeng79/go-log/conf"
	"github.com/xiaomeng79/go-log/plugins/zaplog"
	"testing"
)

//初始化日志,可以再这里初始化不同日志引擎的日志 、、 zap logrous
//初始化zap
//设置日志引擎为刚初始化的
func logInit() {
	log.SetLogger(zaplog.New(
		conf.WithIsStdOut("no"),
	))
}

func BenchmarkLog(b *testing.B) {
	b.StopTimer()
	logInit()
	err := errors.New("my error is test")
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		log.Info("test" + err.Error())
	}
}

func BenchmarkLogF(b *testing.B) {
	b.StopTimer()
	logInit()
	err := errors.New("my error is test2")
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		log.Infof("test:%s", err)
	}
}
