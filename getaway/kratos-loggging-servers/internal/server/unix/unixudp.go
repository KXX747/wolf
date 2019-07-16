package unix

import (
	"net"
	"bufio"
	"github.com/KXX747/wolf/getaway/kratos-loggging-servers/internal/service"
	"github.com/bilibili/kratos/pkg/log"
	"os"
)

var (


)

//init
func NewUnix(svc *service.Service){

	//初始化协议
	go initUnix(svc)

}

func initUnix(svc *service.Service)  {

	//svc.AppConfig.Log.Agent.Addr
	err:=os.Remove(svc.AppConfig.Log.Agent.Addr)
	if err!=nil {
		log.Error("=os.Remove err=%s",err)
		//return
	}
	agent := svc.AppConfig.Log.Agent
	log.Info("agent addr=%s  peoto=%s timeout=%s",agent.Addr,agent.Network,agent.Timeout)
	unixAddr, _ := net.ResolveUnixAddr(agent.Network,agent.Addr)
	unixListener, _ := net.ListenUnix(agent.Network, unixAddr)
	for  {
		unixConn, err := unixListener.AcceptUnix()
		if err != nil {
			panic(err)
		}
		log.Info("A client connected : err= " + unixConn.RemoteAddr().String())
		go unixPipe(unixConn)
	}
}


func unixPipe(conn *net.UnixConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		log.Info("disconnected :" + ipStr)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		log.Info("message =",message)
		msg :=message + "                      回复了。。\n"
		b := []byte(msg)
		conn.Write(b)
	}
}

