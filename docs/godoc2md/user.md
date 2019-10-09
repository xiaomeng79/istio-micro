# user
`import "/home/meng/workspace/go/istio-micro/srv/user/"`

* [Overview](#pkg-overview)
* [Imported Packages](#pkg-imports)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>

## <a name="pkg-imports">Imported Packages</a>

- github.com/Shopify/sarama
- github.com/asaskevich/govalidator
- github.com/grpc-ecosystem/go-grpc-middleware
- github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing
- github.com/jinzhu/copier
- github.com/xiaomeng79/go-log
- github.com/xiaomeng79/istio-micro/cinit
- github.com/xiaomeng79/istio-micro/internal/gateway
- github.com/xiaomeng79/istio-micro/internal/utils
- github.com/xiaomeng79/istio-micro/internal/wrapper
- github.com/xiaomeng79/istio-micro/srv/user/proto
- google.golang.org/grpc
- google.golang.org/grpc/codes
- google.golang.org/grpc/reflection
- google.golang.org/grpc/status

## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func Run()](#Run)
* [func CacheDel(ctx context.Context, id int64)](#CacheDel)
* [func CacheGet(ctx context.Context, id int64) (map[string]string, error)](#CacheGet)
* [func CacheSet(ctx context.Context, id int64)](#CacheSet)
* [type User](#User)
  * [func (m \*User) Add(ctx context.Context) error](#User.Add)
  * [func (m \*User) Delete(ctx context.Context) error](#User.Delete)
  * [func (m \*User) QueryAll(ctx context.Context) ([]\*User, utils.Page, error)](#User.QueryAll)
  * [func (m \*User) QueryOne(ctx context.Context) error](#User.QueryOne)
  * [func (m \*User) Update(ctx context.Context) error](#User.Update)
* [type Server](#Server)
  * [func (s \*Server) UserAdd(ctx context.Context, in \*pb.UserBase) (out \*pb.UserBase, outerr error)](#Server.UserAdd)
  * [func (s \*Server) UserDelete(ctx context.Context, in \*pb.UserID) (out \*pb.UserID, outerr error)](#Server.UserDelete)
  * [func (s \*Server) UserQueryAll(ctx context.Context, in \*pb.UserAllOption) (\*pb.UserAll, error)](#Server.UserQueryAll)
  * [func (s \*Server) UserQueryOne(ctx context.Context, in \*pb.UserID) (out \*pb.UserBase, outerr error)](#Server.UserQueryOne)
  * [func (s \*Server) UserUpdate(ctx context.Context, in \*pb.UserBase) (out \*pb.UserBase, outerr error)](#Server.UserUpdate)

#### <a name="pkg-files">Package files</a>
[cache.go](./cache.go) [common_cache.go](./common_cache.go) [model.go](./model.go) [msg_queue.go](./msg_queue.go) [run.go](./run.go) [user.go](./user.go) 

## <a name="pkg-constants">Constants</a>
``` go
const (
    KeyMaxExpire     = 500// 秒
    AgainGetStopTime = 100 * time.Millisecond
)
```
``` go
const (
    SexMan   = 1
    SexWoman = 2
    SexOther = 3
)
```
性别

``` go
const (
    SN = "srv-user"// 定义services名称
)
```
``` go
const (
    CacheIDPrefix = "ucid"
)
```

## <a name="Run">func</a> [Run](./run.go#L21)
``` go
func Run()
```

## <a name="CacheDel">func</a> [CacheDel](./cache.go#L61)
``` go
func CacheDel(ctx context.Context, id int64)
```

## <a name="CacheGet">func</a> [CacheGet](./cache.go#L16)
``` go
func CacheGet(ctx context.Context, id int64) (map[string]string, error)
```

## <a name="CacheSet">func</a> [CacheSet](./cache.go#L42)
``` go
func CacheSet(ctx context.Context, id int64)
```

## <a name="User">type</a> [User](./model.go#L13-L21)
``` go
type User struct {
    ID       int64  `json:"id" db:"id" valid:"int~用户id类型为int"`
    UserName string `json:"user_name" db:"user_name" valid:"required~用户名称必须存在"`
    Password string `json:"password" db:"password" valid:"required~密码必须存在"`
    Iphone   string `json:"iphone" db:"iphone" valid:"required~手机号码必须存在"`
    Sex      int32  `json:"sex" db:"sex" valid:"required~性别必须存在"`
    IsUsable int32  `json:"-" db:"is_usable"`
    Page     utils.Page
}
```

### <a name="User.Add">func</a> (\*User) [Add](./model.go#L127)
``` go
func (m *User) Add(ctx context.Context) error
```
添加

### <a name="User.Delete">func</a> (\*User) [Delete](./model.go#L163)
``` go
func (m *User) Delete(ctx context.Context) error
```
删除

### <a name="User.QueryAll">func</a> (\*User) [QueryAll](./model.go#L201)
``` go
func (m *User) QueryAll(ctx context.Context) ([]*User, utils.Page, error)
```
查询全部

### <a name="User.QueryOne">func</a> (\*User) [QueryOne](./model.go#L180)
``` go
func (m *User) QueryOne(ctx context.Context) error
```
查询一个

### <a name="User.Update">func</a> (\*User) [Update](./model.go#L145)
``` go
func (m *User) Update(ctx context.Context) error
```
修改

## <a name="Server">type</a> [Server](./user.go#L12)
``` go
type Server struct{}
```

### <a name="Server.UserAdd">func</a> (\*Server) [UserAdd](./user.go#L14)
``` go
func (s *Server) UserAdd(ctx context.Context, in *pb.UserBase) (out *pb.UserBase, outerr error)
```

### <a name="Server.UserDelete">func</a> (\*Server) [UserDelete](./user.go#L60)
``` go
func (s *Server) UserDelete(ctx context.Context, in *pb.UserID) (out *pb.UserID, outerr error)
```

### <a name="Server.UserQueryAll">func</a> (\*Server) [UserQueryAll](./user.go#L106)
``` go
func (s *Server) UserQueryAll(ctx context.Context, in *pb.UserAllOption) (*pb.UserAll, error)
```

### <a name="Server.UserQueryOne">func</a> (\*Server) [UserQueryOne](./user.go#L83)
``` go
func (s *Server) UserQueryOne(ctx context.Context, in *pb.UserID) (out *pb.UserBase, outerr error)
```

### <a name="Server.UserUpdate">func</a> (\*Server) [UserUpdate](./user.go#L37)
``` go
func (s *Server) UserUpdate(ctx context.Context, in *pb.UserBase) (out *pb.UserBase, outerr error)
```

- - -
Generated by [godoc2ghmd](https://github.com/GandalfUK/godoc2ghmd)