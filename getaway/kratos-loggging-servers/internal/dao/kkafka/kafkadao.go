package kkafka

import (
	"fmt"
	"github.com/bilibili/kratos/pkg/cache/kafka"
	"go-common/library/log"
)




//创建消费者
func(s *S) CreateConsumer(mConsumerSub *kafka.ConsumerSub ,conf *kafka.Config)(err error){

	var (
	  mSub *kafka.Sub
	)
	mSub,err=kafka.InitConsumer(mConsumerSub,conf)
	if err!=nil {
		log.Warn("CreateConsumer kafka.InitConsumer mConsumerSub=% conf=% err=%s",mConsumerSub,conf,err)
		return
	}

	for{

		select {

			case msg:=<-mSub.Datach:

				fmt.Println("msg=",msg)
		}

	}

}



//创建生产者生产消息
func(p *P) CreateProduct(mProductPub *kafka.ProductPub ,conf *kafka.Config)  {


}



func insertToMongo()  {

	for{
		select {

		case datach:=<-mongoChan:
			c :=SwichCol(string(datach.key))
			err:=c.Insert(datach)
			fmt.Println("err=",err )
		}

	}

}
