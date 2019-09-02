package kkafka

import (
	"github.com/bilibili/kratos/pkg/database/mongo"
	"time"
	xtime "github.com/bilibili/kratos/pkg/time"
)

const (
	CH_LEN =1024


)

var(
	mongoChan =make(chan MsgData,2024) //
	hbaseChan =make(chan MsgData,2024) //
	esChan =make(chan MsgData,2024) //

	//mongo配置和连接
	mongoConfig *mongo.Config
	mongoAddr ="127.0.0.1:27017"
	mDatabase *mgo.Database
)

func init() {
	mongoConfig =new(mongo.Config)
	mongoConfig.Addr =mongoAddr
	mongoConfig.DSN =mongoAddr
	mongoConfig.ReadDSN = []string{mongoAddr,mongoAddr}
	mongoConfig.IdleTimeout =xtime.Duration(20*time.Millisecond)

	db :=mongo.NewMongoClient(mongoConfig)
	mDatabase=db.Write.DB.DB("test")

	//默认使用log库
	SwichCol(_loggings)
}

//切换库
func SwichCol(name string)(mCollection *mgo.Collection){

	mCollection=mDatabase.C(name)
	return
}


//接收和生成消息
type MsgData struct {
	Key []byte
	Value []byte
}

type MsgDataChan struct {
	Ch chan MsgData
}

type P struct {}
type S struct {}