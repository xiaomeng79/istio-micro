package cinit

import (
	"github.com/jinzhu/configor"
	"github.com/xiaomeng79/istio-micro/pkg/pprof"
	"log"
)

const (
	//上下文
	TRACE_CONTEXT     = "trace_ctx"     //trace
	REQ_PARAM         = "req_param"     //请求参数绑定
	JWT_NAME          = "Authorization" //JWT请求头名称
	JWT_MSG           = "JWT-MSG"       //JWT自定义的消息
	SORT_NUM          = 100             //默认排序号码
	FLOAT_COMPUTE_BIT = 2               //浮点计算位数
	REDIS_NIL         = "redis: nil"
)

//公共配置
var Config = struct {
	Service struct {
		Name      string `default:""`     //服务名称
		Version   string `default:"v1.0"` //服务版本号
		RateTime  int    `default:"1024"` //限制请求
		AppKey    string `default:"admin"`
		AppSecret string `default:"admin"`
	}
	//tracing
	Trace struct { //链路跟踪
		Address       string  `default:"http://jaeger:14268/api/traces?format=jaeger.thrift"` // http://jaeger:14268/api/traces?format=jaeger.thrift
		ZipkinURL     string  `default:""`                                                    //http://zipkin:9411/api/v1/spans
		SamplingRate  float64 `default:"0.1"`                                                 // 采样率 0.01-1范围
		LogTraceSpans bool    `default:"false"`                                               // 日志
	}
	//log config
	Log struct { //日志
		Path         string `default:"tmp"`   //日志保存路径
		IsStdOut     string `default:"yes"`   //是否输出日志到标准输出 yes:输出 no:不输出
		MaxAge       int    `default:"7"`     //日志最大的保存时间，单位天
		RotationTime int    `default:"1"`     //日志分割的时间，单位天
		MaxSize      int    `default:"100"`   //日志分割的尺寸，单位MB
		LogLevel     string `default:"debug"` //日志级别
	}
	//mysql config
	Mysql struct {
		DbName   string `default:"guess"`     //数据库名称
		Addr     string `default:"127.0.0.1"` //地址
		User     string `default:"root"`
		Password string `default:"root"`
		Port     int    `default:"3306"` //required:"true" env:"DB_PROT"
		IdleConn int    `default:"4"`    //空闲连接
		MaxConn  int    `default:"20"`   //最大连接
	}
	//mysql config
	Postgres struct {
		DbName   string `default:"test"`      //数据库名称
		Addr     string `default:"127.0.0.1"` //地址
		User     string `default:"postgres"`
		Password string `default:"postgres"`
		Port     int    `default:"5432"` //required:"true"
		IdleConn int    `default:"4"`    //空闲连接
		MaxConn  int    `default:"20"`   //最大连接
	}
	//redis config
	Redis struct {
		Addr     string `default:"127.0.0.1:6379"` //地址
		Password string `default:""`
		Db       int    `default:"0"`
	}
	//mongo config
	Mongo struct {
		Hosts     string `default:"127.0.0.1:27017"` //数据库地址，可以多个，用逗号分割
		DbName    string `default:"test"`            //数据库名称
		User      string `default:"root"`
		Password  string `default:"root"`
		PoolLimit int    `default:"4096"` //连接池限制
	}
	Kafka struct {
		Addrs string `default:"127.0.0.1:9092"` //数据库地址，可以多个，用逗号分割
	}
	//metrics config
	Metrics struct {
		Enable   string `default:"yes"` //是否启用:yes 启用 no 停用
		Duration int    `default:"5"`   //单位秒
		Url      string `default:"http://influxdb:8086"`
		Database string `default:"test01"`
		UserName string `default:""`
		Password string `default:""`
	}

	//userservice
	SrvUser struct {
		Port              string `default:":5001"`          //定义的端口
		Address           string `default:"127.0.0.1:5001"` //访问地址
		GateWayAddr       string `default:":9999"`          //网关端口
		GateWaySwaggerDir string `default:"/swagger"`       // swagger目录
	}
	//accountservice
	SrvAccount struct {
		Port              string `default:":5003"`          //定义的端口
		Address           string `default:"127.0.0.1:5003"` //访问地址
		GateWayAddr       string `default:":9997"`          //网关端口
		GateWaySwaggerDir string `default:"/swagger"`       // swagger目录
	}
	//api backend
	ApiBackend struct {
		Port    string `default:":8888"`
		Address string `default:"127.0.0.1:8888"`
	}
	//api backend
	ApiFrontend struct {
		Port    string `default:":8889"`
		Address string `default:"127.0.0.1:8889"`
	}
	//gamesocket
	SrvSocket struct {
		Port    string `default:":5002"`          //定义的端口
		Address string `default:"127.0.0.1:5002"` //访问地址
	}
}{}

//初始化配置文件
//配置加载顺序1.是否设置了变量conf，设置了第一个加载，如果文件不存在，加载默认配置文件
//如果设置了环境变量 CONFIGOR_ENV = test等，那么加载config_test.yml的配置文件
//最后加载环境变量,是否设置环境变量前缀,如果设置了CONFIGOR_ENV_PREFIX=WEB,设置环境变量为WEB_DB_NAME=root,否则为DB_NAME=root
func configInit(sn string) {

	//config := flag.String("conf", "conf/config.yml", "you configuer file")
	//flag.Parse()
	//err := configor.Load(&Config, *config)
	configor.Load(&Config, "config.yml")
	if len(Config.Service.Name) == 0 {
		Config.Service.Name = sn //使用传入的名称
	}
	log.Printf("config: %+v\n", Config)
}

//保存需要关闭的选项
var closeArgs []string

//初始化选项
//log:日志(必须) trace:链路跟踪 mysql:mysql数据库 mongo:MongoDB postgres:postgres数据库
func InitOption(sn string, args ...string) {
	//开启pprof
	go pprof.Run()
	//保存需要关闭的参数
	closeArgs = args
	//1.初始化配置参数
	configInit(sn)
	//2.初始化日志
	logInit()
	//3.其他服务
	for _, o := range args {
		switch o {
		case "trace":
			traceInit()
		case "mysql":
			mysqlInit()
		case "mongo":
		case "redis":
			redisInit()
		case "kafka":
			KafkaInit()
		case "metrics":
			metricsInit(sn)
		case "postgres":
			pgInit()
		}

	}
}

//关闭打开的服务
func Close() {
	for _, o := range closeArgs {
		switch o {
		case "trace":
			//关闭链路跟踪
			tracerClose()
		case "mysql":
			//关闭mysql
			//原始
			mysqlClose()
		case "mongo":
		case "redis":
			redisClose()
		case "kafka":
			KafkaClose()
		case "metrics":
		case "postgres":
			pgClose()
		}
	}
}
