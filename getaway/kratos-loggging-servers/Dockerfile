 FROM golang:latest

 MAINTAINER docker747 "docker747@163.com"

 WORKDIR $GOPATH/src/github.com/KXX747/wolf/getaway/kratos-loggging-servers/
 ADD . $GOPATH/src/github.com/KXX747/wolf/getaway/kratos-loggging-servers/

 RUN cd $GOPATH/src/github.com/KXX747/wolf/getaway/kratos-loggging-servers/
 RUN chmod +X $GOPATH/src/github.com/KXX747/wolf/getaway/kratos-loggging-servers/start.sh

 #log和trace日志地址
 RUN mkdir -p /tmp/trace/unix_sock
 RUN mkdir -p /tmp/log/unix_sock
 VOLUME /tmp/trace/unix_sock
 VOLUME /tmp/log/unix_sock

 #RUN pwd
 RUN cd cmd && go build

 EXPOSE 38701 38702

 CMD ["./start.sh"]

 #构建
  #docker build -t kratos-loggging-servers .

 #运行
 #docker run --name=pb_loggging_servers -p 38701:38701 -p 38702:38702  -v /tmp/trace/unix_sock:/tmp/trace/unix_sock -v /tmp/log/unix_sock:/tmp/log/unix_sock -d kratos-loggging-servers

