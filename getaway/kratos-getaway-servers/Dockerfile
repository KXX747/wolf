 FROM golang:latest

 MAINTAINER docker747 "docker747@163.com"

 WORKDIR $GOPATH/src/github.com/KXX747/wolf/getaway/kratos-getaway-servers
 ADD . $GOPATH/src/github.com/KXX747/wolf/getaway/kratos-getaway-servers

 RUN cd $GOPATH/src/github.com/KXX747/wolf/getaway/kratos-getaway-servers
 RUN chmod +X $GOPATH/src/github.com/KXX747/wolf/getaway/kratos-getaway-servers/start.sh

 #log和trace日志地址
 RUN mkdir -p /tmp/trace/unix_sock
 RUN mkdir -p /tmp/log/unix_sock
 VOLUME /tmp/trace/unix_sock
 VOLUME /tmp/log/unix_sock

 #RUN pwd
 RUN cd cmd && go build

 EXPOSE 38705

 #RUN ls &&pwd
 CMD ["./start.sh"]

 #构建
  #docker build -t kratos-getaway-servers .

 #运行
 #docker run --name=pb_getaway_servers -p 38705:38705 -v /tmp/trace/unix_sock:/tmp/trace/unix_sock -v /tmp/log/unix_sock:/tmp/log/unix_sock -d kratos-getaway-servers

