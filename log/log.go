package log

import (
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"fmt"
)

func log()  {
	config := make(map[string]interface{})
	config["filename"] = "../logs/logcollect.log"
	config["level"] = logs.LevelDebug//切换log Level来决定logs.Debug、Trace、Warn是否生效

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}
	//获取配置的路径和日志级别组成json str
	logs.SetLogger(logs.AdapterFile, string(configStr))

	logs.Debug("this is a test, my name is %s", "golang")
	logs.Trace("this is a trace, my name is %s", "cpp")
	logs.Warn("this is a warn, my name is %s", "c")
}