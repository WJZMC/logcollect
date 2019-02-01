package tail

import (
	"fmt"
	"time"
	"github.com/hpcloud/tail"
	"io"
)

func tails()  {
	filename := "../logs/my.log"
	//tmp,err:=ioutil.ReadFile(filename)
	//fmt.Println(tmp)
	tails, err := tail.TailFile(filename, tail.Config{
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
	var msg *tail.Line
	var ok bool
	for true {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("msg:", msg.Text,msg.Time,msg.Err)
	}
}
