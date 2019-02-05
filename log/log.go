package log

import (
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"fmt"
	"logcollect/config"
	"github.com/pkg/errors"
)
func formatLogLevel(level string) (logLevel int){
	switch level {
	case "debug":
		logLevel=logs.LevelDebug
	case "warn":
		logLevel=logs.LevelWarn
	default:
		logLevel=logs.LevelInfo
	}
	return
}
func InitLog() error {
	configTmp := make(map[string]interface{})
	configTmp["filename"] = config.AppConfig.LogPath
	configTmp["level"] = formatLogLevel(config.AppConfig.LogLevel)

	configStr, err := json.Marshal(configTmp)
	if err != nil {
		err:=fmt.Sprintf("InitLog failed, err:", err)
		return errors.New(err)
	}
	//获取配置的路径和日志级别组成json str
	logs.SetLogger(logs.AdapterFile, string(configStr))

	//logs.Debug("this is a test, my name is %s", "golang")
	//logs.Trace("this is a trace, my name is %s", "cpp")
	//logs.Warn("this is a warn, my name is %s", "c")
	return  nil
}