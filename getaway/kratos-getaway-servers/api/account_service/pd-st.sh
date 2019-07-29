#!/usr/bin/env bash


echo "pd.proto生成pd.go"
export KRATOS_HOME="/Users/a747/go/src/github.com/bilibili/kratos"
export KRATOS_DEMO="/Users/a747/go/src/kratos-user-account-server/internal/model/user"

echo $KRATOS_HOME

echo $KRATOS_DEMO

	#生成：api.pb.go
protoc -I$GOPATH/src:$KRATOS_HOME/tool/protobuf/pkg/extensions:$KRATOS_DEMO --gogofast_out=plugins=grpc:. user.proto

protoc -I$GOPATH/src:$KRATOS_HOME/tool/protobuf/pkg/extensions:$KRATOS_DEMO --bm_out=.  user.proto

	# 生成：api.swagger.json
protoc -I$GOPATH/src:$KRATOS_HOME/tool/protobuf/pkg/extensions:$KRATOS_DEMO --bswagger_out=. user.proto

