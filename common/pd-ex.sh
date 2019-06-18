#!/usr/bin/env bash

#export $KRATOS_HOME = "/Users/a747/go/src/github.com/bilibili/kratos"
#export $KRATOS_DEMO = "/Users/a747/go/src/github.com/KXX747/wolf/public/user-acount-server"

# 生成：api.pb.go
#protoc -I$GOPATH/src:$KRATOS_HOME/tool/protobuf/pkg/extensions:$KRATOS_DEMO/api --gogofast_out=plugins=grpc:$KRATOS_DEMO/api $KRATOS_DEMO/api/api.proto

#protoc -I/Users/a747/go/src:/Users/a747/go/src/github.com/bilibili/kratos/tool/protobuf/pkg/extensions:/Users/a747/go/src/github.com/KXX747/wolf/public/user-acount-server/api --gogofast_out=plugins=grpc:/Users/a747/go/src/github.com/KXX747/wolf/public/user-acount-server/api /Users/a747/go/src/github.com/KXX747/wolf/public/user-acount-server/api/api.proto



#生成：api.bm.go
#protoc -I$GOPATH/src:$KRATOS_HOME/tool/protobuf/pkg/extensions:$KRATOS_DEMO/api --bm_out=$KRATOS_DEMO/api $KRATOS_DEMO/api/api.proto

# 生成：api.swagger.json
#protoc -I$GOPATH/src:$KRATOS_HOME/tool/protobuf/pkg/extensions:$KRATOS_DEMO/api --bswagger_out=$KRATOS_DEMO/api$KRATOS_DEMO/api/api.proto



#go get -u github.com/gogo/protobuf/protoc-gen-gogo


#export GO111MODULE=on


export KRATOS_HOME="/Users/a747/go/src/github.com/bilibili/kratos"
export KRATOS_DEMO="/Users/a747/go/src/github.com/KXX747/wolf/public/snowflake-uuid-server/api"

echo $KRATOS_HOME

echo $KRATOS_DEMO

	#生成：api.pb.go
protoc -I$GOPATH/src:$KRATOS_HOME/tool/protobuf/pkg/extensions:$KRATOS_DEMO --gogofast_out=plugins=grpc:. api.proto

	# 生成：api.bm.go
protoc -I$GOPATH/src:$KRATOS_HOME/tool/protobuf/pkg/extensions:$KRATOS_DEMO --bm_out=.  api.proto

	# 生成：api.swagger.json
protoc -I$GOPATH/src:$KRATOS_HOME/tool/protobuf/pkg/extensions:$KRATOS_DEMO --bswagger_out=. api.proto






