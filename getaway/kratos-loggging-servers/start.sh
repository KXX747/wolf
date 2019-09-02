#!/usr/bin/env bash
cd cmd

go build

./cmd -conf ./../configs/

#scp -r kratos-loggging-servers/ root@192.168.57.136:/root/go/src/github.com/KXX747/wolf/getaway/