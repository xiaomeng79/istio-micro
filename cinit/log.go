package cinit

import (
	"github.com/xiaomeng79/go-log"
	"github.com/xiaomeng79/go-log/conf"
	"github.com/xiaomeng79/go-log/plugins/zaplog"
)

// 初始化日志,可以再这里初始化不同日志引擎的日志 、、 zap logrous// 初始化zap// 设置日志引擎为刚初始化的
func logInit() {
	log.SetLogger(zaplog.New(
		conf.WithProjectName(Config.Service.Name),
		conf.WithLogPath(Config.Log.Path),
		conf.WithLogName(Config.Service.Name),
		conf.WithMaxAge(Config.Log.MaxAge),
		conf.WithMaxSize(Config.Log.MaxSize),
		conf.WithIsStdOut(Config.Log.IsStdOut),
		conf.WithLogLevel(Config.Log.LogLevel),
	))
}
