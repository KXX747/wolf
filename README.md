
common 公共组件
getaway 对外网关
public  服务


docker stop pb_loggging_servers
docker stop pb_getaway_servers
docker stop pb_user_server
docker stop pb_stream_server

docker start pb_loggging_servers
docker start pb_user_server
docker start pb_stream_server
docker start pb_getaway_servers



export GO111MODULE=on 
go mod download
go mod vendor 


//添加库
go mod edit -replace=golang.org/x/image@v0.0.0-20180708004352-c73c2afc3b81=github.com/golang/image@v0.0.0-20180708004352-c73c2afc3b81

go mod edit -replace=gopkg.in/mgo.v2=https://github.com/globalsign/mgo

go mod edit -replace=gopkg.in/mgo.v2@0.0.0=github.com/globalsign/mgo@0.0.0


go mod edit -require=github.com/pkg/sftp@v1.10.0
go mod edit -require=github.com/globalsign/mgo@0.0.0

go mod edit -require=github.com/samuel/go-zookeeper@latest
go mod edit -require=github.com/gogo/protobuf/gogoproto@latest

go get -u github.com/gogo/protobuf/protoc-gen-gofast

go get github.com/alecthomas/gometalinter



//使用本地的
replace github.com/gohouse/goroom => /path/to/go/src/github.com/gohouse/goroom


protoc-gen-swagger
export GO111MODULE=on 
go get -u github.com/bazelbuild/bazel-gazelle

yacc安装 let
解决goyacc错误 cmd/goyacc'

go get github.com/golang/tools/cmd/goyacc
https://github.com/golang/tools/tree/master/cmd/goyacc
go build
go intall
