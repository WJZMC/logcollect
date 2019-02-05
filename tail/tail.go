package tail

import (
	"fmt"
	"time"
	"github.com/hpcloud/tail"
	"io"
	"logcollect/config"
)
var(
	TailfObj *Tailf
)
type tailReader struct {
	tails *tail.Tail
	collect config.CollectionLog
}
type TailMsg struct {
	Msg string
	Topic string
}
type Tailf struct {
	msgs []tailReader
	msgchan chan *TailMsg
}

func GetMsg() (msg string,topic string) {
	tailMsg:=<-TailfObj.msgchan
	msg = tailMsg.Msg
	topic = tailMsg.Topic
	return
}

func InitTail()  {
	TailfObj=&Tailf{
		msgchan:make(chan *TailMsg,config.AppConfig.MsgChanSize),
	}
	for _,v:=range config.AppConfig.Collections{
		//filename := "../logs/my.log"
		//tmp,err:=ioutil.ReadFile(filename)
		//fmt.Println(tmp)
		tails, err := tail.TailFile(v.LogPath, tail.Config{
			ReOpen:    true,//异常
			Follow:    true,
			Location:  &tail.SeekInfo{Offset: 0, Whence:io.SeekCurrent},//异常后重定向
			MustExist: false,//是否必须存在，允许不存在
			Poll:      true,//不断地查询
		})
		if err != nil {
			fmt.Println("tail file err:", err)
			return
		}
		tmp:=tailReader{
			tails:tails,
			collect:v,
		}
		TailfObj.msgs= append(TailfObj.msgs, tmp)

		go startCollect(tails,v.Topic)
	}
}

func startCollect(tails *tail.Tail,topic string)  {
	var msg *tail.Line
	var ok bool
	for true {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		msgstr:=fmt.Sprintf("msg:%v   %v  %v", msg.Text,msg.Time,msg.Err)
		tailMsg:=&TailMsg{
			Msg:msgstr,
			Topic:topic,
		}
		TailfObj.msgchan<-tailMsg
		//time.Sleep(time.Second)
	}
}
