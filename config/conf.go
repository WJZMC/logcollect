package config

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"errors"
)

const MsgChanSize  = 100
var(
	AppConfig *LogConfig
)
//demo
func configs()  {
	//初始化
	conf, err := config.NewConfig("ini", "../conf/logcollect.conf")
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}
	//端口
	port, err := conf.Int("server::port")
	if err != nil {
		fmt.Println("read server:port failed, err:", err)
		return
	}

	fmt.Println("Port:", port)
	log_level := conf.String("logs::log_level")
	fmt.Println("log_level:", log_level)

	log_path := conf.String("logs::log_path")
	fmt.Println("log_path:", log_path)
}

type CollectionLog struct {
	LogPath string
	Topic string
}

func LoadCollection(tmp config.Configer) (err error) {
	collect:=CollectionLog{}
	collect.LogPath=tmp.String("collections::LogPath")
	collect.Topic=tmp.String("collections::Topic")
	if len(collect.LogPath)<=0||len(collect.Topic)<=0{
		err=errors.New("collection null logpath or topic")
		return
	}

	AppConfig.Collections= append(AppConfig.Collections, collect)

	return
}

//加载配置
type LogConfig struct {
	LogLevel string//切换log Level来决定logs.Debug、Trace、Warn是否生效
	LogPath string
	MsgChanSize int//收集日志行管道容量
	Collections []CollectionLog
}

const (
	EncodeType_ini =iota
	EncodeType_json
	EncodeType_xml
	EncodeType_yaml
)
//encodeType: EncodeType_xxx
func NewConfig(encodeType int,filePath string) (err error) {
	var encodeName string
	switch encodeType {
	case EncodeType_ini:
		encodeName="ini"
	case EncodeType_json:
		encodeName="json"
	case EncodeType_xml:
		encodeName="xml"
	case EncodeType_yaml:
		encodeName="yaml"
	default:
		encodeName="ini"
	}

	tmp,err:=config.NewConfig(encodeName,filePath)
	if err!=nil{
		return
	}
	logPath:=tmp.String("logs::LogPath")
	logLevel:=tmp.String("logs::LogLevel")

	if len(logPath)<=0{
		err=errors.New("config null logpath")
		return
	}

	chanSize,err:=tmp.Int("logs::MsgChanSize")
	if err!=nil{
		chanSize=MsgChanSize
	}

	AppConfig =&LogConfig{
		LogPath:logPath,
		LogLevel:logLevel,
		MsgChanSize:chanSize,
	}

	LoadCollection(tmp)

	return

}



