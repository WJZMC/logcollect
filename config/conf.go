package config

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

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