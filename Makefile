#定义变量
GOPROXY=https://goproxy.io
GO111MODULE=on
#GOTHIRDPKG=${HOME}/gopkg/third
VERSION=$(shell git describe --abbrev=0 --tags)
COMMIT=$(shell git rev-parse --short HEAD)

#project:game prize pusher socket
#type:api srv web

.PHONY : fmt
fmt :
	@echo "格式化代码"
	@gofmt -l -w ./

.PHONY : vendor
vendor :
	@echo "创建vendor"
	@go mod vendor
	@echo "结束vendor"

.PHONY : test
test : vendor check
	@echo "代码测试[覆盖率]"
	@go test -mod=vendor -race -cover  -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY : onlinetest
onlinetest :
	@echo "代码测试[覆盖率]"
	@go test  -race -cover  -coverprofile=coverage.txt -covermode=atomic ./...

#代码检查
.PHONY : check
check :
	@echo "代码检查"
	@echo "代码静态检查开始"
	@go vet  ./...
	@echo "代码静态检查结束"
	@chmod +x ./scripts/check.sh && ./scripts/check.sh
#	@chmod +x ./scripts/check.sh && ./scripts/check.sh | tee check.log


.PHONY : build
build : proto builddata dockerfile vendor
	@echo "部分编译开始:"$(project)_$(type)
	@chmod +x ./scripts/build.sh && ./scripts/build.sh build $(type) $(project)
	@echo "部分编译结束"



.PHONY : allbuild
allbuild : proto builddata alldockerfile vendor

	@echo "全部编译开始"
	@chmod +x ./scripts/build.sh && ./scripts/build.sh allbuild
	@echo "全部编译结束"



#生成pb文件

.PHONY : proto
proto :

	@echo "生成proto开始"
	@chmod +x ./scripts/proto.sh && ./scripts/proto.sh
	@echo "生成proto结束"

#生成dockerfile
.PHONY : dockerfile
dockerfile :

	@echo "部分生成dockerfile开始"
	@chmod +x ./scripts/dockerfile.sh && ./scripts/dockerfile.sh df $(type) $(project)
	@echo "部分生成Dockerfile结束"


.PHONY : alldockerfile
alldockerfile :

	@echo "全部生成dockerfile开始"
	@chmod +x ./scripts/dockerfile.sh && ./scripts/dockerfile.sh alldf
	@echo "全部生成Dockerfile结束"


#compose命令 bin:up dup stop restart kill rm ps
.PHONY : compose
compose :

	@chmod +x ./scripts/docker-compose.sh && ./scripts/docker-compose.sh $(bin)

.PHONY : builddata
builddata :
		@echo "打包静态数据"
		@chmod +x ./scripts/builddata.sh && ./scripts/builddata.sh

#编辑k8s配置
.PHONY : k8sconfig
k8sconfig :

	@echo "配置k8s"
	@chmod +x ./scripts/k8sconf.sh && ./scripts/k8sconf.sh

#pprof性能分析
.PHONY : pprofon
pprofon :
	@chmod +x ./scripts/pprof.sh && ./scripts/pprof.sh pprofon $(type) $(project)

#pprof性能分析
.PHONY : pprofoff
pprofoff :
	@chmod +x ./scripts/pprof.sh && ./scripts/pprof.sh pprofoff $(type) $(project)

#提交代码
.PHONY : push
push : fmt check test
	git add -A
	git commit -m $(msg)
	git push origin master

#清理
.PHONY : clean
clean :
	@git clean -dxf -e .idea

#发行版本
release :
	@chmod +x ./scripts/release.sh && ./scripts/release.sh

#查看下一个版本号
next-version :
	@chmod +x ./scripts/version.sh && ./scripts/version.sh

#清理没用的docker镜像
docker-clean:
	docker images
	docker image prune --force

docker-kill:
	docker kill `docker ps -q` || true

docker-remove:
	docker rm --force `docker ps -a -q` || true
	docker rmi --force `docker images -q` || true
