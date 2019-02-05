package main

import (
	"logcollect/config"
	"logcollect/log"
	"github.com/astaxie/beego/logs"
	"logcollect/tail"
	"runtime"
	"logcollect/kafka"
)

func main()  {
	err:=config.NewConfig(config.EncodeType_ini,"./conf/logcollect.conf")
	if err!=nil{
		logs.Error("config.NewConfig fail err=%v",err)
		return
	}
	err=log.InitLog()
	if err!=nil{
		logs.Error("log.InitLog fail err=%v",err)
		return
	}

	logs.Debug("load sucess conf=%v",*config.AppConfig)
	logs.Debug("========================%v==============","logs begin")

	tail.InitTail()

	kafka.InitProduct()

	for{
		runtime.GC()
	}


}
