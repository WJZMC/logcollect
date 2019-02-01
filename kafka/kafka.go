package kafka

import (
	"github.com/Shopify/sarama"
	"fmt"
	"time"
)
func product()  {

	config := sarama.NewConfig()
	config.ClientID="test3"
	config.Producer.RequiredAcks = sarama.WaitForAll//ack 确认服务器收到消息

	//随机分区，分片存储到多台服务器
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	config.Producer.Return.Successes = true//写成功配置

	//这里使用了同步客户端
	//还可以使用异步客户端
	client, err := sarama.NewSyncProducer([]string{"192.168.0.103:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}

	defer client.Close()

	go consumer()


	for{
		//构造一个消息实体
		msg := &sarama.ProducerMessage{}
		msg.Topic = "test3"
		msg.Value = sarama.StringEncoder("this is a good test, my message is good")


		//如果是同步客户端，发送完成后就成功写入到kafka了
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send message failed,", err)
			return
		}

		fmt.Printf("pid:%v offset:%v\n", pid, offset)
		time.Sleep(time.Second)
	}
}

func consumer()  {
	//config:=sarama.NewConfig()
	//config.Consumer.Return.Errors=true
	//config.Consumer.Group.

}